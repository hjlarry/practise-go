package basic

import (
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

// TestStruct ... 结构体
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
