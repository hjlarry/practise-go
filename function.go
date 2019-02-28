package main

import "fmt"

func add1(a int) int {
	a = a + 1
	return a
}

func add2(a *int) int {
	*a = *a + 1
	return *a
}

// TestAdd ... 传值和传指针的区别
func TestAdd() {
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

// TestFunc ... 函数作为参数传入
func TestFunc() {
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
