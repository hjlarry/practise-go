package concurrent

import (
	"runtime"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}

func TestGoroutine2(t *testing.T) {
	// 获取当前协程数量
	t.Log(runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		go func() {
			for {
				time.Sleep(time.Second)
			}
		}()
	}
	t.Log(runtime.NumGoroutine())

	// 读取默认的线程数
	t.Log(runtime.GOMAXPROCS(0))
	// 设置默认的线程数
	runtime.GOMAXPROCS(10)
	t.Log(runtime.GOMAXPROCS(0))
}
