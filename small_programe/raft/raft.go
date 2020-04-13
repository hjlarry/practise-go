package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/rpc"
	"time"
)

type Node struct {
	address string
	connect bool
}

type Role int

const (
	Leader    Role = 1
	Candidate Role = 2
	Follower  Role = 3
)

type Raft struct {
	me          int  //当前节点id
	role        Role // 当前角色
	currentTerm int
	voteFor     int           // 给谁投票，未投时为-1
	voteCount   int           // 当前任期内获得的票数
	nodes       map[int]*Node // 集群内除自己的其他节点map

	toLeaderCh  chan bool // 其中取出true，说明计算出自己是Leader
	heartbeatCh chan bool // 其中取出true，说明收到leader发出的心跳

	logs        []LogEntry
	commitIndex int   // 被提交的最大索引
	lastApplied int   // 被应用到状态机的最大索引
	nextIndex   []int // 保存需要发送给每个节点的下一个条目索引
	matchIndex  []int // 保存已经复制给每个节点日志的最高索引
}

type LogEntry struct {
	LogTerm    int
	LogIndex   int
	LogCommand interface{}
}

func (rf *Raft) String() string {
	return fmt.Sprintf("Node: %d, role: %d, voteFor: %d", rf.me, rf.role, rf.voteFor)
}

func newNode(addr string) *Node {
	node := &Node{}
	node.address = addr
	return node
}

func (rf *Raft) rpc(port string) {
	_ = rpc.Register(rf)
	rpc.HandleHTTP()
	go func() {
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatalf("rpc err: %v \n", err)
		}
	}()
}

