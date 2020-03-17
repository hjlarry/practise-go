package main

import (
	"fmt"
)

// 一、数组定义
//func main() {
//	var a [4]int              //元素自动初始化为零值
//	b := [4]int{2, 5}         //未提供初始化值的元素也初始化为零
//	c := [4]int{5, 3: 10}     //可指定索引位置3的值为10
//	d := [...]int{1, 2, 3}    //编译器按初始化值的数量来确定数组长度
//	e := [...]int{10, 3: 100} //也支持索引初始化，但需注意数组的长度
//	fmt.Println(a, b, c, d, e)
//
//	type user struct {
//		name string
//		age  int
//	}
//
//	f := [...]user{ // 复合类型，可省略元素初始化类型标签
//		{"Tome", 20},
//		{"Mary", 18},
//	}
//	fmt.Printf("%#v\n", f)
//
//	g := [...][2][3]int{ //仅第一维度允许使用...
//		{
//			{1, 2},
//			{3, 4},
//		},
//		{
//			{10, 20, 30},
//			{30, 40, 50},
//		},
//	}
//	fmt.Println(len(g), cap(g))
//	fmt.Println(len(g[1]), cap(g[1]))
//	fmt.Println(len(g[1][1]), cap(g[1][1]))
//
//	// 元素类型支持比较，则相同长度的数组也支持
//	h := [2]int{1, 2}
//	i := [2]int{0, 1}
//	fmt.Println(h == i)
//}


// 二、 数组指针和指针数组
//func main() {
//	x, y := 10, 20
//	// 指针数组：表示元素均为指针
//	a := [...]*int{&x, &y}
//	// 数组指针：存储数组地址的指针
//	p := &a
//	fmt.Printf("%T, %v", a, a)
//	fmt.Printf("%T, %v", p, p)
//
//	// 可获取任意元素的地址
//	fmt.Println(&a, &a[0], &a[1])
//
//	// 数组指针可直接操作元素
//	b := [...]int{10, 20}
//	q := &b
//	q[1] += 15
//	fmt.Println(q[1])
//}


// 三、 赋值传参
func testaa(x [2]int) {
	fmt.Printf("x: %p, %v\n", &x, x)
}

// 赋值和传参操作，会复制整个数组的数据
func main() {
	a := [2]int{10, 20}
	var b [2]int
	b = a

	fmt.Printf("a :%p, %v", &a, a)
	fmt.Printf("b :%p, %v", &b, b)

	testaa(a)
}

// 四、 大小端
//func main() {
//	// unsafe.Sizeof(a) 编译期可以计算出，所以能用来做数组长度的定义
//	a := 0x11223344 //小端： int [44 33 22 11 00 00 00 00]  通过gdb x/8xb &a 查看大小端
//	b := *(*[unsafe.Sizeof(a)]byte)(unsafe.Pointer(&a))
//	fmt.Printf("%x \n", b)
//}