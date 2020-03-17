package main

// 一、 单向通道
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	// 通道默认双向
// 	ch := make(chan int)
// 	// 相当于语义上定义了单向接口
// 	var send chan<- int = ch
// 	var recv <-chan int = ch

// 	// ./chan.go:15:6: cannot use send (type chan<- int) as type chan int in assignmen
// 	// 不能从单向往双向转
// 	// var x chan int = send
// 	// _ = x
// 	// 但unsafe的方式是可以转的
// 	var x chan int = *((*chan int)(unsafe.Pointer(&send)))
// 	_ = x

// 	go func() {
// 		defer wg.Done()
// 		for i := range recv {
// 			println(i)
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		defer close(ch)
// 		for i := 0; i < 3; i++ {
// 			send <- i
// 		}
// 	}()

// 	wg.Wait()

// }

// 二、 以工厂模式和goroutine交互
// func test() chan int {
// 	c := make(chan int)
// 	go func() {
// 		c <- 1
// 		c <- 2
// 		c <- 3
// 	}()
// 	return c
// }

// func main() {
// 	c := test()
// 	for i := range c {
// 		println(i)
// 	}
// }

// 三、用通道做并发安全队列，id生成器
// 伪码，每次接收的时候就触发一次
// var ids chan int // 同步通道

// func id() int {
// 	go func() { // 使用sync.Once让它只执行一次
// 		n := 0
// 		for {
// 			ids <- n
// 			n++
// 		}
// 	}()
// 	return <-ids
// }

// 四、实现信号量

func main() {
	sema := make(chan int, 3)
	for i :=0;i<100;i++{ //创建大量goroutine
		go func(){
			sema <- i	//使得同时只能运行sema个goroutine
			...			// 执行实际逻辑
			<- sema
		}()
	}
}
