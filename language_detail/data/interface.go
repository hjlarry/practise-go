package main

import (
	"fmt"
)

/* 一、 简单的接口示例 */
// 接口一般用er结尾命名
//type tester interface {
//	test()
//	string() string
//}
//type dd struct{}
//
//func (*dd) test() {}
//func (dd) string() string {
//	return "ss"
//}
//
//func main() {
//	var d dd
//	// var tt tester = d  报错：missing method test，编译器根据方法集判断是否实现了接口
//	var tt tester = &d
//	tt.test()
//	fmt.Println(tt.string())
//
//	// 接口比较
//	var t1, t2 interface{}
//	fmt.Println(t1 == t2)
//
//	t1, t2 = 100, 100
//	fmt.Println(t1 == t2)
//
//	// 报错，必须实现接口的类型 支持比较才能比较
//	t1, t2 = map[int]int{}, map[int]int{}
//	fmt.Println(t1 == t2)
//}

// --------------------------------------------------------------------------
// 二、接口反汇编
// type Xer interface {
// 	MethodA(a int)
// 	MethodB()
// }

// type X int

// func (x X) MethodA(b int) {}
// func (x X) MethodB()      {}
// func main() {
// 	var o X = 100
// 	var e Xer = o
// 	e.MethodB()
// }

/*
go build -gcflags "-l -N -S" 2>a.txt interface.go

1.数据的方法地址在编译期写入RODATA中
go.itab."".X,"".Xer SRODATA dupok size=40
	rel 0+8 t=1 type."".Xer+0
	rel 8+8 t=1 type."".X+0
	rel 24+8 t=1 "".(*X).MethodA+0
	rel 32+8 t=1 "".(*X).MethodB+0

结合iface的数据结构
type iface struct {
	tab  *itab
	data unsafe.Pointer
}

以及itab的数据结构
type itab struct {
	inter *interfacetype
	_type *_type
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}

go build -gcflags "-l -N" interface.go && go tool objdump -s "main\.main" interface
  interface.go:12	0x4523df		4883ec38		SUBQ $0x38, SP
  interface.go:12	0x4523e3		48896c2430		MOVQ BP, 0x30(SP)
  interface.go:12	0x4523e8		488d6c2430		LEAQ 0x30(SP), BP
  interface.go:13	0x4523ed		48c744241064000000	MOVQ $0x64, 0x10(SP)
  interface.go:14	0x4523f6		48c7042464000000	MOVQ $0x64, 0(SP)
  interface.go:14	0x4523fe		e8bd63fbff		CALL runtime.convT64(SB)
  interface.go:14	0x452403		488b442408		MOVQ 0x8(SP), AX
  interface.go:14	0x452408		4889442418		MOVQ AX, 0x18(SP)
  interface.go:14	0x45240d		488d0d8c020300		LEAQ go.itab.main.X,main.Xer(SB), CX
  interface.go:14	0x452414		48894c2420		MOVQ CX, 0x20(SP)
  interface.go:14	0x452419		4889442428		MOVQ AX, 0x28(SP)
  interface.go:15	0x45241e		8401			TESTB AL, 0(CX)
  interface.go:15	0x452420		488b0d99020300		MOVQ go.itab.main.X,main.Xer+32(SB), CX
  interface.go:15	0x452427		48890424		MOVQ AX, 0(SP)
  interface.go:15	0x45242b		ffd1			CALL CX

2.通过CALL runtime.convT64(SB)把数据100转换为指针地址时有复制行为，且大概率进行了堆上的内存分配
3.通过CX和AX进行调用
*/

// --------------------------------------------------------------------------
// 三、编译器根据方法集判断接口实现
// type X int

// func (x *X) test() {}

// type Xer interface {
// 	test()
// }

// func main() {
// 	var o X = 100
// 	var e Xer = o
// 	e.test()
// }

// ./interface.go:76:6: cannot use o (type X) as type Xer in assignment:
// 	X does not implement Xer (test method has pointer receiver)

// --------------------------------------------------------------------------
// 四、通过gdb观察
// type X int

// func (x *X) test() {}

// type Xer interface {
// 	test()
// }

// func main() {
// 	var o X = 100
// 	var e Xer = &o

// 	var f interface{} = &o
// 	println(e)
// 	println(f)
// }

