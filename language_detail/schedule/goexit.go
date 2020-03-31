package main

import (
	"runtime"
	"testing"
)

// func main() {
// 	c := make(chan struct{})
// 	go func() {
// 		defer close(c)
// 		func() {
// 			// goexit能保证终止整个任务，而return只能终止当前的这个函数
// 			// runtime.Goexit()
// 			return
// 		}()
// 		println("done")
// 	}()
// 	<-c
// 	println("exit")
// }

// 在main中执行Goexit则会等待其他Goroutine执行完毕，然后让进城崩溃
func main() {
	go func() {
		println("a")
		println("b")
		println("c")
	}()
	runtime.Goexit()
}

// Goexit立即终止当前任务，但不能用在main.main
func TestGoroutine8(t *testing.T) {
	exit := make(chan struct{})
	go func() {
		defer close(exit)
		defer println("a")
		func() {
			defer func() {
				println("b", recover() == nil)
			}()
			func() {
				println("c")
				runtime.Goexit()
				println("c done") //不执行
			}()
			println("b done") //不执行
		}()
		println("a done") //不执行
	}()
	<-exit
	println("main exit") //执行
}
