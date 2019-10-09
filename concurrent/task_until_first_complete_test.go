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

func TestFirstResponse(t *testing.T) {
	t.Log("before goroutines:", runtime.NumGoroutine())
	t.Log(firstResponse())
	time.Sleep(time.Second)
	t.Log("after goroutines:", runtime.NumGoroutine())
}
