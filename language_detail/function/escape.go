package main

//go:noinline
func test() *int {
	x := 100
	return &x
}

// 使用go build -gcflags "-m=2" escape.go可以得到一些编译的信息，例如x的逃逸
// 但因为链接器和编译器是阶段性输出这些信息的，即使内联了也会输出x的逃逸信息
// 这时应该主要看反汇编的结果，编译信息只能参考、不是依据
func main() {
	t := test()
	println(t)
}

// func main() {
// 	x := 100
// 	// 使用内置的println不会逃逸，而fmt会，因为fmt接收的是接口，会有x的复制行为
// 	// println(x)
// 	fmt.Println(x)
// }
