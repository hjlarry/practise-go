package main

import (
	"errors"
	"fmt"
)

// 检测错误通过特定的一个变量
// var errTest = errors.New("test")

// func test() error {
// 	return errTest
// }
// func main() {
// 	err := test()
// 	if err == errTest {
// 	}
// }

// 通过自定义类型匹配的处理方式
// func main() {
// 	err := test()
// 	if _, ok := err.(ErrSome); ok {
// 	}
// 	switch err.(type) {
// 	case nil: // no error
// 	case ErrSome: // do something...
// 	default: // unknown error
// 	}
// }

// 由于变量错误对象可被修改，可以变向的使用常量错误值
// type Error string

// func (e Error) Error() string {
// 	return string(e)
// }

// const ErrEOF = Error("EOF")

// func test() error {
// 	return ErrEOF
// }

// func main() {
// 	err := test()
// 	println(err == Error("EOF"))
// }

// 通过对错误的包装，可以使其在不同的层级(上下文)呈现不同的内容
type OrderError struct {
	msg string
}

func (e *OrderError) Error() string {
	return e.msg
}

func main() {
	a := errors.New("a")
	b := fmt.Errorf("b_%w", a)
	c := fmt.Errorf("c_%w", b)
	fmt.Println(c)

	fmt.Println(errors.Unwrap(c) == b)
	fmt.Println(errors.Is(c, a)) //递归匹配

	x := &OrderError{"order"}
	y := fmt.Errorf("y_%w", x)
	z := fmt.Errorf("z_%w", y)
	fmt.Println(z)

	var x2 *OrderError
	if errors.As(z, &x2) { // 按类型递归查找，并填写指针（指针的指针）
		fmt.Println(x == x2)
	}
}
