package main

import (
	"fmt"
)

func main() {
	x := 10
	var p *int = &x //获取地址，保存到指针变量
	*p += 20        // 用指针间接引用，并更新对象
	fmt.Println(p, *p, &x, x)

	m := map[string]int{"a": 1}
	//fmt.Println(&m["a"])  报错：can not take the address of m["a"]
	fmt.Println(&m)

	p2 := &x
	//p2++ 报错：p2++ (non-numeric type *int)
	//var p3 *int= p2+1 报错
	fmt.Println(p2 == p)
}

// go build ptr.go && GODEBUG=gctrace=1 ./ptr
// 观察垃圾是否会回收
//func main() {
//	type data struct {
//		x int
//		_ [1 << 20]byte
//	}
//
//	// var s []*int
//	// var s []unsafe.Pointer
//	var s []uintptr
//	for i := 0; i < 100; i++ {
//		d := data{x: 1}
//		// s = append(s, &d.x)
//		// s = append(s, unsafe.Pointer(&d.x))
//		s = append(s, uintptr(unsafe.Pointer(&d.x)))
//	}
//
//	for {
//		println(len(s), s[0])
//		runtime.GC()
//		time.Sleep(time.Second)
//	}
//}
