package main

// 反汇编说明多返回值就是根据返回值的数量提前分配好内存
// func test(x int) (int, int) {
// 	return x + 1, x + 2
// }

// func main() {
// 	a, b := test(100)
// 	println(a, b)
// }

// 命名返回值本质上就是提前给返回值位置一个名字
// func test() (x int) {
// 	x = 100
// 	return 200
// }

// func main() {
// 	a := test()
// 	println(a)
// }

// 整体的步骤
// defer方法的定义、result = 200、defer方法的执行、return result
func test() (x int) {
	x = 100
	defer func() { // 可以分解为两步 defer方法的定义，defer方法的执行
		x = 88
	}()
	return 200 // 可以分解为两步 result = 200; return result
}

func main() {
	a := test()
	println(a)
}
