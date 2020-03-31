package main

import (
	"runtime"
	"sync"
)

// 一、Goexit
// 1. main中执行
// func main() {
// 	c := make(chan struct{})
// 	go func() {
// 		defer close(c)
// 		func() {
// 			// goexit能保证终止整个任务，而return只能终止当前的这个函数
// 			// runtime.Goexit()
// 			return
// 		}()
// 		println("done")
// 	}()
// 	<-c
// 	println("exit")
// }

// 2. goroutine中执行
// 在main中执行Goexit则会等待其他Goroutine执行完毕，然后让进城崩溃
//func main() {
//	go func() {
//		println("a")
//		println("b")
//		println("c")
//	}()
//	runtime.Goexit()
//}

// 二、使用Gosched可以暂停任务
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

// 三、任务调度观察
//func main() {
//	runtime.GOMAXPROCS(4)
//	for i := 0; i < 10000; i++ {
//		go func() {
//			for x := 0; x < 1e9; x++ {
//
//			}
//		}()
//	}
//	time.Sleep(time.Minute)
//}

/*
[ubuntu] ~/.mac/gocode $ go build runq.go
[ubuntu] ~/.mac/gocode $ GODEBUG=schedtrace=500 ./runq
SCHED 0ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=1 idlethreads=0 runqueue=0 [0 0]
SCHED 504ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9804 [9 175 9 0]
SCHED 1004ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9680 [8 174 8 127]
SCHED 1504ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9684 [7 173 7 126]
SCHED 2015ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9688 [6 172 6 125]
SCHED 2517ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9692 [5 171 5 124]
SCHED 3026ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9696 [4 170 4 123]
SCHED 3532ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9696 [4 170 4 123]
SCHED 4039ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9700 [3 169 3 122]
SCHED 4546ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9704 [2 168 2 121]
SCHED 5054ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9708 [1 167 1 120]
SCHED 5564ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9712 [0 166 0 119]
SCHED 6073ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9460 [127 165 127 118]
SCHED 6583ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9464 [126 164 126 117]
SCHED 7088ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9468 [125 163 125 116]
SCHED 7593ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=9468 [125 163 125 116]

每个P上的任务会逐渐趋于平衡
*/