/*
go build -gcflags "-l -N" interface.go && gdb interface
b 98

(gdb) info locals
o = 100
f = {
  _type = 0x461f00,
  data = 0xc000030720
}
e = {
  tab = 0x482600 <X,main.Xer>,
  data = 0xc000030720
}

(gdb) p/x *e.tab
$4 = {
  inter = 0x4638e0,
  _type = 0x461f00,
  hash = 0x342eca44,
  _ = {0x0, 0x0, 0x0, 0x0},
  fun = {0x4523e0}
}

(gdb) p/x e.tab.fun[0]
$5 = 0x4523e0

(gdb) info symbol 0x4523e0
main.(*X).test in section .text of /root/.mac/gocode/interface
*/

// --------------------------------------------------------------------------
// 五、匿名接口
// type X int

// func (x *X) test() {}

// func main() {
// 	var o X
// 	// 支持匿名接口，有时并不需要提前创建一个接口类型
// 	var e interface {
// 		test()
// 	} = &o
// 	e.test()
// }

// --------------------------------------------------------------------------
// 六、接口嵌入
// 接口嵌入和结构体嵌入不同，相当于嵌入并展开，而结构体嵌入有层级关系
// type Aer interface {
// 	MethodA()
// }

// type Ber interface {
// 	Aer
// 	MethodB()
// }

// type X int

// func (X) MethodA()  {}
// func (*X) MethodB() {}

// func main() {
// 	var o X = 100
// 	var a Aer = o
// 	var b Ber = &o

// 	println(a, b)
// 	// 超集接口可以转换为子集接口
// 	var c Aer = b
// 	println(c)

// }

// --------------------------------------------------------------------------
// 七、 多态
//type fruitable interface {
//	eat()
//}
//
//type fruit struct {
//	name string
//	fruitable
//}
//
//func (f fruit) want() {
//	f.eat()
//}
//
//type apple struct{}
//
//func (a apple) eat() {
//	fmt.Println("eat apple")
//}
//
//type banana struct{}
//
//func (b banana) eat() {
//	fmt.Println("eat banana")
//}
//
//func eatInterface(f fruitable) {
//	f.eat()
//}
//
//// 通过接口模拟其他语言的多态
//func main() {
//	// 使用结构体的组合实现的多态
//	var app = fruit{"Apple", apple{}}
//	app.want()
//	var bana = fruit{"Banana", banana{}}
//	bana.want()
//	// 直接使用接口实现的多态
//	a := &apple{}
//	b := &banana{}
//	eatInterface(a)
//	eatInterface(b)
//}


// --------------------------------------------------------------------------
// 八、接口断言
// 1. 方式一
// type Xer interface {
// 	A()
// }
// type X int

// func (X) A() {}
// func main() {
// 	var o X
// 	var e Xer = o
// 	a, ok := e.(X)
// 	println(a, ok)
// }

// --------------------------------------------------------------------------
// 2. 方式二
// func main() {
// 	var x interface{} = func(x int) string {
// 		return fmt.Sprintf("d:%d", x)
// 	}
// 	switch v := x.(type) {
// 	case nil:
// 		println("nil")
// 	case *int:
// 		println(*v)
// 	case func(int) string:
// 		println(v(100))
// 	case fmt.Stringer:
// 		fmt.Println(v)
// 	default:
// 		println("unknown")
// 	}
// }

// --------------------------------------------------------------------------
// 九、判断变量是否为nil
// func main() {
// 	var o *int = nil
// 	var a interface{} = o // {type: *int, data:nil}
// 	var b interface{}     // {type: nil, data:nil}
// 	println(a == nil, b == nil)

// 	// 那么只能通过反射来判断a是否为nil
// 	v := reflect.ValueOf(a)
// 	if v.IsValid() {
// 		println(v.IsNil())
// 	}
// }

// --------------------------------------------------------------------------
// 十、通过hack的方式修改接口
//type X struct {
//	o int
//}
//
//func main() {
//	b := X{100}
//	var e interface{} = b
//	// 可以通过类型转换后读取该字段
//	println(e.(X).o)
//	// 但不可写，因为接口的data是私有的，不可寻址的
//	// e.(X).o = 200
//
//	// 那么我们可以通过让接口的data持有一个指针
//	var f interface{} = &b
//	// 然后获取这个指针并修改数据
//	p := f.(*X)
//	(*p).o = 200
//	fmt.Println(b)
//	// 简写为
//	f.(*X).o = 300
//	fmt.Println(b)
//}

// --------------------------------------------------------------------------
// 十一、定义函数类型，让相同签名的函数自动实现某接口
type FuncString func() string
func(f FuncString) String() string{
	return f()
}

func main(){
	var tt fmt.Stringer = FuncString(func() string{
		// 转换类型，使其实现Stringer接口
		return "hello world"
	})
	fmt.Println(tt)
}