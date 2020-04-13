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
	commitIndex int
	lastApplied int
	nextIndex   []int
	matchIndex  []int
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
	if args.Term < rf.currentTerm {
		reply.Term = rf.currentTerm
		reply.VoteGrant = false
		return nil
	}

	if rf.voteFor == -1 {
		rf.voteFor = args.Candidate
		rf.currentTerm = args.Term
		reply.Term = rf.currentTerm
		reply.VoteGrant = true
		return nil
	}

	reply.Term = rf.currentTerm
	reply.VoteGrant = false
	return nil
}

type HeartbeatArgs struct {
	Term   int
	Leader int
}

type HeartbeatReply struct {
	Term int
}

func (rf *Raft) broadcastHeartBeat() {
	args := HeartbeatArgs{
		Term:   rf.currentTerm,
		Leader: rf.me,
	}

	for i := range rf.nodes {
		go func(i int) {
			var reply HeartbeatReply
			rf.sendHeartBeat(i, args, &reply)
		}(i)
	}
}

func (rf *Raft) sendHeartBeat(nodeId int, args HeartbeatArgs, reply *HeartbeatReply) {
	client, err := rpc.DialHTTP("tcp", rf.nodes[nodeId].address)
	if err != nil {
		log.Fatalf("rpc err: %v \n", err)
	}
	defer client.Close()

	_ = client.Call("Raft.HeartBeat", args, reply)

	// 收到心跳回复后，发现自己的term小于其他节点，说明已经过期，自己变为Follower，重新选举
	if reply.Term > rf.currentTerm {
		rf.currentTerm = reply.Term
		rf.voteFor = -1
		rf.voteCount = 0
		rf.role = Follower
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

	reply.Term = rf.currentTerm
	rf.heartbeatCh <- true
	return nil
}
