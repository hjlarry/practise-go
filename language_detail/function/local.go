package main

import (
	"runtime"
	"time"
)

// 局部变量的生命周期
func main() {
	x := make([]byte, 1<<20)
	runtime.SetFinalizer(&x, func(o *[]byte) { println("dead") }) //添加x的析构函数
	for i := 0; i < 2; i++ {
		// 说明局部变量并不一定是随栈帧一起被回收的
		runtime.GC()
		time.Sleep(time.Second)
	}
	println("aaa")
	// 如果想延长某变量的生命周期，相当于_ = x
	// runtime.KeepAlive(&x)
	println("bbb")
}
