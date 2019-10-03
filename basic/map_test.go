package basic

import "testing"

func TestMap(t *testing.T) {
	map1 := make(map[int]int)
	map1[3] = 19
	t.Log(map1)

	map2 := map[int]int{
		1: 1,
		2: 4,
		3: 9,
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

	// 遍历数组元素
	for k, v := range map1 {
		t.Log(k, v)
	}
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