package main

import (
	"reflect"
	"testing"
	"unsafe"
)

// go test -bench . -benchmem str_test.go
// BenchmarkTestStr/safe-2         	205193720	         5.52 ns/op	       0 B/op	       0 allocs/op
// BenchmarkTestStr/unsafe-2       	1000000000	         0.279 ns/op	       0 B/op	       0 allocs/op
func BenchmarkTestStr(b *testing.B) {
	bs := []byte("hello, world!")
	b.Run("safe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = string(bs)
		}
	})
	b.Run("unsafe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = *(*reflect.StringHeader)(unsafe.Pointer(&bs))
		}
	})
}
