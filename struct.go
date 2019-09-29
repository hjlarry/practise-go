package main

import (
	"fmt"
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
func TestStruct() {
	tom := person{"Tom", 25}
	fmt.Println(tom)
	fmt.Println(tom.age)

	mark := student{person: person{"Mark", 30}, special: "computer science"}
	fmt.Println(mark)
	fmt.Println(mark.name)
	mark.name = "haha"
	fmt.Println(mark.name)
	mark.person.age -= 2
	fmt.Println(mark.age)

	mark.skills = append(mark.skills, "golang", "python")
	fmt.Println(mark.skills)
	mark.int = 3
	fmt.Println(mark.int)

	larry := employee{person{"larry", 12}, "designer", 20}
	fmt.Println(larry.age)
	fmt.Println(larry.person.age)
}

// Circle ...
type Circle struct {
	x      int
	y      int
	Radius int
}

// TestStruct1 ... 结构体相互赋值本质上是浅拷贝，拷贝了其内部所有字段，而结构体指针的相互赋值仅拷贝了指针地址值
func TestStruct1() {
	var c1 = Circle{Radius: 50}
	var c2 = c1
	fmt.Printf("%+v\n", c1)
	fmt.Printf("%+v\n", c2)
	c1.Radius = 100
	fmt.Printf("%+v\n", c1)
	fmt.Printf("%+v\n", c2)

	var c3 = &Circle{Radius: 50}
	var c4 = c3
	fmt.Printf("%+v\n", c3)
	fmt.Printf("%+v\n", c4)
	c3.Radius = 100
	fmt.Printf("%+v\n", c3)
	fmt.Printf("%+v\n", c4)
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
func TestStruct2() {
	var as = ArrayStruct{[...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}}
	var ss = SliceStruct{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}}
	fmt.Println(unsafe.Sizeof(as), unsafe.Sizeof(ss))
}
