package concurrent

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

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

// 当第一个任务完成时返回
func TestFirstResponse(t *testing.T) {
	t.Log("before goroutines:", runtime.NumGoroutine())
	t.Log(firstResponse())
	time.Sleep(time.Second)
	t.Log("after goroutines:", runtime.NumGoroutine())
}

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

// 当所有任务完成时返回，除了使用csp的方式，也可以使用waitGroup来实现
func TestAllResponse(t *testing.T) {
	t.Log("before goroutines:", runtime.NumGoroutine())
	t.Log(allResponse())
	time.Sleep(time.Second)
	t.Log("after goroutines:", runtime.NumGoroutine())
}
