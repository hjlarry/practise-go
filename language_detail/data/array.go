package main

import "fmt"

// func main() {
// 	// unsafe.Sizeof(a) 编译期可以计算出，所以能用来做数组长度的定义
// 	a := 0x11223344 //小端： int [44 33 22 11 00 00 00 00]  通过gdb x/8xb &a 查看大小端
// 	b := *(*[unsafe.Sizeof(a)]byte)(unsafe.Pointer(&a))
// 	fmt.Printf("%x \n", b)
// }

// func main() {
// 	var a [3]int
// 	var b [3]*int //指针数组

// 	var p *[3]int = &a //数组指针
// 	p[0] = 10

// 	var p1 *int = &a[2]
// 	*p1 = 30

// 	b[0] = &a[0]
// 	b[1] = &a[1]
// 	b[2] = &a[2]

// 	fmt.Println(a, b)
// }

func test(x [3]int) {
	fmt.Println(x)
}

func main() {
	a := [3]int{0x11, 0x22, 0x33}
	test(a)
}

// 不存在动态数组，所谓的动态数组只是采用数组的语法访问特征，但是底层可能是链表、字典等复合结构
