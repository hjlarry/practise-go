package main

import (
	p "./proto"
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func connect() *rpc.Client {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func put(key string, value string) {
	client := connect()
	args := p.PutArgs{
		Key:   key,
		Value: value,
	}
	reply := p.PutReply{}
	err := client.Call("KV.Put", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	_ = client.Close()
}

func get(key string) {
	client := connect()
	args := p.GetArgs{
		Key: key,
	}
	reply := p.GetReply{}
	err := client.Call("KV.Get", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	_ = client.Close()
	fmt.Printf("%v", reply)
}

func main() {
	put("helloworld", "123456")
	time.Sleep(time.Millisecond)
	get("helloworld")
}
