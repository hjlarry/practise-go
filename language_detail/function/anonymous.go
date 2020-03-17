package main

// 匿名函数编译后仍然是有符号和名字的
// 匿名函数可被内联
// go build -gcflags "-N -l -S" 2>a.txt anonymous.go
func main() {
	f := func() { println("abc") }
	f()

	func() { println("def") }()
}
