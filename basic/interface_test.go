package basic

import (
	"fmt"
	"testing"
)

/* 一、 简单的接口示例 */
// 接口一般用er结尾命名
type tester interface {
	test()
	string() string
}
type dd struct{}

func (*dd) test() {}
func (dd) string() string {
	return "ss"
}

func TestSimpleInterface(t *testing.T) {
	var d dd
	// var tt tester = d  报错：missing method test，编译器根据方法集判断是否实现了接口
	var tt tester = &d
	tt.test()
	t.Log(tt.string())
}

func TestEmptyInterface(t *testing.T) {
	var t1, t2 interface{}
	t.Log(t1 == t2)

	t1, t2 = 100, 100
	t.Log(t1 == t2)

	// 报错，必须实现接口的类型支持比较才能比较
	t1, t2 = map[int]int{}, map[int]int{}
	t.Log(t1 == t2)
}

type smaller interface {
	string() string
}

type bigger interface {
	smaller //嵌入其他接口
	test()
}

func qq(a smaller) {
	println(a.string())
}
func TestEmbedInterface(t *testing.T) {
	var d dd
	var big bigger = &d
	big.test()
	t.Log(big.string())

	qq(big)             // 可隐式转换为子集接口
	var s smaller = big // 超集转换为子集
	t.Log(s.string())
	// var t bigger = s   子集无法转换为超集,missing method test
}

type node struct {
	data interface { //匿名接口类型
		string() string
	}
}

func TestAnymousInterface(t *testing.T) {
	var tt interface { // 定义匿名接口变量
		string() string
	} = dd{}
	n := node{
		data: tt,
	}
	t.Log(n.data.string())
}

// 二、 多态
type fruitable interface {
	eat()
}

type fruit struct {
	name string
	fruitable
}

func (f fruit) want() {
	f.eat()
}

type apple struct{}

func (a apple) eat() {
	fmt.Println("eat apple")
}

type banana struct{}

func (b banana) eat() {
	fmt.Println("eat banana")
}

func eatInterface(f fruitable) {
	f.eat()
}

// TestInterface3 ... 通过接口模拟其他语言的多态
func TestInterface3(t *testing.T) {
	// 使用结构体的组合实现的多态
	var app = fruit{"Apple", apple{}}
	app.want()
	var bana = fruit{"Banana", banana{}}
	bana.want()
	// 直接使用接口实现的多态
	a := new(apple)
	b := &banana{}
	eatInterface(a)
	eatInterface(b)
}

// 三、空接口断言
func dosomething(any interface{}) {
	// if i, ok := any.(int); ok {
	// 	println("it`s interger:", i)
	// 	return
	// }
	switch v := any.(type) {
	case int:
		println("it`s interger:", v)
	case string:
		println("it`s string:", v)
	default:
		println("unknown type")
	}
	println(any)
	fmt.Printf("%T\n", any)

}

func TestInterface5(t *testing.T) {
	i := 1000
	dosomething(i)
	b := "abc"
	dosomething(b)
	p := &b
	dosomething(p)
}

type myint int

func (d myint) String() string{
	return fmt.Sprintf("data: %d", d)
}
func TestInterfaceConversion(t *testing.T){
	var d myint = 15
	var x interface{} = d

	if n, ok := x.(fmt.Stringer); ok{  // 转换为更具体的接口类型
		t.Log(n)
		t.Log(x)
	}

	if d2, ok:= x.(myint);ok{  // 转换为原始的类型
		t.Log(d2)
	}

	// 错误，interface conversion: basic.myint is not error: missing method Error [recovered]
	// ok-idom模式可避免引发panic
	// e:=x.(error)  
	// t.Log(e)

	// 使用switch在多种类型间做匹配
	var i interface{} = func(x int) string{
		return fmt.Sprintf("d: %d", x)
	}
	switch v:= i.(type){
	case nil:
		t.Log("nil")
	case *int:
		t.Log("*int")
	case func(int) string:
		t.Log(v(100))
	case fmt.Stringer:
		t.Log(v)
	default:
		t.Log("unknown")
	}
}

// 定义函数类型，让相同签名的函数自动实现某接口
type FuncString func() string
func(f FuncString) String() string{
	return f()
}

func TestFuncString(t *testing.T){
	var tt fmt.Stringer = FuncString(func() string{
		// 转换类型，使其实现Stringer接口
		return "hello world"
	})
	t.Log(tt)
}