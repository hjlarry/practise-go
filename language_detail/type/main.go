package main

/*
一、同名遮蔽示例
var x = 100

func main() {
	x := x               // 0x18中的x = main.x
	if x := 20; x > 10 { // 0x10中的x赋值20
		println(x) // 打印0x10中的x
		x := x     // 0x8中的x赋值0x10中的x
		println(x) // 打印0x8中的x
	}
	println(x) // 打印0x18中的x
}
*/

/*
二、退化赋值示例
反汇编技巧：编译时使用默认的优化模式，但源码中针对某个方法定义禁止内联，这样这个方法就成了一个标志，方便查找和阅读


//go:noinline
func test() (int, int) {
	return 1, 2
}

func main() {
	a, x := test()
	println(a, x)
	b, x := test()
	println(b, x)
}
*/

/*
三、多变量赋值示例
反汇编证明多变量赋值和分别赋值是一样的速度


//go:noinline
//go:nosplit
func test() (int, int) {
	a, b := 1, 2
	a, b = b+1, a+2
	return a, b
}

func main() {
	a, b := test()
	println(a, b)
}
*/

/*
四、 动态修改字符串变量
go build -ldflags "-X main.BuildTime=$(date +'%Y.%m.%d')"


var BuildTime string

func main() {
	println(BuildTime)
}
*/

/*
五、 空标识符的本质

//go:noinline
//go:nosplit
func test() (int, int) {
	return 1, 2
}
func main() {
	x := 100
	_ = x
	a, _ := test()
	println(a)
}
*/

/*
六、 空标识符的使用场景
检测X是否实现了Xer接口中的方法

[ubuntu] ~/.mac/gocode $ go build -gcflags "-N -l"
# _/root/.mac/gocode
./main.go:92:5: cannot use X(0) (type X) as type Xer in assignment:
	X does not implement Xer (missing A method)



type Xer interface {
	A()
}

type X int

// func (x X) A() {}

var _ Xer = X(0)

func main() {

}
*/

/*
七、 常量


const x, s = 1, "abc"

func main() {
	const a = 1
	{
		const a = "abc"
		println(a)
	}
	println(a)
	const x int32 = 100
	const s uintptr = unsafe.Sizeof(0)
	const n int = len("abc")
	const (
		c int = 1 * int(unsafe.Sizeof("abc"))
		b
	)
}
*/

/*
八、 枚举
*/

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

type color byte

const (
	black color = iota
	red
	blue
)

func test(c color) {
	println(c)
}

func main() {
	println(KB, MB, GB)

	test(red)
	test(100) // 100是字面量，若未超过类型byte的范围，会被隐式转换为相应的类型

	x := 2
	// test(x) // x则被限制传入枚举类型，说明不允许对变量做隐式转换
	test(color(x)) // 显示转换合法
}
