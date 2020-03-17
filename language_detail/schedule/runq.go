package main

import "time"
import "runtime"

func main() {
	runtime.GOMAXPROCS(4)
	for i := 0; i < 10000; i++ {
		go func() {
			for x := 0; x < 1e9; x++ {

			}
		}()
	}
	time.Sleep(time.Minute)
}
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