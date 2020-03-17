package main

import (
	"fmt"
	"reflect"
)

// 一、定义
// ./method.go:7:6: cannot define new methods on non-local type http.Clien
// func (x http.Client) method() {}

// // ./method.go:9:6: cannot define new methods on non-local type int
// func (x int) method() {}

// type X int

// // ./method.go:14:6: invalid receiver type **X (*X is not a defined type)
// func (x **X) method() {}
// func main() {

// }

// --------------------------------------------------------------------------
// type X int

// //go:noinline
// func (x X) value() {
// 	println(x)
// }

// //go:noinline
// func (p *X) pointer() {
// 	println(p, &p)
// }

// func main() {
// 	var x X = 100
// 	x.value()   // 编译器实际上是call main.X.value(x_copy_value)
// 	x.pointer() // call main.(*X).pointer(&x_copy)

// 	p := &x
// 	p.value() // 编译器会把指针转为值，生成的指令和上面是相同的
// 	p.pointer()

// 	p2 := &p
// 	// ./method.go:40:4: calling method value with receiver p2 (type **X) requires explicit dereference
// 	// p2.value() 不能使用多级指针，只能先取出再调用
// 	(*p2).value()
// }

// --------------------------------------------------------------------------

// type X int

// //go:noinline
// func (x X) add(a int) X {
// 	return x + X(a)
// }

// // 两种调用方式生成的汇编指令是一样的，从实现上方法只是编译器自动插入receiver
// // 把第一种方式称为值模式（value）
// // 把第二种方式称为表达式模式（expression）
// func main() {
// 	var o X = 20
// 	// 设计上讲，值模式先拿到状态，再给这个状态做逻辑
// 	// 多次调用，状态维持
// 	o.add(99)
// 	// 表达式模式先拿到逻辑，再把状态当参数放进去
// 	// 多次调用，应传入相同的状态
// 	X.add(o, 99)
// }

// --------------------------------------------------------------------------
// go tool objdump -s "main\.main" method
// go tool objdump -s "A-fm" method
// type X int

// //go:noinline
// func (x *X) A(a int) {
// 	println(x, a)
// }

// func main() {
// 	var o X = 100
// 	v := o.A //把一个方法值赋值给一个变量，变量能直接调用，这个方法值的签名应为func(a int)，那么这个o本身去哪了？
// 	v(2)
// }

// 类似于闭包的原理，编译器生成一个main.(*X).A-fm(SB)的方法，o本身放在DX
// --------------------------------------------------------------------------
// 验证方法表达式和方法值的签名
type X int

func (x X) add(a int) X {
	return x + X(a)
}

func main() {
	var o X = 20

	e := X.add
	// func(main.X, int) main.X
	fmt.Println(reflect.TypeOf(e), e(o, 2))

	v := o.add
	// func(int) main.X
	fmt.Println(reflect.TypeOf(v), v(2))
}
