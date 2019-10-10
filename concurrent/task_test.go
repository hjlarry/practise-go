package concurrent

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
	"unsafe"
)

/* 一、使用sync.Once在多个协程中只执行一次任务 */
type Singleten struct{}

var single *Singleten

var once sync.Once

func getSingleton() *Singleten {
	once.Do(func() {
		single = new(Singleten)
		println("create success")
	})
	return single
}

func TestOnceTask(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			sin := getSingleton()
			println(unsafe.Pointer(sin))
		}()
	}
	time.Sleep(time.Second)
}

/* 二、多个任务执行时，当第一个任务完成时返回 */
func runTask(id int) string {
	time.Sleep(time.Millisecond * 50)
	return fmt.Sprintf("the result is from : %d", id)
}

func firstResponse() string {
	//ch := make(chan string)
	// 使用buffer chan 可使程序运行前后goroutines数量都是2
	ch := make(chan string, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- runTask(i)
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("before goroutines:", runtime.NumGoroutine())
	t.Log(firstResponse())
	time.Sleep(time.Second)
	t.Log("after goroutines:", runtime.NumGoroutine())
}

/* 三、多个任务执行时，当所有任务都完成时返回 */
func allResponse() string {
	ch := make(chan string, 4)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- runTask(i)
		}(i)
	}

	ret := ""
	for j := 0; j < 10; j++ {
		ret = ret + <-ch + "\n"
	}
	return ret
}

// 除了使用csp的方式，也可以使用waitGroup来实现
func TestAllResponse(t *testing.T) {
	t.Log("before goroutines:", runtime.NumGoroutine())
	t.Log(allResponse())
	time.Sleep(time.Second)
	t.Log("after goroutines:", runtime.NumGoroutine())
}

func isCancelled(cancelChan chan int) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

/* 四、通过close channel取消任务 */
// 只能取消一个任务
func cancel1(cancelChan chan int) {
	cancelChan <- 0
}

// 可取消全部任务
func cancel2(cancelChan chan int) {
	close(cancelChan)
}

func TestTaskClose(t *testing.T) {
	cancelChan := make(chan int)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan int) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			println(i, "cancelled")
		}(i, cancelChan)
	}
	// cancel1(cancelChan)
	cancel2(cancelChan)

	time.Sleep(time.Second)
}

/* 五、通过context取消任务 */
// 使用context取消任务，则可以取消该任务的子任务
func isCancelledByCtx(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestTaskCtxClose(t *testing.T) {
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
