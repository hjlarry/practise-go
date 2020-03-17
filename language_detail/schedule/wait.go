package main

import (
	"sync"
)

// 一、使用sleep等待goroutine执行
// import "time"

// func main() {
// 	go func() {
// 		println("hello world")
// 	}()
// 	// 虽然这种等待方式看起来蠢，但它没有对被观察对象做任何侵入，有时候会很有用
// 	time.Sleep(time.Second)
// }

// 二、使用通道发送信号
// func main() {
// 	c := make(chan struct{})
// 	go func() {
// 		println("hello world")
// 		// 发送什么数据无所谓，空结构体不占内存
// 		c <- struct{}{}
// 	}()
// 	<-c
// }

// 三、关闭通道来发送信号
// func main() {
// 	c := make(chan struct{})
// 	go func() {
// 		defer close(c)
// 		println("hello world")
// 	}()
// 	// 关闭通道后会立即接收到1个零值
// 	<-c
// }

// 四、多个goroutine编排顺序
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

// 五、使用waitgroup等待
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
