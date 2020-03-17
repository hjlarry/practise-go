package main

import "sync"

/*
[ubuntu] ~/.mac/gocode $ go vet
# _/root/.mac/gocode
./go_vet.go:7:8: assignment copies lock value to m2: sync.Mutex
./go_vet.go:8:6: assignment copies lock value to _: sync.Mutex
*/
func main() {
	var m sync.Mutex
	m2 := m
	_ = m2
}
