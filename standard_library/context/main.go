package main

import (
	"context"
	"fmt"
)

// 一、 通过上下文传值
// 查找值时实际上是递归查找，下例中b方法包装了新的值，最后就是新值
// func a(ctx context.Context) {
// 	ctx = context.WithValue(ctx, "aa", "hello world!")
// 	b(ctx)
// }

// func b(ctx context.Context) {
// 	ctx = context.WithValue(ctx, "aa", 1234)
// 	c(ctx)
// }

// func c(ctx context.Context) {
// 	time.Sleep(time.Second * 2)
// 	fmt.Println("c exec.", ctx.Value("aa"))
// }

// func main() {
// 	a(context.Background())
// }

// 二、 超时
// 用来判断某次链式调用中是否有超时
// func a(ctx context.Context) {
// 	b(ctx)
// }

// func b(ctx context.Context) {
// 	// 基于当前上下文定义超时时间
// 	ctx, _ = context.WithTimeout(ctx, time.Second*3)
// 	// 添加一个异步调用，正常来说不会阻塞当前程序，那么我们如何知道它超时了
// 	done := make(chan struct{})
// 	go func() {
// 		defer close(done)
// 		c(ctx)
// 	}()
// 	// 打印输出什么时候会超时
// 	fmt.Println(ctx.Deadline())
// 	// 判断是否超时
// 	select {
// 	// context信号先收到，说明超时之前未完成
// 	case <-ctx.Done():
// 		fmt.Println("timeout, ", ctx.Err())
// 	// 先收到通道消息，说明超时前已完成
// 	case <-done:
// 		fmt.Println("exec done")
// 	}
// }

// func c(ctx context.Context) {
// 	time.Sleep(time.Second * 4)
// 	fmt.Println("c exec.")
// }

// func main() {
// 	a(context.Background())
// }

// 三、取消
// 取消之后，会沿着链式调用向后传输，后续环节可以根据是否收到取消信号做不同逻辑
func a(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	// 假设提起调用取消操作，可能是在某个逻辑或某个goroutine中
	cancel()
	b(ctx)
}

func b(ctx context.Context) {
	c(ctx)
	_, ok := <-ctx.Done()
	fmt.Println("b:", ok)
}

func c(ctx context.Context) {
	// 模拟c能否检测到已经取消
	_, ok := <-ctx.Done()
	fmt.Println("c:", ok)
}

func main() {
	a(context.Background())
}
