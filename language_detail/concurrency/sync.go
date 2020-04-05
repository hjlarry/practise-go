package main

import (
	"fmt"
	"sync"
)

//一、使用sync.Once在多个协程中只执行一次任务
//type Singleten struct{}
//
//var single *Singleten
//
//var once sync.Once
//
//func getSingleton() *Singleten {
//	once.Do(func() {
//		single = new(Singleten)
//		println("create success")
//	})
//	return single
//}
//
//func main() {
//	for i := 0; i < 10; i++ {
//		go func() {
//			sin := getSingleton()
//			println(unsafe.Pointer(sin))
//		}()
//	}
//	time.Sleep(time.Second)
//}

//二、使用sync.Pool
// 1. 单个goroutine
//func main() {
//	pool := sync.Pool{
//		New: func() interface{} {
//			fmt.Println("create a new obj")
//			return 10
//		},
//	}
//
//	fmt.Println(pool.Get())
//	fmt.Println(pool.Get())
//	pool.Put(3)
//	fmt.Println(pool.Get())
//	fmt.Println(pool.Get())
//	pool.Put(3)
//	runtime.GC() //GC 会清除sync.pool中缓存的对象
//	fmt.Println(pool.Get())
//}

// 2. 多个goroutine
// 创建对象有开销，但使用sync.Pool在协程间同步时背后加锁也有开销，需要权衡
func main() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new obj")
			return 100
		},
	}
	pool.Put(9)
	pool.Put(9)
	pool.Put("haha")

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(pool.Get())
			wg.Done()
		}()
	}
	wg.Wait()
}