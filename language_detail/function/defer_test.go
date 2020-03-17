package main

import (
	"sync"
	"testing"
)

var m sync.Mutex

func test1() {
	m.Lock()
	m.Unlock()
}

// 使用defer虽然造成了一定的性能损耗，但它会确保程序其他的逻辑不会造成意外情况，是我们大多数情况的选择
// 在一些锁的粒度很细、锁的使用很频繁的场景，我们基于性能考虑可以不用defer
func test2() {
	m.Lock()
	defer m.Unlock()
}

// go test -bench . defer_test.go
// BenchmarkTestDefer/nodefer-2         	93130644	        15.6 ns/op
// BenchmarkTestDefer/defer-2           	34221590	        48.9 ns/op
func BenchmarkTestDefer(b *testing.B) {
	b.Run("nodefer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			test1()
		}
	})
	b.Run("defer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			test2()
		}
	})
}
