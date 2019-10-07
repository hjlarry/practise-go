package concurrent

import (
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "done"
}

func otherTask() {
	println("some work start...")
	time.Sleep(time.Millisecond * 100)
	println("some work end")
}

func TestService(t *testing.T) {
	println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string)
	// 带buffer的chan
	// retCh := make(chan string, 1)
	go func() {
		ret := service()
		println("returned result")
		retCh <- ret
		println("service exit") //执行顺序会和使用哪种chan相关
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	println(<-retCh)
	time.Sleep(time.Second)
}
