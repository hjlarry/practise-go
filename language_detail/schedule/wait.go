package main

import (
	"sync"
	"testing"
	"time"
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


// 进程退出时并不会等待gorutine结束，可以使用通道阻塞并发出退出信号
func TestGoroutine4(t *testing.T) {
	exit := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 3)
		println("gorutine done")
		close(exit)
	}()
	println("main")
	<-exit
	println("main done")
}

// 等待多个任务结束，可以使用WaitGroup
func TestGoroutine5(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Second)
			println("gorutine", id, "done")
		}(i)
	}
	println("main")
	wg.Wait()
	println("main done")
}
