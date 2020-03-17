package main

/* break跳转 */
// func main() {
// 	// x:
// 	for i := 0; i < 10; i++ {

// 	}
// x:
// 	for i := 0; i < 10; i++ {
// 		break x // break+label时必须使跳出标签和循环直接关联
// 	}
// }

/* goto跳转 */
// func main(){
// 	goto Done
// 	v := 10
// Done:
// 	println(v)  //goto Done jumps over declaration of v at ./break.go:19:4
// }

/* 代码块重构 */

func main() { 
	a := 1
	// 以下代码目的是构建一个封闭的作用域区间，使用goto跳转的方式，会使内部代码块逻辑依赖外部的标签
	// {
	// 	a := "abc"
	// 	println(a)
	// 	goto done
	// 	a = "xxx"
	// 	println(a)
	// }
	// 一种方式是重构为函数，在函数中可以控制自己的作用域和跳转逻辑
	// func() {
	// 	a := "abc"
	// 	println(a)
	// 	return
	// 	a = "xxx"
	// 	println(a)
	// }()
	// 另一种方式是重构为只执行一次的while循环
	for {
		a := "abc"
		println(a)
		break
		a = "xxx" 
		println(a)
	}
// done:
	println(a)
}