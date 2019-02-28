package main

import (
	"errors"
	"fmt"
)

var global = 1

// Allglobal ... 首字母大写 则其他包能访问
var Allglobal = 10

// A ...
const A = 100

func main() {
	varFunc()
	testAdd()
	testFunc()
	testStruct()
}

func varFunc() {
	fmt.Println("hello world")
	var s1 int = 10
	var s2 = 10
	s3 := 10
	fmt.Println(s2)
	fmt.Println(s3)
	for index := 0; index < 4; index++ {
		fmt.Println(s1)
	}
	var local = 2
	fmt.Println(global)
	fmt.Println(local)

	// 有符号整数，可以表示正负
	var a int8 = 1  // 1 字节
	var b int16 = 2 // 2 字节
	var c int32 = 3 // 4 字节
	var d int64 = 4 // 8 字节
	fmt.Println(a, b, c, d)

	// 无符号整数，只能表示非负数
	var ua uint8 = 1
	var ub uint16 = 2
	var uc uint32 = 3
	var ud uint64 = 4
	fmt.Println(ua, ub, uc, ud)

	// int 类型，在32位机器上占4个字节，在64位机器上占8个字节
	var e int = 5
	var ue uint = 5
	fmt.Println(e, ue)

	// bool 类型
	var f bool = true
	fmt.Println(f)

	// 字节类型
	var j byte = 'a'
	fmt.Println(j)

	// 字符串类型
	var g string = "abcdefg"
	fmt.Println(g)

	// 浮点数
	var h float32 = 3.14
	var i float64 = 3.141592653
	fmt.Println(h, i)

	var a_array [9]int
	fmt.Println(a_array)

	var aa = [10]int{1, 3, 4}
	var bb [10]int = [10]int{1, 3, 4}
	cc := [10]int{1, 3, 4}
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)

	var seq [9]int
	for i := 0; i < len(seq); i++ {
		seq[i] = i * i
	}
	fmt.Println(seq)

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
}

func add1(a int) int {
	a = a + 1
	return a
}

func add2(a *int) int {
	*a = *a + 1
	return *a
}

// 传值和传指针的区别
func testAdd() {
	x := 3
	fmt.Println("x=", x)
	x1 := add1(x)
	fmt.Println("x1=", x1)
	fmt.Println("x=", x)

	y := 3
	fmt.Println("y=", y)
	y1 := add2(&y)
	fmt.Println("y1=", y1)
	fmt.Println("y=", y)
}

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

// 函数作为参数传入
func testFunc() {
	slice := []int{1, 3, 4, 6, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("odd = ", odd)
	even := filter(slice, isEven)
	fmt.Println("even = ", even)
}

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

func testStruct() {
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
