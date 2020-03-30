package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"

	"github.com/pkg/errors"
)

func TestTypeAndValue(t *testing.T) {
	var f uint64 = 99
	t.Log(reflect.ValueOf(f), reflect.TypeOf(f))
	t.Log(reflect.ValueOf(f).Type())

	// Type表示其真实类型，Kind表示其底层类型
	type X int
	var a X = 100
	r := reflect.TypeOf(a)
	t.Log(r.Name(), r.Kind())

	// 可直接构造一些基础的复合类型
	b := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	c := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	t.Log(b)
	t.Log(c)

	// 传入对象应区分基类型和指针类型
	y := 100
	ty, tp := reflect.TypeOf(y), reflect.TypeOf(&y)
	t.Log(ty, tp)
	t.Log(ty.Kind(), tp.Kind())
	t.Log(ty == tp.Elem())

	// Elem返回指针、数组、切片、字典值、通道的基类型
	t.Log(reflect.TypeOf(map[string]int{}).Elem())
	t.Log(reflect.TypeOf([]int32{}).Elem())

}

/* 遍历结构体字段 */
type user struct {
	name string
	age  int
}
type manager struct {
	user
	title string
}

func TestStructField(t *testing.T) {
	var m manager
	tt := reflect.TypeOf(&m)
	if tt.Kind() == reflect.Ptr { // 遍历struct的字段，只能获取结构体指针的基类型
		tt = tt.Elem()
	}
	for i := 0; i < tt.NumField(); i++ {
		f := tt.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset)
		if f.Anonymous { // 若是匿名字段结构
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Println(af.Name, af.Type, af.Offset)
			}
		}
	}

	// 按名称查找，但不支持多级名称，如果同名遮蔽须通过匿名字段二次获取
	name, _ := tt.FieldByName("name")
	fmt.Println(name.Name, name.Type)
	// 按多级索引查找
	age := tt.FieldByIndex([]int{0, 1})
	fmt.Println(age.Name, age.Type)
}

/* 遍历方法集，也区分指针类型或基类型 */
type A int
type B struct {
	A
}

func (A) Av()  {}
func (*A) Ap() {}
func (B) Bv()  {}
func (*B) Bp() {}

func TestMethodSet(t *testing.T) {
	var b B
	tt := reflect.TypeOf(&b)
	ss := tt.Elem()
	fmt.Println(tt, ":")
	for i := 0; i < tt.NumMethod(); i++ {
		fmt.Println(" ", tt.Method(i))
	}
	fmt.Println(ss, ":")
	for i := 0; i < ss.NumMethod(); i++ {
		fmt.Println(" ", ss.Method(i))
	}
}

func checkType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int32, reflect.Int64:
		println("it`s int")
	case reflect.Float32, reflect.Float64:
		println("it`s float")
	default:
		println("unknown other")
	}
}

func TestCheckType(t *testing.T) {
	var f float64 = 20
	checkType(f)
}

func TestDeepEqual(t *testing.T) {
	a1 := []int{1, 2, 3}
	a2 := []int{1, 2, 3}
	a3 := []int{2, 3, 1}
	//t.Log(a1 == a2)   slice不能直接比较，map也是
	t.Log(reflect.DeepEqual(a1, a2))
	t.Log(reflect.DeepEqual(a1, a3))

	b1 := map[string]string{"name": "mike", "age": "40"}
	b2 := map[string]string{"name": "mike", "age": "40"}
	//t.Log(b1 == b2)
	t.Log(reflect.DeepEqual(b1, b2))

	c1 := Customer{"1", "mike", 40}
	c2 := Customer{"1", "mike", 40}
	t.Log(c1 == c2)
	t.Log(reflect.DeepEqual(c1, c2))
}

type Customer struct {
	CustomerId string
	Name       string
	Age        int
}

type Employee struct {
	EmployeeId string
	name       string `format:"normal"`
	age        int
}

// 一定要首字母大写的方法才能反射
func (e *Employee) UpdateAge(newVal int) {
	e.age = newVal
}

func TestInvokeByName(t *testing.T) {
	e := Employee{"ss", "mike", 40}
	t.Logf("value: %[1]v , other: %[1]T", reflect.ValueOf(e).FieldByName("name"))

	if namefield, ok := reflect.TypeOf(e).FieldByName("name"); !ok {
		t.Error("fail to get name")
	} else {
		t.Log(namefield.Tag.Get("format"))
	}

	reflect.ValueOf(&e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(20)})
	t.Log(e)
}

