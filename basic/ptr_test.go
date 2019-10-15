package basic

import "testing"

func TestPtr(t *testing.T) {
	x := 10
	var p *int = &x //获取地址，保存到指针变量
	*p += 20        // 用指针间接引用，并更新对象
	t.Log(p, *p, &x, x)

	m := map[string]int{"a": 1}
	//t.Log(&m["a"])  报错：can not take the address of m["a"]
	t.Log(&m)

	p2 := &x
	//p2++ 报错：p2++ (non-numeric type *int)
	//var p3 *int= p2+1 报错
	t.Log(p2 == p)
}
