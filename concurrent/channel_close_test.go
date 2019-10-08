package concurrent

import (
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			// data := <-ch
			// channel关闭后使用上面的语句会获取到零值，使用下面的语句可以判断channel的状态
			// 向关闭的channel发送数据会抛出异常
			if data, ok := <-ch; ok {
				println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
