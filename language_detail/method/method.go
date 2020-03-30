package main

import (
	"fmt"
)

// 一、定义
// 不能为非本地包定义方法
// ./method.go:7:6: cannot define new methods on non-local type http.Client
// func (x http.Client) method() {}

// 也不能为系统内置的类型定义方法
// // ./method.go:9:6: cannot define new methods on non-local type int
// func (x int) method() {}

// type X int

// // ./method.go:14:6: invalid receiver type **X (*X is not a defined type)
// func (x **X) method() {}
// func main() {

// }

// --------------------------------------------------------------------------
// 二、receiver为值或指针
// 用T的情况：无需修改状态的小对象或固定值，引用类型、字符串、函数等指针包装对象
// 用*T的情况：要修改实例的状态、大对象、包含Mutex等同步字段、其他无法确定的情况
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
// 三、两种调用方式
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
// 四、方法赋值给变量
// 1. 反编译
// go tool objdump -s "main\.main" method
// go tool objdump -s "A-fm" method
// 类似于闭包的原理，编译器生成一个main.(*X).A-fm(SB)的方法，o本身放在DX
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

// 2. 分为expression和value两种状态
type M int
func (m M) test() {
	fmt.Printf("test.m: %p, %d\n", &m, m)
}
func main() {
	var m M = 26
	fmt.Printf("main.m: %p, %d\n", &m, m)
	// 通过类型引用的method expression会被还原为普通函数，调用时receiver作为参数传入
	f1 := M.test
	f1(m)
	// 类型是T或*T都可以，目标方法在其方法集中即可
	f2 := (*M).test
	f2(&m) //须按方法集中的签名传入正确的参数

	M.test(m)
	(*M).test(&m)

	p := &m
	m++
	// 基于指针或实例引用的method value依旧正常调用
	f3 := m.test //但当其赋值给变量或作为参数传递时，会立即复制该方法所需的receiver对象并与其绑定
	m++
	f4 := p.test
	m++
	fmt.Printf("main.m: %p, %d\n", p, m)
	f3()
	f4()
}
// --------------------------------------------------------------------------
// 五、验证方法表达式和方法值的签名
//type X int
//
//func (x X) add(a int) X {
//	return x + X(a)
//}
//
//func main() {
//	var o X = 20
//
//	e := X.add
//	// func(main.X, int) main.X
//	fmt.Println(reflect.TypeOf(e), e(o, 2))
//
//	v := o.add
//	// func(int) main.X
//	fmt.Println(reflect.TypeOf(v), v(2))
//}

// type A int

// type X struct {
// 	A  // 匿名字段是语法糖，相当于一个字段的名称是A，类型也是A
// 	*A // 而*A相当于名称是A，类型是*A，所以两个同名字段不能同时存在
// }

// --------------------------------------------------------------------------
// 六、匿名字段
// 使用匿名字段可以实现看上去像是继承的调用效果，但实际上是组合
// type A int
// type B struct {
// 	A
// 	x int
// }

// //go:noinline
// func (a A) ma() {
// 	println("A.a")
// }

// //go:noinline
// func (b B) mb() {
// 	println("B.b")
// }

// func main() {
// 	o := B{A: 100, x: 200}
// 	o.mb()

// 	// 通过反汇编，以下两种方式的指令都是CALL main.A.ma(SB)
// 	o.ma()
// 	o.A.ma()
// }

// --------------------------------------------------------------------------
// 七、同名遮蔽
// 1. 通过同名遮蔽，可以做到类似于继承中override的效果
// 这种查找相应的方法进行调用和python的mro不同，是编译期编译器就已经确定的
// type A int
// type B struct {
// 	A
// }

// func (A) test() {
// 	println("A.test")
// }

// func (B) test() {
// 	println("B.test")
// }

// func main() {
// 	o := B{}
// 	o.test()
// 	o.A.test()
// }

// --------------------------------------------------------------------------
// 2. 但是当嵌套的深度一样时会报错，只能去显式调用相应的方法
// ./method2.go:79:3: ambiguous selector o.test
//type A int
//type B int
//type C struct {
//	A
//	B
//}
//
//func (A) test() {
//	println("A.test")
//}
//
//func (B) test() {
//	println("B.test")
//}
//
//func main() {
//	o := C{}
//	o.test()
//}