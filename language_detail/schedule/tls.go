package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 通过参数来实现TLS(线程内本地存储)
func main() {
	var gs [5]struct {
		id     int
		result int
		// 填充无效内容，避免伪共享
		_ [64 - 8*2]byte
	}

	for i := 0; i < len(gs); i++ {
		go func(x int) {
			gs[x].id = x
			gs[x].result = x + 100
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Printf("%v", gs)

}


// TLS
func TestGoroutine6(t *testing.T) {
	var wg sync.WaitGroup
	var gs [5]struct { // 用于实现类似TLS（局部存储）功能
		id     int
		result int
	}

	for i := 0; i < len(gs); i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			gs[id].id = id
			gs[id].result = (id + 1) * 100
		}(i)
	}
	wg.Wait()
	t.Logf("%+v", gs)
}
