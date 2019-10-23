package concurrent

import (
	"runtime"
	"testing"
	"time"
	"sync"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}


func TestGoroutine2(t *testing.T) {
	// 获取当前协程数量
	t.Log(runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		go func() {
			for {
				time.Sleep(time.Second)
			}
		}()
	}
	t.Log(runtime.NumGoroutine())

	// 读取默认的线程数
	t.Log(runtime.GOMAXPROCS(0))
	// 设置默认的线程数
	runtime.GOMAXPROCS(10)
	t.Log(runtime.GOMAXPROCS(0))
}


var c int
func counter() int{
	c++
	return c
}

// 说明gorutine会立即计算并复制执行的参数
func TestGoroutine3(t *testing.T){
	a:=100
	go func(x, y int){
		time.Sleep(time.Second) // 让main中的任务先执行
		t.Log("go:", x, y)
	}(a, counter())
	a += 100
	t.Log("main:", a, counter())
	time.Sleep(time.Second * 3) //等待go中的任务
}

// 进程退出时并不会等待gorutine结束，可以使用通道阻塞并发出退出信号
func TestGoroutine4(t *testing.T){
	exit := make(chan struct{})
	go func(){
		time.Sleep(time.Second*3)
		println("gorutine done")
		close(exit)
	}()
	println("main")
	<- exit
	println("main done")
}

// 等待多个任务结束，可以使用WaitGroup
func TestGoroutine5(t *testing.T){
	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		wg.Add(1)
		go func(id int){
			defer wg.Done()
			time.Sleep(time.Second)
			println("gorutine", id, "done")
		}(i)
	}
	println("main")
	wg.Wait()
	println("main done")
} 


// TLS
func TestGoroutine6(t *testing.T){
	var wg sync.WaitGroup
	var gs[5] struct{  // 用于实现类似TLS（局部存储）功能
		id int
		result int
	}

	for i:=0;i<len(gs);i++{
		wg.Add(1)
		go func(id int){
			defer wg.Done()
			gs[id].id = id
			gs[id].result = (id+1)*100
		}(i)
	}
	wg.Wait()
	t.Logf("%+v", gs)
}

// 使用Gosched暂停，释放线程去执行其他任务
func TestGoroutine7(t *testing.T){
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})
	go func(){  // 任务a
		defer close(exit)
		go func(){  // 任务b
			println("b")
		}()
		for i:=0;i<4;i++{
			println("a:", i)
			if i == 1{ // 让出当前线程，调度执行b
				runtime.Gosched()
			}
		}
	}()
	<- exit
}

// Goexit立即终止当前任务，但不能用在main.main
func TestGoroutine8(t *testing.T){
	exit := make(chan struct{})
	go func(){
		defer close(exit)
		defer println("a")
		func(){
			defer func(){
				println("b", recover()==nil)
			}()
			func(){
				println("c")
				runtime.Goexit()
				println("c done") //不执行
			}()
			println("b done")//不执行
		}()
		println("a done")	//不执行	
	}()
	<-exit
	println("main exit")//执行
}