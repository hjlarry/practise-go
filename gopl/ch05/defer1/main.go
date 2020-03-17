package main

import "fmt"

// 当f(0)被调用时，先发生panic异常，之前被defer的函数依次执行；中断执行后，输出panic信息和堆栈信息
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func main() {
	f(3)
}
