package main

import (
	"testing"
)

// 将发往通道的数据打包，可减少传输次数，提升性能
const (
	max     = 5000000
	block   = 500
	bufsize = 100
)

func testNoBlock() {
	done := make(chan struct{})
	c := make(chan int, bufsize)
	go func() {
		count := 0
		for x := range c {
			count += x
		}
		close(done)
	}()
	for i := 0; i < max; i++ {
		c <- i
	}
	close(c)
	<-done
}

func testBlock() {
	done := make(chan struct{})
	c := make(chan [block]int, bufsize)
	go func() {
		count := 0
		for a := range c {
			for _, x := range a {
				count += x
			}
		}
		close(done)
	}()
	for i := 0; i < max; i += block {
		var b [block]int
		for n := 0; n < block; n++ {
			b[n] = i + n
			if i+n == max-1 {
				break
			}
		}
		c <- b
	}
	close(c)
	<-done
}

// BenchmarkTestNoBlock-4   	       3	 421774198 ns/op	    1152 B/op	       2 allocs/op
func BenchmarkTestNoBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testNoBlock()
	}
}

// BenchmarkTestBlock-4   	     100	  12488588 ns/op	  401527 B/op	       2 allocs/op
func BenchmarkTestBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBlock()
	}
}
