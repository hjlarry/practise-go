package concurrent

import (
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			t.Log("create a new obj")
			return 10
		},
	}

	t.Log(pool.Get())
	t.Log(pool.Get())
	pool.Put(3)
	t.Log(pool.Get())
	t.Log(pool.Get())
	pool.Put(3)
	runtime.GC() //GC 会清除sync.pool中缓存的对象
	t.Log(pool.Get())
}

// 创建对象有开销，但使用sync.Pool在协程间同步时背后加锁也有开销，需要权衡
func TestSyncPoolMultiGoroutine(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			t.Log("create a new obj")
			return 100
		},
	}
	pool.Put(9)
	pool.Put(9)
	pool.Put("haha")

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			t.Log(pool.Get())
			wg.Done()
		}()
	}
	wg.Wait()
}
