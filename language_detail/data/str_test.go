package main

import (
	"bytes"
	"reflect"
	"testing"
	"unsafe"
)

/* safe和unsafe的方式字符串转换性能比较 */
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


/* 直接拼接、使用join、使用bytes.buffer的性能比较 */
func plus() string {
	var s string
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	return s
}

//BenchmarkPlus-4   	   10000	    106563 ns/op
func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plus()
	}
}

func join(a []string, sep string) string {
	n := len(sep) * (len(a) - 1)  // 统计分隔符长度
	for i := 0; i < len(a); i++ { // 统计所有待拼接字符长度
		n += len(a[i])
	}
	b := make([]byte, n) //一次分配所需长度的数组空间
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

func testJoin() string {
	s := make([]string, 1000) //准备足够内存，避免中途扩张
	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}
	return join(s, "")
}

//BenchmarkJoin-4   	  200000	      9145 ns/op
func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testJoin()
	}
}

func buffer() string {
	var b bytes.Buffer
	b.Grow(1000) //准备足够内存，避免中途扩张
	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}
	return b.String()
}

//BenchmarkBuffer-4   	  200000	      6213 ns/op
func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buffer()
	}
}
