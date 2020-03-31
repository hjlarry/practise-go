package main

import (
	"runtime"
	"sync"
	"testing"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for x := 0; x < 10; x++ {
				println(id, ":", x)
				// 主动调度有时候会有用
				// runtime.Gosched()
			}
		}(i)
	}
	wg.Wait()
}

/*
默认情况:
[ubuntu] ~/.mac/gocode $ go run gosched.go
1 : 0
1 : 1
1 : 2
1 : 3
1 : 4
1 : 5
1 : 6
1 : 7
1 : 8
1 : 9
0 : 0
0 : 1
0 : 2
0 : 3
0 : 4
0 : 5
0 : 6
0 : 7
0 : 8
0 : 9

主动切换:
[ubuntu] ~/.mac/gocode $ go run gosched.go
1 : 0
0 : 0
1 : 1
0 : 1
1 : 2
0 : 2
1 : 3
0 : 3
1 : 4
0 : 4
1 : 5
0 : 5
1 : 6
0 : 6
1 : 7
0 : 7
1 : 8
0 : 8
1 : 9
0 : 9
*/


// 使用Gosched暂停，释放线程去执行其他任务
func TestGoroutine7(t *testing.T) {
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})
	go func() { // 任务a
		defer close(exit)
		go func() { // 任务b
			println("b")
		}()
		for i := 0; i < 4; i++ {
			println("a:", i)
			if i == 1 { // 让出当前线程，调度执行b
				runtime.Gosched()
			}
		}
	}()
	<-exit
}