func fillBySettings(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer to the struct other")
	}
	// Elem方法签名: func (v Value) Elem() Value
	// Elem returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr. It returns the zero Value if v is nil.
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct other")
	}

	if settings == nil {
		return errors.New("settings is nil.")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}

		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "hello", "Age": 5}
	c := Customer{}

	if err := fillBySettings(&c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(c)
}

/* Implements、ConvertibleTo、AssignableTo使用*/
type Y int

func (Y) String() string {
	return ""
}
func TestAssitMethod(t *testing.T) {
	var a Y
	tt := reflect.TypeOf(a)
	// Implements不能直接使用类型作为参数
	st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	t.Log(tt.Implements(st))

	it := reflect.TypeOf(0)
	t.Log(tt.ConvertibleTo(it))

	t.Log(tt.AssignableTo(st), tt.AssignableTo(it))
}

type User struct {
	Name string
	code int
}

func TestRelectValue(t *testing.T) {
	a := 100
	va, vp := reflect.ValueOf(a), reflect.ValueOf(&a).Elem()
	t.Log(va.CanAddr(), va.CanSet()) // 接口变量会复制变量，且unaddressable
	t.Log(vp.CanAddr(), vp.CanSet())

	p := new(User)
	v := reflect.ValueOf(p).Elem()
	name := v.FieldByName("Name")
	code := v.FieldByName("code")
	t.Logf("name: can addr= %v, can set = %v", name.CanAddr(), name.CanSet())
	t.Logf("code: can addr= %v, can set = %v", code.CanAddr(), code.CanSet())
	if name.CanSet() {
		name.SetString("Tom")
	}
	// 非导出字段不能直接设置
	if code.CanAddr() {
		*(*int)(unsafe.Pointer(code.UnsafeAddr())) = 100
	}
	t.Logf("%+v", *p)

	// 通过Interface方法进行类型推断和转换，也可通过Value.Int等方法进行类型转换，但不支持ok-idom，失败会panic
	u := User{"hjl", 30}
	d := reflect.ValueOf(&u)
	if !d.CanInterface() {
		println("Can interface fall")
		return
	}
	p, ok := d.Interface().(*User)
	if !ok {
		println("Can interface fall")
		return
	}
	p.code++
	t.Logf("%+v", u)

	// channel对象设置示例
	c := make(chan int, 4)
	vc := reflect.ValueOf(c)
	if vc.TrySend(reflect.ValueOf(100)) {
		t.Log(vc.TryRecv())
	}

	// 接口有两种nil状态，可用IsNil判断
	var ia interface{} = nil
	var ib interface{} = (*int)(nil)
	t.Log(ia == nil, ib == nil)
	t.Log(reflect.ValueOf(ib).IsNil())
}

type X struct{}

func (X) Test(x, y int) (int, error) {
	return x + y, fmt.Errorf("err: %d", x+y)
}

func (X) Format(s string, a ...interface{}) string {
	return fmt.Sprintf(s, a...)
}

func TestReflectMethod(t *testing.T) {
	var x X
	v := reflect.ValueOf(&x)
	m := v.MethodByName("Test")
	in := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	}
	out := m.Call(in)
	for _, v := range out {
		t.Log(v)
	}

	// 对于变参，CallSlice更方便一些
	m2 := v.MethodByName("Format")
	out = m2.Call([]reflect.Value{
		reflect.ValueOf("%s=%d"),
		reflect.ValueOf("x"),
		reflect.ValueOf(101),
	})
	t.Log(out)
	out = m2.CallSlice([]reflect.Value{
		reflect.ValueOf("%s=%d"),
		reflect.ValueOf([]interface{}{"x", 102}),
	})
	t.Log(out)
}

func add(args []reflect.Value) (results []reflect.Value) {
	if len(args) == 0 {
		return nil
	}
	var ret reflect.Value
	switch args[0].Kind() {
	case reflect.Int:
		n := 0
		for _, a := range args {
			n += int(a.Int())
		}
		ret = reflect.ValueOf(n)
	case reflect.String:
		ss := make([]string, 0, len(args))
		for _, s := range args {
			ss = append(ss, s.String())
		}
		ret = reflect.ValueOf(strings.Join(ss, ""))
	}
	results = append(results, ret)
	return
}

func makeAdd(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), add) //实现通用模板，对应不同类型
	fn.Set(v)
}

func TestReflectMakeFunc(t *testing.T) {
	var intAdd func(x, y int) int
	var strAdd func(a, b string) string
	makeAdd(&intAdd)
	makeAdd(&strAdd)
	t.Log(intAdd(100, 200))
	t.Log(strAdd("hello ", "world"))
}
