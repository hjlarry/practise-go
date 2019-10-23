package concurrent

import (
	"sync"
	"testing"
	"unsafe"
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

	// 通道变量本身就是指针，可用相等操作符判断
	var a, b chan int = make(chan int, 3), make(chan int)
	var c chan bool
	t.Log(a == b)
	t.Log(c == nil)
	t.Logf("%p, %d", a, unsafe.Sizeof(a))

	// cap返回缓冲区大小，据此可判断通道是同步还是异步的
	t.Log("a:", len(a), cap(a))
	a <- 1
	a <- 2
	t.Log("a:", len(a), cap(a))
	t.Log("b:", len(b), cap(b))

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

/* 三、测试channel的收发 */
// ok-idom模式
func TestChannelSend(t *testing.T) {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)
		for {
			x, ok := <-c
			if !ok { // 据此判断通道的关闭
				return
			}
			t.Log(x)
		}
	}()
	c <- 3
	c <- 2
	c <- 1
	close(c)
	<-done
}

func TestChannelSend2(t *testing.T) {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)
		for x := range c { //循环获取消息，直到通道关闭
			t.Log(x)
		}
	}()
	c <- 3
	c <- 2
	c <- 1
	close(c)
	<-done
}

/* 四、单向通道 */
// 通道默认双向，不区分发送和接收端，通常使用类型转换获取单向通道
func TestChannelSingle(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c

	go func() {
		defer wg.Done()
		for x := range recv {
			println(x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0; i < 3; i++ {
			send <- i
		}
	}()
	wg.Wait()
}

// 不能在单向通道做逆向操作,不能close接收通道
func TestChannelSingle1(t *testing.T) {
	c := make(chan int, 2)
	var send chan<- int = c
	var recv <-chan int = c
	// invalid operation: <-send (receive from send-only type chan<- int)
	// <-send
	// invalid operation: recv <- 1 (send to receive-only type <-chan int)
	// recv <- 1
	// invalid operation: close(recv) (cannot close receive-only channel)
	// close(recv)
	t.Log(send, recv)
}
