package basic

import (
	"fmt"
	"testing"
)

// 一、 简单的接口示例
type GoPrograme struct {
	name string
}
type Programe interface {
	hello() string
}

func (g *GoPrograme) hello() string {
	return "hello go"
}
func TestInterface4(t *testing.T) {
	g := &GoPrograme{}
	var p Programe
	p = g
	t.Log(p.hello())
}

// 二、 复杂的接口示例
// Human ...
type Human struct {
	name  string
	age   int
	phone string
}

// Student ...
type Student struct {
	Human  //匿名字段Human
	school string
	loan   float32
}

// Employee ...
type Employee struct {
	Human   //匿名字段Human
	company string
	money   float32
}

// SayHi ... Human对象实现Sayhi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Sing ... Human对象实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

// Guzzle ... Human对象实现Guzzle方法
func (h Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// SayHi ... Employee重载Human的Sayhi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

// BorrowMoney ...Student实现BorrowMoney方法
func (s Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

// SpendSalary ... Employee实现SpendSalary方法
func (e Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

// Men ...
type Men interface {
	SayHi()
	Sing(lyrics string)
}

// YoungChap ...
type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

// ElderlyGent ...
type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

// TestInterface ...
func TestInterface(t *testing.T) {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}

// 三、 多态
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

// 四、空接口断言

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