func (rf *Raft) start() {
	rf.role = Follower
	rf.currentTerm = 0
	rf.voteFor = -1
	rf.voteCount = 0
	rf.toLeaderCh = make(chan bool)
	rf.heartbeatCh = make(chan bool)

	go func() {
		rand.Seed(time.Now().UnixNano())
		// 不间断处理节点行为和RPC
		for {
			switch rf.role {
			case Follower:
				select {
				// 心跳超时，follower变为Candidate
				case <-time.After(time.Duration(rand.Intn(500-300)+300) * time.Millisecond):
					fmt.Printf("follower %d timeout \n", rf.me)
					rf.role = Candidate
				//	成功接收到心跳
				case <-rf.heartbeatCh:
					fmt.Printf("follower %d received heartbeat \n", rf.me)
				}
			case Candidate:
				fmt.Printf("Node: %d, candidate \n", rf.me)
				rf.currentTerm += 1
				rf.voteFor = rf.me
				rf.voteCount = 1

				go rf.broadcastVoteRequest()

				select {
				// 选举超时，candiate变为follower
				case <-time.After(time.Duration(rand.Intn(500-300)+300) * time.Millisecond):
					fmt.Printf("timeout, candiate %d become follower", rf.me)
					rf.role = Follower
				//	选举成功
				case <-rf.toLeaderCh:
					fmt.Printf("candiate %d become leader", rf.me)
					rf.role = Leader
					rf.nextIndex = make([]int, len(rf.nodes))
					rf.matchIndex = make([]int, len(rf.nodes))
					// 为每个节点初始化nextIndex和matchIndex，这里不考虑leader重新选举的情况
					for i := range rf.nodes {
						rf.nextIndex[i] = 1
						rf.matchIndex[i] = 0
					}

					// 模拟客户端每3秒发送一条command
					go func() {
						i := 0
						for {
							i++
							newEntry := LogEntry{
								LogTerm:    rf.currentTerm,
								LogIndex:   i,
								LogCommand: fmt.Sprintf("user send: %d", i),
							}
							rf.logs = append(rf.logs, newEntry)
							time.Sleep(time.Second * 3)
						}
					}()
				}
			case Leader:
				rf.broadcastHeartBeat()
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
}

type RequestVoteArgs struct {
	Term      int
	Candidate int
}

type RequestVoteReply struct {
	Term      int
	VoteGrant bool
}

func (rf *Raft) broadcastVoteRequest() {
	args := RequestVoteArgs{
		Term:      rf.currentTerm,
		Candidate: rf.me,
	}

	for i := range rf.nodes {
		go func(i int) {
			var reply RequestVoteReply
			rf.sendRequestVote(i, args, &reply)
		}(i)
	}

}

func (rf *Raft) sendRequestVote(nodeId int, args RequestVoteArgs, reply *RequestVoteReply) {
	client, err := rpc.DialHTTP("tcp", rf.nodes[nodeId].address)
	if err != nil {
		log.Fatalf("rpc err: %v \n", err)
	}
	defer client.Close()
	_ = client.Call("Raft.RequestVote", args, reply)

	if reply.Term > rf.currentTerm {
		rf.currentTerm = reply.Term
		rf.voteFor = -1
		rf.role = Follower
		return
	}

	// 每次成功接收到投票时，计算有没有超过半数，超过则通知toLeaderCh
	if reply.VoteGrant {
		rf.voteCount += 1

		if rf.voteCount > len(rf.nodes)/2+1 {
			rf.toLeaderCh <- true
		}
	}
}

func (rf *Raft) RequestVote(args RequestVoteArgs, reply *RequestVoteReply) error {
	// 作为其他节点，收到一个投票请求，确认是否要投票
	// Candidate节点过时，拒绝投票
	if args.Term < rf.currentTerm {
		reply.Term = rf.currentTerm
		reply.VoteGrant = false
		return nil
	}

	// 说明还没有投票给其他人，投票成功
	if rf.voteFor == -1 {
		rf.voteFor = args.Candidate
		rf.currentTerm = args.Term
		reply.Term = rf.currentTerm
		reply.VoteGrant = true
		return nil
	}

	// 可能已投票给其他人，投票失败
	reply.Term = rf.currentTerm
	reply.VoteGrant = false
	return nil
}

type HeartbeatArgs struct {
	Term   int
	Leader int
	// follower依据PrevLogIndex和PrevLogTerm和自己的本地日志对比，决定是否要同步Entries
	PrevLogIndex int        //新日志之前的索引
	PrevLogTerm  int        //PrevLogIndex对应的term
	Entries      []LogEntry // 准备存储的日志条目，心跳时为空
	LeaderCommit int        //Leader已经commit的索引值
}

type HeartbeatReply struct {
	Term      int
	Success   bool
	NextIndex int //如果Follower index小于leader，会告诉leader下次开始发送的索引位置
}

func (rf *Raft) broadcastHeartBeat() {
	for i := range rf.nodes {
		args := HeartbeatArgs{
			Term:         rf.currentTerm,
			Leader:       rf.me,
			LeaderCommit: rf.commitIndex,
		}
		prevLogIndex := rf.nextIndex[i] - 1
		// 如果没有可以发送的LogEntry
		if rf.getLastIndex() > prevLogIndex {
			args.PrevLogIndex = prevLogIndex
			args.PrevLogTerm = rf.logs[prevLogIndex].LogTerm
			args.Entries = rf.logs[prevLogIndex:]
			log.Printf("send entries: %v \n", args.Entries)
		}

		go func(i int, args HeartbeatArgs) {
			var reply HeartbeatReply
			rf.sendHeartBeat(i, args, &reply)
		}(i, args)
	}
}

func (rf *Raft) sendHeartBeat(nodeId int, args HeartbeatArgs, reply *HeartbeatReply) {
	client, err := rpc.DialHTTP("tcp", rf.nodes[nodeId].address)
	if err != nil {
		log.Fatalf("rpc err: %v \n", err)
	}
	defer client.Close()

	_ = client.Call("Raft.HeartBeat", args, reply)

	if reply.Success {
		if reply.NextIndex > 0 {
			rf.nextIndex[nodeId] = reply.NextIndex
			rf.matchIndex[nodeId] = rf.nextIndex[nodeId] - 1
			// TODO
			// 如果大于半数节点同步成功
			// 1. 更新 leader 节点的 commitIndex
			// 2. 返回给客户端
			// 3. 应用状态就机
			// 4. 通知 Followers Entry 已提交
		}
	} else {
		// 收到心跳回复后，发现自己的term小于其他节点，说明已经过期，自己变为Follower，重新选举
		if reply.Term > rf.currentTerm {
			rf.currentTerm = reply.Term
			rf.voteFor = -1
			rf.voteCount = 0
			rf.role = Follower
		}
	}

}

func (rf *Raft) HeartBeat(args HeartbeatArgs, reply *HeartbeatReply) error {
	if args.Term < rf.currentTerm {
		reply.Term = rf.currentTerm
		return nil
	}

	if args.Term > rf.currentTerm {
		rf.currentTerm = args.Term
		rf.voteFor = -1
		rf.role = Follower
	}

	// 没有entries，仅heartbeat的情况
	rf.heartbeatCh <- true
	if len(args.Entries) == 0 {
		reply.Success = true
		reply.Term = rf.currentTerm
		return nil
	}

	// 有entries
	// 但leader维护的logindex大于follower的，follower之前失联过，follower告知当前自己的最大索引，下次心跳leader返回相应的
	if args.PrevLogIndex > rf.getLastIndex() {
		reply.Success = false
		reply.Term = rf.currentTerm
		reply.NextIndex = rf.getLastIndex() + 1
		return nil
	}

	// 添加leader发送的log并更新commitIndex
	rf.logs = append(rf.logs, args.Entries...)
	rf.commitIndex = rf.getLastIndex()
	reply.Success = true
	reply.Term = rf.currentTerm
	reply.NextIndex = rf.getLastIndex() + 1
	return nil
}

func (rf *Raft) getLastIndex() int {
	if len(rf.logs) == 0 {
		return 0
	}
	return rf.logs[len(rf.logs)-1].LogIndex
}
