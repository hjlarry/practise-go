package professional

import (
	"fmt"
	"reflect"
	"testing"

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
		println("unknown type")
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
	t.Logf("value: %[1]v , type: %[1]T", reflect.ValueOf(e).FieldByName("name"))

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
		return errors.New("the first param should be a pointer to the struct type")
	}
	// Elem方法签名: func (v Value) Elem() Value
	// Elem returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr. It returns the zero Value if v is nil.
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct type")
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
