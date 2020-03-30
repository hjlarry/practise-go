package main

// import "fmt"

// type X int // 自定义类型
// type Y=int // 别名 

// func test(x int){}

// func main(){
// 	var a X = 1
// 	var b Y = 2

// 	// test(a)  //cannot use a (type X) as type int in argument to test
// 	test(b)
// 	fmt.Printf("%T, %T\n", a, b)  //main.X, int
// }


// go build -gcflags "-N -l -S" 2>alias.txt alias.go
type data struct{}

type Tester interface{
	// 接口会接收传入的类型，所以可以用接口观察编译器如何看待别名
	test()
}

type YYY = data
func(y YYY) test(){}

func main(){
	var a YYY
	// go.itab."".data,"".Tester(SB), AX
	// 说明编译器会把YYY都替换为data
	var t Tester = a  
	_ = t
}