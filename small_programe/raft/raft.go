package raft

import (
	"fmt"
	"log"
	"math/rand"
	"net/rpc"
	"sync"
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
	mu sync.Locker

	me          int
	role        Role
	currentTerm int
	voteFor     int
	voteCount   int
	nodes       map[int]*Node

	toLeaderCh  chan bool
	heartbeatCh chan bool
}

func newNode(addr string) *Node {
	node := &Node{}
	node.address = addr
	return node
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
		for {
			switch rf.role {
			case Follower:
				select {
				case <-rf.heartbeatCh:
					fmt.Printf("follower %d received heartbeat \n", rf.me)
				case <-time.After(time.Duration(rand.Intn(500-300)+300) * time.Millisecond):
					fmt.Printf("follower %d timeout \n", rf.me)
					rf.role = Candidate
				}
			case Candidate:
				fmt.Printf("Node: %d, candidate \n", rf.me)
				rf.currentTerm += 1
				rf.voteFor = rf.me
				rf.voteCount = 1

				go rf.broadcastVoteRequest()

				select {
				case <-time.After(time.Duration(rand.Intn(500-300)+300) * time.Millisecond):
					fmt.Printf("timeout, candiate %d become follower")
					rf.role = Follower
				case <-rf.toLeaderCh:
					fmt.Printf("candiate %d become leader")
					rf.role = Leader
				}
			}
		}
	}()
}

type requestVoteArgs struct {
	Term      int
	Candidate int
}

type requestVoteReply struct {
	Term      int
	VoteGrant bool
}

func (rf *Raft) broadcastVoteRequest() {
	args := requestVoteArgs{
		Term:      rf.currentTerm,
		Candidate: rf.me,
	}

	for i := range rf.nodes {
		go func(i int) {
			var reply requestVoteReply
			rf.sendRequestVote(i, args, &reply)
		}(i)
	}

}

func (rf *Raft) sendRequestVote(nodeId int, args requestVoteArgs, reply *requestVoteReply) {
	client, err := rpc.Dial("tcp", rf.nodes[nodeId].address)
	if err != nil {
		log.Fatalf("rpc err: %v \n", err)
	}
	defer client.Close()

	client.Call("Raft.RequestVote", args, reply)

	if reply.Term > rf.currentTerm {
		rf.currentTerm = reply.Term
		rf.voteFor = -1
		rf.role = Follower
		return
	}

	if reply.VoteGrant {
		rf.voteCount += 1

		if rf.voteCount > len(rf.nodes)/2+1 {
			rf.toLeaderCh <- true
		}
	}
}

func (rf *Raft) RequestVote(args requestVoteArgs, reply requestVoteReply) error {
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
