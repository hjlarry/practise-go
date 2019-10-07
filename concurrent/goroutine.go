package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

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

func testChannel() {
	var ch = make(chan int, 1)
	// 子协程循环读
	go recv(ch)
	// 主协程循环写
	send(ch)
}

func testChannel2() {
	var ch = make(chan int, 4)
	ch <- 10
	ch <- 20
	close(ch)
	// value := <-ch
	// fmt.Printf("Receive %d \n", value)
	// value = <-ch
	// fmt.Printf("Receive %d \n", value)
	// value = <-ch
	// fmt.Printf("Receive %d \n", value)

	for value := range ch {
		fmt.Printf("Receive %d \n", value)
	}
}

func send1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 计数值减一
	i := 0
	for i < 4 {
		i++
		ch <- i
	}
}

func recv1(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func testChannel3() {
	var ch = make(chan int, 4)
	var wg = new(sync.WaitGroup)
	wg.Add(2) // 增加计数值
	go send1(ch, wg)
	go send1(ch, wg)
	go recv1(ch)
	// Wait()会阻塞等待所有的写通道协程结束，因为计数值变为0，Wait()才会返回
	wg.Wait()
	close(ch)
	time.Sleep(time.Second)
}

// 没过一会产生一个数
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

func testChannel4() {
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

// testChannel4的语法糖形式
func testChannel5() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go send4(ch1, time.Second)
	go send4(ch2, 2*time.Second)
	recv5(ch1, ch2)
}

func send6(ch1 chan int, ch2 chan int) {
	i := 0
	for {
		i++
		select {
		case ch1 <- i:
			fmt.Println("send ch1", i)
		case ch2 <- i:
			fmt.Println("send ch2", i)
		// 决定读写操作是否阻塞
		default:
		}
	}
}

func recv6(ch chan int, gap time.Duration, name string) {
	for v := range ch {
		fmt.Printf("received %s %d \n", name, v)
		time.Sleep(gap)
	}
}

func testChannel6() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go recv6(ch1, time.Second, "ch1")
	go recv6(ch2, 2*time.Second, "ch2")
	send6(ch1, ch2)
}
