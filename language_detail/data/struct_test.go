package main

import (
	"sync"
	"testing"
)

// CPU cache是以行为单位缓存的，每个cpu读cache就读一行
// 每行多大可以通过cat /proc/cpuinfo | grep cache_alignment， 一般为64
// 那么当示例中的counter没有按64字节对齐的时候，每个cpu就会同时缓存多个counter
// 当对某个counter有修改的时候，其他cpu为了保证缓存一致性会把整行的cache都删掉
// 由于go中goroutine的调度机制复杂，无法排除一些干扰因素，以下测试可能需要多次运行来对比
func test1() {
	var counter [8]struct {
		x int
		// _ [64 - 8]byte
	}
	var wg sync.WaitGroup
	wg.Add(len(counter))
	for i := 0; i < len(counter); i++ {
		go func(id int) {
			for n := 0; n < 100000; n++ {
				counter[id].x++
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func test2() {
	var counter [8]struct {
		x int
		_ [64 - 8]byte
	}
	var wg sync.WaitGroup
	wg.Add(len(counter))
	for i := 0; i < len(counter); i++ {
		go func(id int) {
			for n := 0; n < 100000; n++ {
				counter[id].x++
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// [ubuntu] ~/.mac/gocode $ go test -bench . -benchmem struct_test.go
// goos: linux
// goarch: amd64
// BenchmarkTest/nopad-2         	    1180	    995898 ns/op	      80 B/op	       2 allocs/op
// BenchmarkTest/pad-2           	    1686	    705098 ns/op	     528 B/op	       2 allocs/op
func BenchmarkTest(b *testing.B) {
	b.Run("nopad", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			test1()
		}
	})
	b.Run("pad", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			test2()
		}
	})
}
