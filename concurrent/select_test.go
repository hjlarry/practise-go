package concurrent

import (
	"fmt"
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

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService2():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}

/* 二、是否使用select的写法比较 */
// 每过一会产生一个数
func send4(ch chan int, gap time.Duration) {
	i := 0
	for {
		i++
		ch <- i
		time.Sleep(gap)
	}
}

// 将多个原通道内容拷贝到单一目标通道
func collect4(source chan int, target chan int) {
	for v := range source {
		target <- v
	}
}

// 从目标通道消费数据
func recv4(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

// 不使用select的写法
func TestChannel3(t *testing.T) {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var ch3 = make(chan int)
	go send4(ch1, time.Second)
	go send4(ch2, 2*time.Second)
	go collect4(ch1, ch3)
	go collect4(ch2, ch3)
	recv4(ch3)
}

func recv5(ch1 chan int, ch2 chan int) {
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
func TestSelect2(t *testing.T) {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go send4(ch1, time.Second)
	go send4(ch2, 2*time.Second)
	recv5(ch1, ch2)
}
