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

type data struct {
	sync.Mutex
}

// func (d data) test(s string) 则锁失效
// 将Mutex作为匿名字段时，receiver需使用pointer
func (d *data) test(s string) {
	d.Lock()
	defer d.Unlock()
	for i := 0; i < 5; i++ {
		println(s, i)
		time.Sleep(time.Second)
	}
}

func TestLock(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data
	go func() {
		defer wg.Done()
		d.test("read")
	}()
	go func() {
		defer wg.Done()
		d.test("write")
	}()
	wg.Wait()
}

type cache struct {
	sync.RWMutex
	data []int
}

func (c *cache) count() int {
	c.Lock()
	n := len(c.data)
	c.Unlock()
	return n
}

func (c *cache) get() int {
	c.Lock()
	defer c.Unlock()
	var d int
	if n := c.count(); n > 0 { // count重复锁定，造成死锁
		d = c.data[0]
		c.data = c.data[1:]
	}
	return d
}

func TestDeadLock(t *testing.T) {
	c := cache{
		data: []int{1, 2, 3, 4},
	}
	t.Log(c.get())
}

/*
相关建议：
1. 对性能要求高时，应避免使用defer Unlock(这条存疑，defer在新版go中性能已经提升了，是否还需要避免?)
2. 读写并发时，用RWMutex性能会更好
3. 对单个数据读写保护，可尝试用原子操作
4. 执行严格测试，尽可能打开数据竞争检查
*/
