package basic

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

type N int

// 不引用实例时，可只保留类名
func (N) test() {
	println("hellow")
}

func (n N) value() {
	n++
	fmt.Printf("v :%p, %v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("p :%p, %v\n", n, *n)
}

// 用T的情况：无需修改状态的小对象或固定值，引用类型、字符串、函数等指针包装对象
// 用*T的情况：要修改实例的状态、大对象、包含Mutex等同步字段、其他无法确定的情况
func TestMethod(t *testing.T) {
	var a N = 25
	a.test() //receiver会被复制
	a.value()
	a.pointer()
	fmt.Printf("a :%p, %v\n", &a, a)
}

// color
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type color byte
type box struct {
	width, height, depth float64
	color                color
}
type boxlist []box

func (b *box) volume() float64 {
	return b.width * b.height * b.depth
}

func (b *box) setColor(c color) {
	b.color = c
}

func (bl boxlist) biggestColor() color {
	v := 0.00
	k := color(WHITE)
	for _, b := range bl {
		if bv := b.volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl boxlist) PaintItBlack() {
	for i := range bl {
		bl[i].setColor(BLACK)
	}
}

func (c color) string() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

// TestMethod2 ...
func TestMethod2(t *testing.T) {
	boxes := boxlist{
		box{4, 4, 4, RED},
		box{10, 2, 10, BLUE},
		box{5, 5, 20, YELLOW},
		box{1, 4, 4, BLACK},
		box{4, 30, 4, WHITE},
	}
	t.Logf("we have %d boxes", len(boxes))
	t.Log("first box volumn:", boxes[0].volume(), "cm3")
	t.Log("last box color:", boxes[len(boxes)-1].color.string())
	t.Log("biggest:", boxes.biggestColor().string())

	t.Log("paint all to black:")
	boxes.PaintItBlack()
	t.Log("last box color:", boxes[len(boxes)-1].color.string())
}

type human struct {
	name string
	age  int
}

type student1 struct {
	human
	school string
}

type employee1 struct {
	human
	company string
}

func (h human) sayHi() {
	fmt.Println("this is human", h.name)
}

func (e employee1) sayHi() {
	fmt.Println("this is employee", e.name, e.company)
}

// TestMethod3 ... method继承和重写
func TestMethod3(t *testing.T) {
	mark := student1{human{"mark", 30}, "MIT"}
	sam := employee1{human{"sam", 30}, "Golang Inc."}

	mark.sayHi()
	sam.sayHi()
}

type data struct {
	sync.Mutex
	buf [1024]byte
}

func TestMethod4(t *testing.T) {
	d := data{}
	d.Lock() //可以直接调用匿名成员的方法，编译时会处理为sync.(*Mutex).Lock()的调用
	defer d.Unlock()
}

type S struct{}
type T struct {
	S
}

func (S) SVal()  {}
func (*S) SPtr() {}
func (T) TVal()  {}
func (*T) TPtr() {}

// 方法集仅影响接口实现和方法表达式转换，与通过实例或实例指针调用方法无关
func methodSet(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println(t)
	// NumMethod返回所有公开的方法
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		println(m.Name, m.Type)
	}
}

func TestMethodSet(t *testing.T) {
	var tt T
	methodSet(tt)
	println("-----")
	methodSet(&tt)
}

type M int

func (m M) test() {
	fmt.Printf("test.m: %p, %d\n", &m, m)
}

// 方法和函数一样也可以赋值给变量或作为参数传递，按引用方式，分为expression和value两种状态
func TestMethodTransfer(t *testing.T) {
	var m M = 26
	fmt.Printf("main.m: %p, %d\n", &m, m)
	// 通过类型引用的method expression会被还原为普通函数，调用时receiver作为参数传入
	f1 := M.test
	f1(m)
	// 类型是T或*T都可以，目标方法在其方法集中即可
	f2 := (*M).test
	f2(&m) //须按方法集中的签名传入正确的参数

	M.test(m)
	(*M).test(&m)

	p := &m
	m++
	// 基于指针或实例引用的method value依旧正常调用
	f3 := m.test //但当其赋值给变量或作为参数传递时，会立即复制该方法所需的receiver对象并与其绑定
	m++
	f4 := p.test
	m++
	fmt.Printf("main.m: %p, %d\n", p, m)
	f3()
	f4()
}
