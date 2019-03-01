package main

import (
	"errors"
	"fmt"
)

var global = 1

// Allglobal ... 首字母大写 则其他包能访问
var Allglobal = 10

// A ...
const A = 100

// VarFunc ... 定义变量
func VarFunc() {
	fmt.Println("hello world")
	// var s1 int = 10
	var s2 = 10
	s3 := 10
	fmt.Println(s2)
	fmt.Println(s3)
	for index := 0; index < 4; index++ {
		// fmt.Println(s1)
	}
	var local = 2
	fmt.Println(global)
	fmt.Println(local)

	// 有符号整数，可以表示正负
	var a int8 = 1  // 1 字节
	var b int16 = 2 // 2 字节
	var c int32 = 3 // 4 字节
	var d int64 = 4 // 8 字节
	fmt.Println(a, b, c, d)

	// 无符号整数，只能表示非负数
	var ua uint8 = 1
	var ub uint16 = 2
	var uc uint32 = 3
	var ud uint64 = 4
	fmt.Println(ua, ub, uc, ud)

	// int 类型，在32位机器上占4个字节，在64位机器上占8个字节
	// var e int = 5
	// var ue uint = 5
	// fmt.Println(e, ue)

	// bool 类型
	// var f bool = true
	// fmt.Println(f)

	// 字节类型
	var j byte = 'a'
	fmt.Println(j)

	// 字符串类型
	// var g string = "abcdefg"
	// fmt.Println(g)

	// 浮点数
	// var h float32 = 3.14
	// var i float64 = 3.141592653
	// fmt.Println(h, i)

	var aArray [9]int
	fmt.Println(aArray)

	var aa = [10]int{1, 3, 4}
	// var bb [10]int = [10]int{1, 3, 4}
	cc := [10]int{1, 3, 4}
	fmt.Println(aa)
	// fmt.Println(bb)
	fmt.Println(cc)

	var seq [9]int
	for i := 0; i < len(seq); i++ {
		seq[i] = i * i
	}
	fmt.Println(seq)

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
}
