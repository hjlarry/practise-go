package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 练习 4.3： 重写reverse函数，使用数组指针代替slice。
func reverse2(ints *[5]int) {
	for i := 0; i < len(ints)/2; i++ {
		end := len(ints) - i - 1
		ints[i], ints[end] = ints[end], ints[i]
	}
}

// 练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。
func rotate(ints []int) {
	first := ints[0]
	copy(ints, ints[1:])
	ints[len(ints)-1] = first
}

func main() {
	a := [...]int{1, 2, 3, 4, 5} // 数组类型
	reverse(a[:])
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5} // 切片类型
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"

	b := [...]int{1, 2, 3, 4, 5}
	reverse2(&b)
	fmt.Println(b)

	c := []int{0, 1, 2, 3, 4, 5}
	rotate(c)
	rotate(c)
	fmt.Println(c)
}
