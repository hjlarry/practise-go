package main

// type A int

// type X struct {
// 	A  // 匿名字段是语法糖，相当于一个字段的名称是A，类型也是A
// 	*A // 而*A相当于名称是A，类型是*A，所以两个同名字段不能同时存在
// }

// --------------------------------------------------------------------------
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
// 通过同名遮蔽，可以做到类似于继承中override的效果
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
// 但是当嵌套的深度一样时会报错，只能去显式调用相应的方法
// ./method2.go:79:3: ambiguous selector o.test
type A int
type B int
type C struct {
	A
	B
}

func (A) test() {
	println("A.test")
}

func (B) test() {
	println("B.test")
}

func main() {
	o := C{}
	o.test()
}
