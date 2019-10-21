package basic

import (
	"errors"
	"reflect"
	"testing"
	"unsafe"
)

func TestSlice(t *testing.T) {
	var a []int               //仅定义，并未初始化
	b := []int{}              //初始化表达式完成了全部创建过程
	t.Log(a == nil, b == nil) //切片只能和nil去比较
	t.Logf("a: %#v", (*reflect.SliceHeader)(unsafe.Pointer(&a)))
	t.Logf("b: %#v", (*reflect.SliceHeader)(unsafe.Pointer(&b)))
	t.Logf("a size: %d", unsafe.Sizeof(a)) //虽未初始化，但分配了内存
	t.Logf("b size: %d", unsafe.Sizeof(b))

	var ss1 = make([]int, 5)
	var ss2 = make([]int, 0, 5)
	t.Log(len(ss1), cap(ss1))
	t.Log(len(ss2), cap(ss2))
	s3 := append(ss1, 99)
	s4 := append(ss2, 88)
	t.Log(ss1, ss2)
	t.Log(s3, s4)

	s := []int{0, 1, 2, 3, 4}
	p := &s
	p0 := &s[0] //可获取元素地址，但不能像数组那样直接用指针访问元素内容
	p1 := &s[1]
	t.Log(p, p0, p1)

	(*p)[0] += 100
	*p1 += 100
	t.Log(s)

	// 任一切片中的值变化会导致底层数组和其他引用该数组的值的变化
	var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ss3 := arr[2:5]
	ss4 := arr[2:]
	ss3[2] = 0
	t.Log(ss3)
	t.Log(ss4)
	t.Log(arr)

}

func array() [1024]int {
	var x [1024]int
	for i := 0; i < 1024; i++ {
		x[i] = i
	}
	return x
}

func slice() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

// 并非所有时候都适合用切片代替数组，因为切片底层数组可能会在堆上分配内存，且小数组在栈上拷贝的消耗也未必比make大
// BenchmarkArray-4   	 2000000	       961 ns/op	       0 B/op	       0 allocs/op
func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

// BenchmarkSlice-4   	 1000000	      1791 ns/op	    8192 B/op	       1 allocs/op
func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}

func TestStack(t *testing.T) {
	stack := make([]int, 0, 5)

	push := func(x int) error {
		n := len(stack)
		if n == cap(stack) {
			return errors.New("stack is full")
		}
		stack = stack[:n+1]
		stack[n] = x
		return nil
	}

	pop := func() (int, error) {
		n := len(stack)
		if n == 0 {
			return 0, errors.New("stack is empty")
		}
		x := stack[n-1]
		stack = stack[:n-1]
		return x, nil
	}

	for i := 0; i < 7; i++ {
		t.Logf("push %d: %v, %v", i, push(i), stack)
	}

	for i := 0; i < 7; i++ {
		x, err := pop()
		t.Logf("pop %d, %v, %v", x, err, stack)
	}
}

func TestAppend(t *testing.T) {
	s := make([]int, 0, 100)
	s1 := s[:2:4]
	// 当超出s1的cap时就会新分配数组，而非底层数组的长度限制
	s2 := append(s1, 1, 2, 3, 4, 5, 6)
	t.Logf("s1: %p: %v", &s1[0], s1)
	t.Logf("s2: %p: %v", &s2[0], s2)
	t.Logf("s data: %v", s[:10])
	t.Logf("s1 cap: %d, s2 cap: %d", cap(s1), cap(s2))
}

func TestCopy(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[5:8]
	// 在同一底层数组的不同区间复制
	n := copy(s[4:], s1)
	t.Log(n, s)
	s2 := make([]int, 6)
	// 在不同数组间复制
	n = copy(s2, s)
	t.Log(n, s2)
}
