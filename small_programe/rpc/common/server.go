package main

import (
	p "./proto"
	"log"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type KV struct {
	mu   sync.Mutex
	data map[string]string
}

func (kv *KV) Get(args *p.GetArgs, reply *p.GetReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	val, ok := kv.data[args.Key]
	if ok {
		reply.Err = p.OK
		reply.Value = val
	} else {
		reply.Err = p.ErrNoKey
		reply.Value = ""
	}
	return nil
}

func (kv *KV) Put(args *p.PutArgs, reply *p.PutReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[args.Key] = args.Value
	reply.Err = p.OK
	return nil
}

func server() {
	kv := new(KV)
	kv.data = map[string]string{}
	rpcs := rpc.NewServer()
	_ = rpcs.Register(kv)
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen err:", err)
	}
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				break
			}
			go rpcs.ServeConn(conn)
		}
		_ = l.Close()
	}()
}

func main() {
	server()
	for {
		time.Sleep(time.Second)
	}
}
