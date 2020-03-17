package main

import (
	"errors"
	"fmt"
)

// 一、切片定义
// 切片是对数组的描述，ptr描述从哪开始，len描述当前可读写的是哪一段，cap描述当前可以动态扩容的有多少
// 切片本身是只读的，我们操作的只是底层数组
//func main() {
//	var a []int               //仅定义，并未初始化
//	b := []int{}              //初始化表达式完成了全部创建过程
//	fmt.Println(a == nil, b == nil) //切片只能和nil去比较
//	fmt.Printf("a: %#v \n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))
//	fmt.Printf("b: %#v \n", (*reflect.SliceHeader)(unsafe.Pointer(&b)))
//	fmt.Printf("a size: %d \n", unsafe.Sizeof(a)) //虽未初始化，但分配了内存
//	fmt.Printf("b size: %d \n", unsafe.Sizeof(b))
//	fmt.Println("-------------------")
//
//	aa := [7]int{1, 2, 3, 4, 5, 6, 7}
//	s1 := aa[:]
//	s2 := s1[2:4]
//	s3 := s2[:5] // {3,4,5,6,7} 虽然s2只有两个元素，但我们要记住切片背后只是定义了一个结构体
//	fmt.Println(s1, len(s1), cap(s1))
//	fmt.Println(s2, len(s2), cap(s2))
//	fmt.Println(s3, len(s3), cap(s3))
//	fmt.Println("-------------------")
//
//	// 任一切片中的值变化会导致底层数组和其他引用该数组的值的变化
//	var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	ss3 := arr[2:5]
//	ss4 := arr[2:]
//	ss3[2] = 0
//	fmt.Println(ss3)
//	fmt.Println(ss4)
//	fmt.Println(arr)
//}

// 二、切片append
// func main() {
// 	var a [10]int

// 	s := a[:2:3] //{ptr:a[0], len:2, cap:3}
// 	s[0] = 100
// 	s[1] = 200

// 	s1 := append(s, 300)
// 	s2 := append(s1, 400) // 尽管原底层数组有空余，但s2并不知道，所以超出cap一定会复制原底层数组新建

// 	fmt.Printf("array: %p \n", &a)
// 	fmt.Println(a)
// 	fmt.Printf("%#v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s)))
// 	fmt.Printf("%#v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)))
// 	fmt.Printf("%#v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))
// }

// 三、切片copy
// func main() {
// 	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	s1 := a[1:3]
// 	s2 := a[6:]
// 	// 大的往小的copy，有多少个格子copy多少个
// 	fmt.Println(copy(s1, s2))
// 	fmt.Println(a)
// }


// 四、应避免长时间引用大数组
// go build slice.go && GODEBUG=gctrace=1 ./slice
//func test_slice() []byte {
//	s := make([]byte, 0, 100<<20)
//	s = append(s, 1, 2, 3, 4)
//	// 直接return 内存无法释放
//	// return s
//
//	s2 := make([]byte, len(s))
//	copy(s2, s)
//	return s2
//}
//
//func main() {
//	s := test_slice()
//	for {
//		fmt.Println(s)
//		runtime.GC()
//		time.Sleep(time.Second)
//	}
//}

// 五、stack
func main() {
	stack := make([]int, 0, 5)

	push := func(x int) error {
		n := len(stack)
		if n == cap(stack) {
			return errors.New("stack is full")
		}
		stack = stack[:n+1]
		stack[n] = x
		return nil
	}

	pop := func() (int, error) {
		n := len(stack)
		if n == 0 {
			return 0, errors.New("stack is empty")
		}
		x := stack[n-1]
		stack = stack[:n-1]
		return x, nil
	}

	for i := 0; i < 7; i++ {
		fmt.Printf("push %d: %v, %v \n", i, push(i), stack)
	}

	for i := 0; i < 7; i++ {
		x, err := pop()
		fmt.Printf("pop %d, %v, %v \n", x, err, stack)
	}
}