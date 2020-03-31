package main

import (
	"sync"
)

// 进程退出时并不会等待gorutine结束，有一些方法可以等待goroutine执行
// 一、使用sleep等待goroutine执行
// import "time"

// func main() {
// 	go func() {
// 		println("hello world")
// 	}()
// 	// 虽然这种等待方式看起来蠢，但它没有对被观察对象做任何侵入，有时候会很有用
// 	time.Sleep(time.Second)
// }

// 二、使用通道阻塞住
// 1. goroutine执行后可以发送个信号
// func main() {
// 	c := make(chan struct{})
// 	go func() {
// 		println("hello world")
// 		// 发送什么数据无所谓，空结构体不占内存
// 		c <- struct{}{}
// 	}()
// 	<-c
// }

// 2. 可以关闭通道，也是一种发送信号
// func main() {
// 	c := make(chan struct{})
// 	go func() {
// 		defer close(c)
// 		println("hello world")
// 	}()
// 	// 关闭通道后会立即接收到1个零值
// 	<-c
// }

// 三、使用waitgroup等待
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		println("a")
	}()
	go func() {
		defer wg.Done()
		println("b")
	}()

	println("main")
	wg.Wait()
}

// 四、发送信号和关闭通道还可以给多个goroutine编排顺序
// func main() {
// 	a, b := make(chan struct{}), make(chan struct{})
// 	go func() {
// 		<-a
// 		println("a")
// 	}()
// 	go func() {
// 		<-b
// 		println("b")
// 	}()

// 	time.Sleep(time.Second)
// 	println("main")

// 	a <- struct{}{}
// 	b <- struct{}{}
// 	time.Sleep(time.Second)
// }
// 关闭通道的方式就显得简单一些
// func main() {
// 	x := make(chan struct{})
// 	go func() {
// 		<-x
// 		println("a")
// 	}()
// 	go func() {
// 		<-x
// 		println("b")
// 	}()

// 	time.Sleep(time.Second)
// 	println("main")

// 	close(x)
// 	time.Sleep(time.Second)
// }
