package concurrent

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	t.Log(counter)
}

func TestCounterLock(t *testing.T) {
	counter := 0
	var mut sync.Mutex
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second)
	t.Log(counter)
}

func TestCounterWg(t *testing.T) {
	// waitgroup可替代time sleep
	counter := 0
	var mut sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1) // 标记开始一个事件
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Add(-1)
			// wg.Done() 标记结束一个事件
		}()
	}
	wg.Wait() // 阻塞至wg都结束
	t.Log(counter)
}
