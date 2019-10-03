package basic

import (
	"fmt"
	"testing"
	"unsafe"
)

type person struct {
	name string
	age  int
}

type student struct {
	person
	special string
	skills
	int
}

type skills []string

type employee struct {
	person
	special string
	age     int
}

func TestStructInit(t *testing.T) {
	// 结构体初始化的三种方式
	e1 := person{name: "e1", age: 10}
	e2 := person{"e2", 11}
	e3 := new(person)
	e3.name = "e3"
	e3.age = 12
	t.Log(e1)
	t.Log(e2)
	t.Log(e3)
	// e3这种方式创建的为引用类型
	t.Logf("%T", e1)
	t.Logf("%T", e3)
}

// TestStruct ... 结构体的嵌套
func TestStruct(t *testing.T) {
	tom := person{"Tom", 25}
	t.Log(tom)
	t.Log(tom.age)

	mark := student{person: person{"Mark", 30}, special: "computer science"}
	t.Log(mark)
	t.Log(mark.name)
	mark.name = "haha"
	t.Log(mark.name)
	mark.person.age -= 2
	t.Log(mark.age)

	mark.skills = append(mark.skills, "golang", "python")
	t.Log(mark.skills)
	mark.int = 3
	t.Log(mark.int)

	larry := employee{person{"larry", 12}, "designer", 20}
	t.Log(larry.age)
	t.Log(larry.person.age)
}

// Circle ...
type Circle struct {
	x      int
	y      int
	Radius int
}

// TestStruct1 ... 结构体相互赋值本质上是浅拷贝，拷贝了其内部所有字段，而结构体指针的相互赋值仅拷贝了指针地址值
func TestStruct1(t *testing.T) {
	var c1 = Circle{Radius: 50}
	var c2 = c1
	t.Log(c1)
	t.Log(c2)
	c1.Radius = 100
	t.Log(c1)
	t.Log(c2)

	var c3 = &Circle{Radius: 50}
	var c4 = c3
	t.Log(c3)
	t.Log(c4)
	c3.Radius = 100
	t.Log(c3)
	t.Log(c4)
}

// ArrayStruct ...
type ArrayStruct struct {
	value [9]int
}

// SliceStruct ...
type SliceStruct struct {
	value []int
}

// TestStruct2 ... 数组占用的大小是数组容量，而切片只是头部容量
func TestStruct2(t *testing.T) {
	var as = ArrayStruct{[...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}}
	var ss = SliceStruct{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}}
	t.Log(unsafe.Sizeof(as), unsafe.Sizeof(ss))
}

func (p person) String1() string {
	fmt.Println("memory addr is ", unsafe.Pointer(&p.name))
	return "person func without *"
}

// 使用这种方式定义的方法不会造成内存的复制
func (p *person) String2() string {
	fmt.Println("memory addr is ", unsafe.Pointer(&p.name))
	return "person func with *"
}

func TestStructFunc(t *testing.T) {
	var p person
	fmt.Println("memory addr is ", unsafe.Pointer(&p.name))
	t.Log(p.String1())
	t.Log(p.String2())
}
