package main

import "unsafe"

func main(){
	// 变量定义时已经分配内存
	var a []int
	println(unsafe.Sizeof(a))

	// nil并不是关键字
	nil := 100
	println(nil)
}