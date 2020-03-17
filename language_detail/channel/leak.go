package main

// 资源泄漏的问题
import (
	"runtime"
	"time"
)

func test() {
	c := make(chan int)

	// 创建10个goroutine，每个从channel里接收数据
	for i := 0; i < 10; i++ {
		go func() {
			<-c
			// select {
			// case <-c:
			// default:
			}
		}()
	}
}

func main() {
	test()

	// 执行到这里，我们往往会认为上面的垃圾已经被回收了，但实际上不是
	for {
		time.Sleep(time.Second * 3)
		runtime.GC()
	}
}

/*

[ubuntu] ~/.mac/gocode $ go build leak.go
[ubuntu] ~/.mac/gocode $ GODEBUG="gctrace=1,schedtrace=1000,scheddetail=1" ./leak
SCHED 0ms: gomaxprocs=2 idleprocs=1 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
  P0: status=0 schedtick=1 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
  P1: status=1 schedtick=3 syscalltick=0 m=0 runqsize=0 gfreecnt=0
  M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 spinning=false blocked=false lockedg=-1
  M0: p=1 curg=1 mallocing=0 throwing=0 preemptoff= locks=2 dying=0 spinning=false blocked=false lockedg=-1
  G1: status=2(chan receive) m=0 lockedm=-1
  G2: status=4(force gc (idle)) m=-1 lockedm=-1
  G3: status=4(GC sweep wait) m=-1 lockedm=-1
  G4: status=4(GC scavenge wait) m=-1 lockedm=-1
  G17: status=1() m=-1 lockedm=-1
SCHED 1004ms: gomaxprocs=2 idleprocs=2 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
  P0: status=0 schedtick=1 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
  P1: status=0 schedtick=13 syscalltick=1 m=-1 runqsize=0 gfreecnt=0
  M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 spinning=false blocked=false lockedg=-1
  M0: p=-1 curg=27 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  G1: status=4(sleep) m=-1 lockedm=-1
  G2: status=4(force gc (idle)) m=-1 lockedm=-1
  G3: status=4(GC sweep wait) m=-1 lockedm=-1
  G4: status=4(GC scavenge wait) m=-1 lockedm=-1
  G17: status=4(chan receive) m=-1 lockedm=-1
  G18: status=4(chan receive) m=-1 lockedm=-1
  G19: status=4(chan receive) m=-1 lockedm=-1
  G20: status=4(chan receive) m=-1 lockedm=-1
  G21: status=4(chan receive) m=-1 lockedm=-1
  G22: status=4(chan receive) m=-1 lockedm=-1
  G23: status=4(chan receive) m=-1 lockedm=-1
  G24: status=4(chan receive) m=-1 lockedm=-1
  G25: status=4(chan receive) m=-1 lockedm=-1
  G26: status=4(chan receive) m=-1 lockedm=-1
  G27: status=3() m=0 lockedm=-1

我们发现这10个Goroutine一直处于接收通道数据的状态，就会造成内存泄漏

一种解决方案是把channel变为全局的，在运行完test以后关闭它，另一种方式是使用select如代码中的注释
*/
