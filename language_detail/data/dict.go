package main

// type User struct {
// 	name string
// 	age  int
// }

// func main() {
// 	u := map[int]User{
// 		1: {"mike", 20},
// 		2: {"john", 23},
// 	}

// 	// 这一步可以分解为拿到u[1].age的内存地址，然后++运算，然后写回
// 	// 但字典是可能正在rehash的，所以不允许这样操作，字典元素设计为不可寻址（not addressable）
// 	// u[1].age++ //cannot assign to struct field u[1].age in map

// 	x := u[1]
// 	x.age++
// 	// 这样可以，是因为使用字典本身提供的接口进行一次性赋值
// 	u[1] = x

// 	fmt.Printf("%v \n", u)
// }

func main() {
	var n map[int]int  // nil，只是空指针
	m := map[int]int{} // empty，空字典已分配内存

	println(n[100])
	// n[100] = 1  空指针字典可读不可写

	println(n == nil, m == nil)
	println(len(n), len(m))

}
