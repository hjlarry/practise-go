package main

import (
	"sync"
	"testing"
	"time"
)

// 一、并发操作
// 某任务对字典写操作时，其他任务就不能对该字典执行并发操作，否则会导致进程崩溃
// 可使用 go run-race test.go 来详细输出数据竞争的情况
func TestMapConcurrent(t *testing.T) {
	m := make(map[string]int)
	go func() {
		for {
			m["a"] += 1 //写
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			_ = m["b"] //读
			time.Sleep(time.Microsecond)
		}
	}()
	select {} //阻止进程退出
}

func TestMapConcurrentWithLock(t *testing.T) {
	var lock sync.RWMutex //使用读写锁以获得最佳性能
	m := make(map[string]int)
	go func() {
		for {
			lock.Lock() //注意锁的粒度
			m["a"] += 1
			lock.Unlock() //不能使用defer
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			lock.RLock()
			_ = m["b"]
			lock.RUnlock()
			time.Sleep(time.Microsecond)
		}
	}()
	select {} //阻止进程退出
}

// 二、初始化时填入cap的性能差异
func testNoCap() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	return m
}
func testCap() map[int]int {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	return m
}

// BenchmarkMap-4   	   20000	     97917 ns/op	   86810 B/op	      66 allocs/op
func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testNoCap()
	}
}

// BenchmarkCapMap-4   	   50000	     38309 ns/op	   41161 B/op	       7 allocs/op
func BenchmarkCapMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCap()
	}
}