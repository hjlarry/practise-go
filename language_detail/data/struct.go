package main

import (
	"fmt"
	"unsafe"
)

//一、 结构体初始化
//type person struct {
//	name string
//	age  int
//}
//
//func main() {
//	// 结构体初始化的三种方式
//	e1 := person{name: "e1", age: 10}
//	e2 := person{"e2", 11}
//	e3 := new(person)
//	e3.name = "e3"
//	e3.age = 12
//	fmt.Println(e1)
//	fmt.Println(e2)
//	fmt.Println(e3)
//	// e3这种方式创建的为引用类型
//	fmt.Printf("%T", e1)
//	fmt.Printf("%T", e3)
//
//	type node struct {
//		_    int //补位
//		id   int
//		next *node
//	}
//	n1 := node{id: 1}
//	n2 := node{id: 2, next: &n1}
//	fmt.Println(n1, n2)
//
//	// 直接定义匿名结构类型变量
//	u := struct {
//		name string
//		age  int
//	}{
//		name: "Tom",
//		age:  12,
//	}
//	type file struct {
//		name string
//		attr struct { //定义匿名结构类型字段
//			owner int
//			perm  int
//		}
//	}
//	f := file{
//		name: "test.dat",
//		// 错误初始化方式： missing type in composite literal
//		// attr: {
//		// 	owner: 1,
//		// 	perm:  0755,
//		// },
//	}
//	// 正确初始化方式
//	f.attr.owner = 1
//	f.attr.perm = 0755
//	fmt.Println(u, f)
//
//	type data struct {
//		*int //除接口指针和多级指针的任何命名类型都能做为匿名字段
//		string
//	}
//	x := 100
//	d := data{
//		int:    &x,
//		string: "abc",
//	}
//	fmt.Printf("%#v", d)
//}

// 二、空结构体
//func main() {
//	var a struct{}
//	var b [100]struct{} //空结构体做为数组元素类型，长度也为0
//	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))
//	// 尽管未分配数组内存，但仍然可以操作元素
//	s := b[:]
//	b[1] = struct{}{}
//	s[2] = struct{}{}
//	fmt.Println(s[3], len(s), cap(s))
//	// 实际上这类长度为0的对象通常指向runtime.zerobase变量
//	c := b[:]
//	d := [0]int{}
//	fmt.Printf("%p, %p, %p", &b[0], &c[0], &d)
//}

// 三、结构体相互赋值
//type Circle struct {
//	Radius int
//}
//
//// 结构体相互赋值本质上是浅拷贝，拷贝了其内部所有字段，而结构体指针的相互赋值仅拷贝了指针地址值
//func main() {
//	var c1 = Circle{Radius: 50}
//	var c2 = c1
//	fmt.Println(c1, c2)
//	c1.Radius = 100
//	fmt.Println(c1, c2)
//
//	var c3 = &Circle{Radius: 50}
//	var c4 = c3
//	fmt.Println(c3, c4)
//	c3.Radius = 100
//	fmt.Println(c3, c4)
//}

// 四、结构体内存分配
type ArrayStruct struct {
	value [9]int
}

type SliceStruct struct {
	value []int
}

func main() {
	// 数组占用的大小是数组容量，而切片只是头部容量
	var arrs = ArrayStruct{[...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}}
	var slis = SliceStruct{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}}
	fmt.Println(unsafe.Sizeof(arrs), unsafe.Sizeof(slis))

	// 内存总是一次性分配，各字段在相邻的地址空间按顺序排序
	type point struct {
		x, y, z int
	}
	type value struct {
		id   int
		name string
		data []byte
		next *value
		point
	}

	v := value{
		id:    1,
		name:  "test",
		data:  []byte{1, 2, 3, 4},
		point: point{x: 100, y: 200, z: 300},
	}

	s := `
		v: %p~%x, size: %d, align: %d
		field  address     offset   size
		----------+----------+-------+---------
		id        %p         %d      %d
		name      %p         %d      %d
		data      %p         %d      %d
		next      %p         %d      %d
		x         %p         %d      %d
		z         %p         %d      %d
	`
	fmt.Printf(s,
		&v, uintptr(unsafe.Pointer(&v))+unsafe.Sizeof(v), unsafe.Sizeof(v), unsafe.Alignof(v),
		&v.id, unsafe.Offsetof(v.id), unsafe.Sizeof(v.id),
		&v.name, unsafe.Offsetof(v.name), unsafe.Sizeof(v.name),
		&v.data, unsafe.Offsetof(v.data), unsafe.Sizeof(v.data),
		&v.next, unsafe.Offsetof(v.next), unsafe.Sizeof(v.next),
		&v.x, unsafe.Offsetof(v.x), unsafe.Sizeof(v.x),
		&v.z, unsafe.Offsetof(v.z), unsafe.Sizeof(v.z),
	)

	// 字段做对齐处理时，通常以所有字段中最长的基础类型宽度为标准
	v1 := struct {
		a byte
		b byte
		c int32 // 对齐宽度为4
	}{}
	v2 := struct {
		a byte
		b byte // 对齐宽度为1
	}{}
	v3 := struct {
		a byte
		b []int // 基础类型int，对齐宽度为8
		c byte
	}{}
	fmt.Printf("v1: %d, %d", unsafe.Alignof(v1), unsafe.Sizeof(v1))
	fmt.Printf("v2: %d, %d", unsafe.Alignof(v2), unsafe.Sizeof(v2))
	fmt.Printf("v3: %d, %d", unsafe.Alignof(v3), unsafe.Sizeof(v3))

	vv := struct {
		a struct{}
		b int
		c struct{} // 最后一个字段若是空结构类型字段，则会把它当做长度为1的类型做对齐处理，以免地址越界
	}{}

	ss := `
		v: %p~%x, size: %d, align: %d
		field  address     offset   size
		----------+----------+-------+---------
		a         %p         %d      %d
		b         %p         %d      %d
		c         %p         %d      %d
	`
	fmt.Printf(ss,
		&vv, uintptr(unsafe.Pointer(&vv))+unsafe.Sizeof(vv), unsafe.Sizeof(vv), unsafe.Alignof(vv),
		&vv.a, unsafe.Offsetof(vv.a), unsafe.Sizeof(vv.a),
		&vv.b, unsafe.Offsetof(vv.b), unsafe.Sizeof(vv.b),
		&vv.c, unsafe.Offsetof(vv.c), unsafe.Sizeof(vv.c),
	)
}
