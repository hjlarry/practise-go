package main

import "sync"

// 一、 是否用锁的对比
//func main() {
//	counter := 0
//	for i := 0; i < 5000; i++ {
//		go func() {
//			counter++
//		}()
//	}
//	time.Sleep(time.Second)
//	fmt.Println(counter)
//}

//func main() {
//	counter := 0
//	var mut sync.Mutex
//	for i := 0; i < 5000; i++ {
//		go func() {
//			defer func() {
//				mut.Unlock()
//			}()
//			mut.Lock()
//			counter++
//		}()
//	}
//	time.Sleep(time.Second)
//	fmt.Println(counter)
//}

// 二、锁为结构体的匿名字段时
//type data struct {
//	sync.Mutex
//}
//
//// 将Mutex作为结构体的匿名字段时，receiver需使用pointer，如果写成func (d data) test(s string)则锁会失效
//func (d *data) test(s string) {
//	d.Lock()
//	defer d.Unlock()
//	for i := 0; i < 5; i++ {
//		println(s, i)
//		time.Sleep(time.Second)
//	}
//}
//
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(2)
//	var d data
//	go func() {
//		defer wg.Done()
//		d.test("read")
//	}()
//	go func() {
//		defer wg.Done()
//		d.test("write")
//	}()
//	wg.Wait()
//}

// 三、死锁
var m sync.Mutex

func A() {
	m.Lock()
	defer m.Unlock()
	B()
}

// 会造成死锁，Go是不支持递归锁的(即可重入锁)
func B() {
	m.Lock()
	defer m.Unlock()
	println("hello")
}
func main() {
	A()
}

/*

相关建议：
1. 对性能要求高时，应避免使用defer Unlock
2. 读写并发时，用RWMutex性能会更好
3. 对单个数据读写保护，可尝试用原子操作
4. 执行严格测试，尽可能打开数据竞争检查
*/
