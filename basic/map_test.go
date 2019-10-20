package basic

import "testing"

func TestMap(t *testing.T) {
	map1 := make(map[int]int)
	map1[3] = 19
	t.Log(map1)

	map2 := map[int]struct {
		x int
	}{
		1: {x: 100},
		2: {x: 200},
	}
	t.Log(map2)

	// map在make创建时，第二个参数实际是cap，却无法用cap(map3)来求容量
	// 因为容量可变，只是初始化时给一个容量使内存分配好，避免随元素增多map的不断扩容迁移，提升了性能
	map3 := make(map[int]int, 3)
	map3[1] = 1
	map3[2] = 1
	map3[3] = 1
	map3[4] = 1
	t.Log(map3)

	var m1 map[string]int //nil字典，未初始化，不能写入
	m2 := map[string]int{}
	t.Log(m1 == nil, m2 == nil)
}

func TestMapElement(t *testing.T) {
	map1 := map[int]int{
		1: 0,
		2: 3,
	}
	// 无论是某个key本身对应的是零值，还是这个key不存在，都会返回零值
	t.Log(map1[1])
	t.Log(map1[3])

	// 判断元素是否存在的方案
	v, exist := map1[1]
	t.Log(v, exist)
	v, exist = map1[3]
	t.Log(v, exist)

	// 删除元素，不存在不会报错
	delete(map1, 4)

	// 遍历数组元素，每次遍历顺序可能不同
	for k, v := range map1 {
		t.Log(k, v)
	}
}

func TestModify(t *testing.T) {
	type user struct {
		name string
		age  int
	}
	m := map[int]user{1: {"Tome", 19}}
	// 字典是not addressable，不能直接修改value成员
	// m[1].age += 1  错误：can not assign to struct field m[1].age in map
	// 而应该先返回整个value，再更新map
	u := m[1]
	u.age += 1
	m[1] = u
	// 或使用指针类型的value
	m2 := map[int]*user{1: &user{"Jack", 20}}
	m2[1].age += 1
	t.Log(m, m2, m2[1].age)
}

func TestMapForFunc(t *testing.T) {
	// key是整数，value是func，使用map创建了工厂方法
	map1 := map[int]func(op int) int{}
	map1[1] = func(op int) int { return op }
	map1[2] = func(op int) int { return op * op }
	map1[3] = func(op int) int { return op * op * op }
	t.Log(map1[1](3))
	t.Log(map1[2](3))
	t.Log(map1[3](3))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	// 添加
	mySet[1] = true
	mySet[2] = true
	mySet[1] = true
	// 判断是否存在
	if mySet[2] {
		t.Log("exist")
	}
	// set元素数量
	t.Log(len(mySet))
	// 删除
	delete(mySet, 2)
	if !mySet[2] {
		t.Log("not exist")
	}
}

// 在迭代期间删除或新增键值对是安全的
func TestIterMap(t *testing.T) {
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		m[i] = i + 10
	}
	for k := range m {
		if k == 5 {
			m[100] = 1000
		}
		delete(m, k)
		t.Log(k, m)
	}
}
