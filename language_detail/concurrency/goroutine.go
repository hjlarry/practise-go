package main

import (
	"fmt"
	"time"
)

// 一、设置协程数量
//func main() {
//	// 获取当前协程数量
//	fmt.Println(runtime.NumGoroutine())
//	for i := 0; i < 10; i++ {
//		go func() {
//			for {
//				time.Sleep(time.Second)
//			}
//		}()
//	}
//	fmt.Println(runtime.NumGoroutine())
//
//	// 读取默认的线程数
//	fmt.Println(runtime.GOMAXPROCS(0))
//	// 设置默认的线程数
//	runtime.GOMAXPROCS(10)
//	fmt.Println(runtime.GOMAXPROCS(0))
//}

// 二、goroutine会立即计算并复制执行的参数
var c int

func counter() int {
	c++
	return c
}

func main() {
	a := 100
	go func(x, y int) {
		time.Sleep(time.Second) // 让main中的任务先执行
		fmt.Println("go:", x, y)
	}(a, counter())
	a += 100
	fmt.Println("main:", a, counter())
	time.Sleep(time.Second * 3) //等待go中的任务
}
