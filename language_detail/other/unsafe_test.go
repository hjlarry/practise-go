package main

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type MyInt int

func TestUnsafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	//d := *(*MyInt())(unsafe.Pointer(&i)) invalid

	t.Log(unsafe.Pointer(&i))
	t.Log(f)

	// 合理的类型转换
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)
	t.Logf("%T", b[0])

}

// 原子类型操作
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}

	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		t.Log(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
