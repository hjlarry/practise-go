package main

import (
	"fmt"
)

// 一、 栈帧
/*
(gdb) info frame
Stack level 0, frame at 0xc00002e748:
Arglist at 0xc00002e698, args: x=100, y= []uint8 = {...}
Locals at 0xc00002e698, Previous frame's sp is 0xc00002e748 (gdb) x/3xg &x
0xc00002e748: 0x0000000000000064 0x000000c000068010 0xc00002e758: 0x0000000000000003
(gdb) x/3xb 0x000000c000068010 0xc000068010: 0x01 0x02 0x03
*/
//func test(x int, y ...byte) {
//	println(y)
//}
//
//func main() {
//	test(100, 1, 2, 3)
//}

//二、 传值和传指针的区别
//func add1(a int) int {
//	a = a + 1
//	return a
//}
//
//func add2(a *int) int {
//	*a = *a + 1
//	return *a
//}
//
//// TestAdd ...
//func main() {
//	x := 3
//	fmt.Println("x=", x)
//	x1 := add1(x)
//	fmt.Println("x1=", x1)
//	fmt.Println("x=", x)
//
//	y := 3
//	fmt.Println("y=", y)
//	y1 := add2(&y)
//	fmt.Println("y1=", y1)
//	fmt.Println("y=", y)
//}

// 三、 一次接收多个参数
//func sum(ops ...int) int {
//	ret := 0
//	for _, op := range ops {
//		ret += op
//	}
//	return ret
//}
//
//func main() {
//	fmt.Println(sum(1, 2, 3, 4))
//}

// 四、函数作为参数传入
type testInt func(int) bool

func isOdd(integer int) bool {
	return integer%2 != 0
}

func isEven(integer int) bool {
	return integer%2 == 0
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 3, 4, 6, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("odd = ", odd)
	even := filter(slice, isEven)
	fmt.Println("even = ", even)
}
