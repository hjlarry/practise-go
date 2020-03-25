package main

import (
	"fmt"
	"unsafe"
)

// 一、创建map
//func main() {
//	map1 := make(map[int]int)
//	map1[3] = 19
//	fmt.Println(map1)
//
//	map2 := map[int]struct {
//		x int
//	}{
//		1: {x: 100},
//		2: {x: 200},
//	}
//	fmt.Println(map2)
//
//	// map在make创建时，第二个参数实际是cap，却无法用cap(map3)来求容量
//	// 因为容量可变，只是初始化时给一个容量使内存分配好，避免随元素增多map的不断扩容迁移，提升了性能
//	map3 := make(map[int]int, 3)
//	map3[1] = 1
//	map3[2] = 1
//	map3[3] = 1
//	map3[4] = 1
//	fmt.Println(map3)
//
//	var m1 map[string]int //nil字典，未初始化，不能写入
//	m2 := map[string]int{}
//	fmt.Println(m1 == nil, m2 == nil)
//}

// 二、操作map元素
//func main() {
//	map1 := map[int]int{
//		1: 0,
//		2: 3,
//	}
//	// 无论是某个key本身对应的是零值，还是这个key不存在，都会返回零值
//	fmt.Println(map1[1])
//	fmt.Println(map1[3])
//
//	// 判断元素是否存在的方案
//	v, exist := map1[1]
//	fmt.Println(v, exist)
//	v, exist = map1[3]
//	fmt.Println(v, exist)
//
//	// 删除元素，不存在不会报错
//	delete(map1, 4)
//
//	// 遍历数组元素，每次遍历顺序可能不同
//	for k, v := range map1 {
//		fmt.Println(k, v)
//	}
//}


// 三、字典是不可寻址的
//func main() {
//	type user struct {
//		name string
//		age  int
//	}
//	m := map[int]user{1: {"Tome", 19}}
//
//	// m[1].age += 1  错误：can not assign to struct field m[1].age in map
//	// 这一步可以分解为拿到u[1].age的内存地址，然后++运算，然后写回
//	// 但字典是可能正在rehash的，所以不允许这样操作，字典元素设计为不可寻址（not addressable）
//
//	u := m[1]
//	u.age += 1
//	// 这样可以，是因为使用字典本身提供的接口进行一次性赋值
//	m[1] = u
//
//	// 或使用指针类型的value
//	m2 := map[int]*user{1: &user{"Jack", 20}}
//	m2[1].age += 1
//	fmt.Println(m, m2, m2[1].age)
//}


// 四、 空指针字典
//func main() {
//	var n map[int]int  // nil，只是空指针
//	m := map[int]int{} // empty，空字典已分配内存
//
//	println(n[100])
//	// n[100] = 1  空指针字典可读不可写
//
//	println(n == nil, m == nil)
//	println(len(n), len(m))
//
//}


// 五、map的特殊使用场景
// 1. 创建工厂方法
//func main() {
//	map1 := map[int]func(op int) int{}
//	map1[1] = func(op int) int { return op }
//	map1[2] = func(op int) int { return op * op }
//	map1[3] = func(op int) int { return op * op * op }
//	fmt.Println(map1[1](3))
//	fmt.Println(map1[2](3))
//	fmt.Println(map1[3](3))
//}

// 2. 当做set
//func main() {
//	mySet := map[int]bool{}
//	// 添加
//	mySet[1] = true
//	mySet[2] = true
//	mySet[1] = true
//	// 判断是否存在
//	if mySet[2] {
//		fmt.Println("exist")
//	}
//	// set元素数量
//	fmt.Println(len(mySet))
//	// 删除
//	delete(mySet, 2)
//	if !mySet[2] {
//		//fmt.Println("not exist")
//	}
//}

// 六、迭代时删除或新增键值对是安全的
//func main() {
//	m := make(map[int]int)
//	for i := 0; i < 10; i++ {
//		m[i] = i + 10
//	}
//	for k := range m {
//		if k == 5 {
//			m[100] = 1000
//		}
//		delete(m, k)
//		fmt.Println(k, m)
//	}
//}

// 七、字典对象本身就是指针的包装，传参时无须再次取地址
func testmm(x map[string]int) {
	fmt.Printf("x: %p \n", x)
}
func main() {
	m := make(map[string]int)
	testmm(m)
	fmt.Printf("m: %p, %d \n", m, unsafe.Sizeof(m))

	m2 := map[string]int{}
	testmm(m2)
	fmt.Printf("m2: %p, %d \n", m2, unsafe.Sizeof(m2))
}
