package concurrent

import (
	"testing"
	"time"
)

func service2() string {
	time.Sleep(time.Millisecond * 500)
	return "done"
}

func AsyncService2() chan string {
	retCh := make(chan string)
	go func() {
		ret := service2()
		println("returned result")
		retCh <- ret
		println("service exit") //执行顺序会和使用哪种chan相关
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService2():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
