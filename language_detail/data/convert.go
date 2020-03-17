package main

import "unsafe"
import "fmt"
/* 类型转换 */
// func main(){
// 	type A int
// 	type B int
// 	var x A = 100
// 	// var y B = x  // cannot use x (type A) as type B in assignment
// 	// 只能显式转换
// 	var y B = B(x)
// 	_ = y


// 	type C struct{
// 		// tag并不影响类型转换
// 		x int  `tag:all`
// 	}
// 	type D struct{
// 		x int
// 	}
// 	// 类型转换的前提是内存布局和名字都相同
// 	type E struct{
// 		a int
// 	}
// 	c := C{x:1}
// 	var d D = D(c)
// 	// var e E = E(c) // cannot convert c (type C) to type E
// 	_ = d
// }



/* 局部转换 */
func main(){
	type A struct{
		x int
		y string
		z int
	}
	type B struct{
		x int
	}
	// b := B{99}
	// // a := A(b)  cannot convert b (type B) to type A
	// // 通过内存地址操作，先用unsafe.Pointer获得一个无类型指针，然后直接读地址
	// var a A = *(*A)(unsafe.Pointer(&b))
	// fmt.Printf("%#v", a)  //main.A{x:99, y:"syntax error scanning boolean", z:5097440}

	aa := A{123, "hello", 88}
	var bb B = *(*B)(unsafe.Pointer(&aa))
	fmt.Printf("%#v", bb)  // main.B{x:123}

	type C struct{
		c string
		d byte
	}
	// 通过地址操作可以直接拿到A中y的地址，即使C中标签名和类型不对也不会影响基于内存地址的操作
	// 这种基于指针的操作往往是系统编程中提升运行效率的方式
	var cc C = *(*C)(unsafe.Pointer(&aa.y))
	fmt.Printf("%#v", cc)  //main.C{c:"hello", d:0x58}
}