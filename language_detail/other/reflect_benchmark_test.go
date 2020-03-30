package main

import (
	"reflect"
	"testing"
)

/* 直接赋值和反射赋值的性能差异 */
type Data struct {
	X int
}

var d = new(Data)

func set(x int) {
	d.X = x
}

func rset(x int) {
	v := reflect.ValueOf(d).Elem()
	f := v.FieldByName("X")
	f.Set(reflect.ValueOf(x))
}

// BenchmarkSet-4   	2000000000	         0.44 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set(100)
	}
}

// BenchmarkRSet-4   	20000000	       122 ns/op	      16 B/op	       2 allocs/op
func BenchmarkRSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rset(100)
	}
}

// 缓存反射数据
var v = reflect.ValueOf(d).Elem()
var f = v.FieldByName("X")

func rset2(x int) {
	f.Set(reflect.ValueOf(x))
}

// BenchmarkRSet2-4   	50000000	        35.6 ns/op	       8 B/op	       1 allocs/op
func BenchmarkRSet2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rset2(100)
	}
}

/* 方法直接调用和反射调用的性能差异 */
func (x *Data) Inc() {
	x.X++
}

var dd = new(Data)
var vv = reflect.ValueOf(dd)
var mm = vv.MethodByName("Inc")

func call() {
	dd.Inc()
}
func rcall() {
	mm.Call(nil)
}

// BenchmarkCall-4   	2000000000	         1.55 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call()
	}
}

// BenchmarkRcall-4   	10000000	       166 ns/op	       8 B/op	       1 allocs/op
func BenchmarkRcall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rcall()
	}
}
