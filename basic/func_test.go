package basic

import "testing"

func add1(a int) int {
	a = a + 1
	return a
}

func add2(a *int) int {
	*a = *a + 1
	return *a
}

// TestAdd ... 传值和传指针的区别
func TestAdd(t *testing.T) {
	x := 3
	t.Log("x=", x)
	x1 := add1(x)
	t.Log("x1=", x1)
	t.Log("x=", x)

	y := 3
	t.Log("y=", y)
	y1 := add2(&y)
	t.Log("y1=", y1)
	t.Log("y=", y)
}

type testInt func(int) bool

func isOdd(integer int) bool {
	return integer%2 != 0
}

func isEven(integer int) bool {
	return integer%2 == 0
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
func TestFunc(t *testing.T) {
	slice := []int{1, 3, 4, 6, 7}
	t.Log("slice = ", slice)
	odd := filter(slice, isOdd)
	t.Log("odd = ", odd)
	even := filter(slice, isEven)
	t.Log("even = ", even)
}

func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestMiltiVars(t *testing.T) {
	t.Log(sum(1, 2, 3, 4))
}

func clear() {
	println("clear some resource")
}

func TestDefer(t *testing.T) {
	defer clear()
	println("start")
	panic("some err")
}
