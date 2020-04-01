package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// 一、channel的基本使用
//func main() {
//	var ch = make(chan int, 4)
//	ch <- 10
//	ch <- 20
//	close(ch)
//
//	for value := range ch {
//		fmt.Printf("Receive %d \n", value)
//	}
//	// 通道变量本身就是指针，可用相等操作符判断
//	var a, b chan int = make(chan int, 3), make(chan int)
//	var c chan bool
//	fmt.Println(a == b)
//	fmt.Println(c == nil)
//	fmt.Printf("%p, %d \n", a, unsafe.Sizeof(a))
//
//	// cap返回缓冲区大小，据此可判断通道是同步还是异步的
//	fmt.Println("a:", len(a), cap(a))
//	a <- 1
//	a <- 2
//	fmt.Println("a:", len(a), cap(a))
//	fmt.Println("b:", len(b), cap(b))
//}

// 二、channel的关闭
//func dataProducer(ch chan int, wg *sync.WaitGroup) {
//	go func() {
//		for i := 0; i < 10; i++ {
//			ch <- i
//		}
//		close(ch)
//		wg.Done()
//	}()
//}
//
//func dataReceiver(ch chan int, wg *sync.WaitGroup) {
//	go func() {
//		for {
//			// data := <-ch
//			// channel关闭后使用上面的语句会获取到零值，使用下面的语句可以判断channel的状态
//			// 向关闭的channel发送数据会抛出异常
//			if data, ok := <-ch; ok {
//				println(data)
//			} else {
//				break
//			}
//		}
//		wg.Done()
//	}()
//}
//
//func main() {
//	ch := make(chan int)
//	var wg sync.WaitGroup
//	wg.Add(1)
//	dataProducer(ch, &wg)
//	wg.Add(1)
//	dataReceiver(ch, &wg)
//	wg.Add(1)
//	dataReceiver(ch, &wg)
//	wg.Wait()
//}

// 三、channel的收发
// 1. ok-idom模式
//func main() {
//	done := make(chan struct{})
//	c := make(chan int)
//
//	go func() {
//		defer close(done)
//		for {
//			x, ok := <-c
//			if !ok { // 据此判断通道的关闭
//				return
//			}
//			fmt.Println(x)
//		}
//	}()
//	c <- 3
//	c <- 2
//	c <- 1
//	close(c)
//	<-done
//}

// 2. 通过循环
//func main() {
//	done := make(chan struct{})
//	c := make(chan int)
//
//	go func() {
//		defer close(done)
//		for x := range c { //循环获取消息，直到通道关闭
//			fmt.Println(x)
//		}
//	}()
//	c <- 3
//	c <- 2
//	c <- 1
//	close(c)
//	<-done
//}

// 四、 单向通道
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

// 五、 以工厂模式和goroutine交互
//func test() chan int {
//	c := make(chan int)
//	go func() {
//		c <- 1
//		c <- 2
//		c <- 3
//	}()
//	return c
//}
//
//func main() {
//	c := test()
//	for i := range c {
//		println(i)
//	}
//}

// 六、用通道做并发安全队列，id生成器
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

// 七、实现信号量
//func main() {
//	runtime.GOMAXPROCS(4)
//	var wg sync.WaitGroup
//	sema := make(chan struct{}, 3)
//	for i := 0; i < 10; i++ { //创建大量goroutine
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			sema <- struct{}{} //获取信号
//			defer func() { // 执行实际逻辑
//				<-sema
//			}()
//			time.Sleep(time.Second)
//			fmt.Println(i, time.Now())
//		}(i)
//	}
//	wg.Wait()
//}

// 八、标准库time提供的timeout和tick channel的实现
//func main() {
//	go func() {
//		for {
//			select {
//			case <-time.After(time.Second * 5):
//				fmt.Println("timeout ...")
//				os.Exit(0)
//			}
//		}
//	}()
//
//	go func() {
//		tick := time.Tick(time.Second)
//		for {
//			select {
//			case <-tick:
//				fmt.Println(time.Now())
//			}
//		}
//	}()
//
//	<-(chan struct{})(nil) // 使用nil channel阻塞线程
//}

// 九、捕获INT和TERM信号，顺便实现一个简易的atexit函数
var exits = &struct {
	sync.RWMutex
	funcs   []func()
	signals chan os.Signal
}{}

func atexit(f func()) {
	exits.Lock()
	defer exits.Unlock()
	exits.funcs = append(exits.funcs, f)
}

func waitExit() {
	if exits.signals == nil {
		exits.signals = make(chan os.Signal)
		signal.Notify(exits.signals, syscall.SIGINT, syscall.SIGTERM)
	}
	exits.RLock()
	for _, f := range exits.funcs {
		defer f()
	} // 延迟调用按FILO顺序执行
	exits.RUnlock()
	//<-exits.signals
}

func main() {
	atexit(func() { println("exit1") })
	atexit(func() { println("exit2") })
	waitExit()
}