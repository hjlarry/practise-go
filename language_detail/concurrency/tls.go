package main

import (
	"fmt"
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



