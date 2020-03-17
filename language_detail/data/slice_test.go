package main

import (
	"testing"
)

func array() [1024]int {
	var x [1024]int
	for i := 0; i < 1024; i++ {
		x[i] = i
	}
	return x
}

func slice() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

// 并非所有时候都适合用切片代替数组，因为切片底层数组可能会在堆上分配内存，且小数组在栈上拷贝的消耗也未必比make大
// BenchmarkArray-4   	 2000000	       961 ns/op	       0 B/op	       0 allocs/op
func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

// BenchmarkSlice-4   	 1000000	      1791 ns/op	    8192 B/op	       1 allocs/op
func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}





