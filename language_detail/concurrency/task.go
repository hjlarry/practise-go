package main

import (
	"context"
	"time"
)

// 一、多个任务执行时，当第一个任务完成时返回
//func runTask(id int) string {
//	time.Sleep(time.Millisecond * 50)
//	return fmt.Sprintf("the result is from : %d", id)
//}
//
//func firstResponse() string {
//	//ch := make(chan string)
//	// 使用buffer chan 可使程序运行前后goroutines数量都是2
//	ch := make(chan string, 10)
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			ch <- runTask(i)
//		}(i)
//	}
//	return <-ch
//}
//
//func main() {
//	fmt.Println("before goroutines:", runtime.NumGoroutine())
//	fmt.Println(firstResponse())
//	time.Sleep(time.Second)
//	fmt.Println("after goroutines:", runtime.NumGoroutine())
//}

// 二、多个任务执行时，当所有任务都完成时返回
//func runTask(id int) string {
//	time.Sleep(time.Millisecond * 50)
//	return fmt.Sprintf("the result is from : %d", id)
//}
//
//func allResponse() string {
//	ch := make(chan string, 4)
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			ch <- runTask(i)
//		}(i)
//	}
//
//	ret := ""
//	for j := 0; j < 10; j++ {
//		ret = ret + <-ch + "\n"
//	}
//	return ret
//}
//
//func main() {
//	fmt.Println("before goroutines:", runtime.NumGoroutine())
//	fmt.Println(allResponse())
//	time.Sleep(time.Second)
//	fmt.Println("after goroutines:", runtime.NumGoroutine())
//}

// 三、通过close channel取消任务
//func isCancelled(cancelChan chan int) bool {
//	select {
//	case <-cancelChan:
//		return true
//	default:
//		return false
//	}
//}
//
//// 只能取消一个任务
//func cancel1(cancelChan chan int) {
//	cancelChan <- 0
//}
//
//// 可取消全部任务
//func cancel2(cancelChan chan int) {
//	close(cancelChan)
//}
//
//func main() {
//	cancelChan := make(chan int)
//	for i := 0; i < 5; i++ {
//		go func(i int, cancelCh chan int) {
//			for {
//				if isCancelled(cancelChan) {
//					break
//				}
//				time.Sleep(time.Millisecond * 5)
//			}
//			println(i, "cancelled")
//		}(i, cancelChan)
//	}
//	cancel1(cancelChan)
//	//cancel2(cancelChan)
//
//	time.Sleep(time.Second)
//}

// 四、通过context取消任务
// 使用context取消任务，则可以取消该任务的子任务
func isCancelledByCtx(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelledByCtx(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			println(i, "cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second)
}
