package concurrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/* 一、使用select获取通道数据 */
func service2() string {
	// 调整时间可得到超时的case
	time.Sleep(time.Millisecond * 50)
	return "done"
}

func AsyncService2() chan string {
	retCh := make(chan string)
	go func() {
		ret := service2()
		retCh <- ret
	}()
	return retCh
}

func TestSelectBasic(t *testing.T) {
	select {
	case ret := <-AsyncService2():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}

/* 二、是否使用select的写法比较 */
// 每过一会产生一个数
func send(ch chan int, gap time.Duration) {
	i := 0
	for {
		i++
		ch <- i
		time.Sleep(gap)
	}
}

// 将多个原通道内容拷贝到单一目标通道
func collect(source chan int, target chan int) {
	for v := range source {
		target <- v
	}
}

// 从目标通道消费数据
func recv(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

// 不使用select的写法
func TestWithoutSelect(t *testing.T) {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var ch3 = make(chan int)
	go send(ch1, time.Second)
	go send(ch2, 2*time.Second)
	go collect(ch1, ch3)
	go collect(ch2, ch3)
	recv(ch3)
}

func recvSelect(ch1 chan int, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Println("receved from ch1", v)
		case v := <-ch2:
			fmt.Println("receved from ch2", v)
		}
	}
}

// 使用select的写法，实际为TestChannel3的语法糖形式
func TestWithSelect(t *testing.T) {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go send(ch1, time.Second)
	go send(ch2, 2*time.Second)
	recvSelect(ch1, ch2)
}

/* 三、使用select随机选择通道发送接收 */
func TestSelect(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	a, b := make(chan int), make(chan int)
	go func() { // 接收端
		defer wg.Done()
		for {
			var (
				name string
				x    int
				ok   bool
			)
			select { // 随机从一个通道接收
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}
			if !ok { // 若任一通道关闭，则终止接收
				return
			}
			println(name, x)
		}
	}()

	go func() { // 发送端
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 10; i++ {
			select { // 随机从一个通道发送
			case a <- i:
			case b <- i * 10:
			}
		}
	}()
	wg.Wait()
}

// 可将已完成的通道设为nil，这样它会被阻塞不再被select选中，可等待全部通道消息处理结束
func TestSelect2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	a, b := make(chan int), make(chan int)
	go func() { //接收端
		defer wg.Done()
		for {
			select {
			case x, ok := <-a:
				if !ok { // 如果通道关闭就设为nil
					a = nil
					break
				}
				println("a:", x)
			case x, ok := <-b:
				if !ok {
					b = nil
					break
				}
				println("b:", x)
			}

			if a == nil && b == nil { // 全部结束则退出循环
				return
			}
		}
	}()

	go func() { // 发送端a
		defer wg.Done()
		defer close(a)
		for i := 0; i < 3; i++ {
			a <- i
		}
	}()

	go func() { // 发送端b
		defer wg.Done()
		defer close(b)
		for i := 0; i < 5; i++ {
			b <- i * 10
		}
	}()
	wg.Wait()
}

// 同一通道，也会随机选择case执行
func TestSelect3(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)
	go func() { // 接收端
		defer wg.Done()
		for {
			var (
				x  int
				ok bool
			)
			select { // 随机从一个通道接收
			case x, ok = <-c:
				println("x1:", x)
			case x, ok = <-c:
				println("x2:", x)
			}
			if !ok {
				return
			}
		}
	}()

	go func() { // 发送端
		defer wg.Done()
		defer close(c)

		for i := 0; i < 10; i++ {
			select {
			case c <- i:
			case c <- i * 10:
			}
		}
	}()
	wg.Wait()
}

/* 四、select with default */
func TestSelect4(t *testing.T) {
	done := make(chan struct{})
	c := make(chan int)
	go func() {
		defer close(done)
		for {
			select {
			case x, ok := <-c:
				if !ok {
					return
				}
				fmt.Println("data:", x)
			default: // 可避免select阻塞
			}
			fmt.Println(time.Now())
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 5)
	c <- 100
	close(c)

	<-done
}

func TestSelect5(t *testing.T) {
	done := make(chan struct{})
	data := []chan int{
		make(chan int, 3),
	}
	go func() {
		defer close(done)
		for i := 0; i < 10; i++ {
			select {
			case data[len(data)-1] <- i: // 生产数据
			default: // 当前通道已满，生成新的缓存通道
				data = append(data, make(chan int, 3))
			}
		}
	}()
	<-done

	for i := 0; i < len(data); i++ { //显示所有数据
		c := data[i]
		close(c)
		for x := range c {
			println(x)
		}
	}

}
