package concurrent

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/* 一、channel的基本使用 */
func TestChannel(t *testing.T) {
	var ch = make(chan int, 4)
	ch <- 10
	ch <- 20
	close(ch)

	for value := range ch {
		t.Logf("Receive %d \n", value)
	}
}

/* 二、测试channel的关闭 */
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

/* 三、测试channel的读写 */
func send(ch chan int) {
	for {
		var value = rand.Intn(100)
		ch <- value
		fmt.Printf("Send %d \n", value)
	}
}

func recv(ch chan int) {
	for {
		value := <-ch
		fmt.Printf("Receive %d \n", value)
		time.Sleep(time.Second)
	}
}

func TestChannel2(t *testing.T) {
	var ch = make(chan int, 1)
	// 子协程循环读
	go recv(ch)
	// 主协程循环写
	send(ch)
}
