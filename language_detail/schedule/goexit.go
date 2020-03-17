package main

import "runtime"

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
