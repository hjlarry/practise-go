package main

import (
	"flag"
	"strings"
)

func main() {
	port := flag.String("port", ":9091", "rpc listen port")
	cluster := flag.String("cluster", "127.0.0.1:9091", "use comma to seprate")
	id := flag.Int("id", 1, "node id")

	flag.Parse()
	clusters := strings.Split(*cluster, ",")
	nodes := make(map[int]*Node)
	for k, v := range clusters {
		nodes[k] = newNode(v)
	}

	raft := &Raft{}
	raft.me = *id
	raft.nodes = nodes

	raft.rpc(*port)
	raft.start()
	select {}
}
