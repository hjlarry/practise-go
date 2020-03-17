package main

import (
	"runtime"
	"time"
	"unsafe"
)

// go build ptr.go && GODEBUG=gctrace=1 ./ptr
// 观察垃圾是否会回收
func main() {
	type data struct {
		x int
		_ [1 << 20]byte
	}

	// var s []*int
	// var s []unsafe.Pointer
	var s []uintptr
	for i := 0; i < 100; i++ {
		d := data{x: 1}
		// s = append(s, &d.x)
		// s = append(s, unsafe.Pointer(&d.x))
		s = append(s, uintptr(unsafe.Pointer(&d.x)))
	}

	for {
		println(len(s), s[0])
		runtime.GC()
		time.Sleep(time.Second)
	}
}
