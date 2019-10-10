package professional

import (
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f uint64 = 99
	t.Log(reflect.ValueOf(f))
	t.Log(reflect.TypeOf(f))
	t.Log(reflect.ValueOf(f).Type())
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
	name       string
	age        int
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
