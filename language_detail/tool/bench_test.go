package main

import "testing"

/*
[ubuntu] ~/.mac/gocode $ go test -bench . -benchmem
goos: linux
goarch: amd64
BenchmarkAdd-2   	1000000000	         0.560 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	_/root/.mac/gocode	0.641s
*/
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
	}
}
