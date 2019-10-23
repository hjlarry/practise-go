package concurrent

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"testing"
	"time"
)

type receiver struct {
	sync.WaitGroup
	data chan int
}

// 使用工厂方法将goroutine和通道绑定
func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}
	r.Add(1)
	go func() {
		defer r.Done()
		for x := range r.data {
			println("recv:", x)
		}
	}()
	return r
}

func TestReceiver(t *testing.T) {
	r := newReceiver()
	r.data <- 1
	r.data <- 2
	close(r.data)
	r.Wait()
}

// 使用通道实现信号量
func TestSemaphore(t *testing.T) {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	sem := make(chan struct{}, 2) // 最多允许两个并发同时执行
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- struct{}{} // 获取信号
			defer func() {    // 释放信号
				<-sem
			}()
			time.Sleep(time.Second * 2)
			fmt.Println(id, time.Now())
		}(i)
	}
	wg.Wait()
}

// 标准库time提供的timeout和tick channel的实现
func TestTick(t *testing.T) {
	go func() {
		for {
			select {
			case <-time.After(time.Second * 5):
				fmt.Println("timeout ...")
				os.Exit(0)
			}
		}
	}()

	go func() {
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				fmt.Println(time.Now())
			}
		}
	}()

	<-(chan struct{})(nil) // 使用nil channel阻塞线程
}

//捕获INT和TERM信号，顺便实现一个简易的atexit函数
var exits = &struct {
	sync.RWMutex
	funcs   []func()
	signals chan os.Signal
}{}

func atexit(f func()) {
	exits.Lock()
	defer exits.Unlock()
	exits.funcs = append(exits.funcs, f)
}

func waitExit() {
	if exits.signals == nil {
		exits.signals = make(chan os.Signal)
		signal.Notify(exits.signals, syscall.SIGINT, syscall.SIGTERM)
	}
	exits.RLock()
	for _, f := range exits.funcs {
		defer f()
	} // 延迟调用按FILO顺序执行
	exits.RUnlock()
	// <-exits.signals
}

func TestSignal(t *testing.T) {
	atexit(func() { println("exit1") })
	atexit(func() { println("exit2") })
	waitExit()
}

// 将发往通道的数据打包，可减少传输次数，提升性能
const (
	max     = 5000000
	block   = 500
	bufsize = 100
)

func testNoBlock() {
	done := make(chan struct{})
	c := make(chan int, bufsize)
	go func() {
		count := 0
		for x := range c {
			count += x
		}
		close(done)
	}()
	for i := 0; i < max; i++ {
		c <- i
	}
	close(c)
	<-done
}

func testBlock() {
	done := make(chan struct{})
	c := make(chan [block]int, bufsize)
	go func() {
		count := 0
		for a := range c {
			for _, x := range a {
				count += x
			}
		}
		close(done)
	}()
	for i := 0; i < max; i += block {
		var b [block]int
		for n := 0; n < block; n++ {
			b[n] = i + n
			if i+n == max-1 {
				break
			}
		}
		c <- b
	}
	close(c)
	<-done
}

// BenchmarkTestNoBlock-4   	       3	 421774198 ns/op	    1152 B/op	       2 allocs/op
func BenchmarkTestNoBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testNoBlock()
	}
}

// BenchmarkTestBlock-4   	     100	  12488588 ns/op	  401527 B/op	       2 allocs/op
func BenchmarkTestBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBlock()
	}
}
