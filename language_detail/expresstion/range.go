package main

import "fmt"

// func main() {
// 	data := [4]int{0x11, 0x22, 0x33, 0x44}
// 	// 反汇编证明range中的data是隐式复制了main中的data
//  // 编译优化打开，则这个示例中range中的data并不复制main中的，因为没有发生数据修改行为编译器认为没必要复制
// 	for i, d := range data {
// 		_, _ = i, d
// 	}
// }

// func main() {
// 	data := [4]int{0x11, 0x22, 0x33, 0x44}
// 	fmt.Printf("%p\n", &data)
// 	for i, d := range data {
// 		_, _ = i, d
// 		// 输出得到的地址相同，说明这里拿到的data就是main中的data
// 		fmt.Printf("%p\n", &data)
// 	}
// }

func main() {
	data := [4]int{100, 200, 300, 400}
	// data := []int{100, 200, 300, 400}  如果是切片，则range中复制的也是切片引用

	for i, d := range data {
		// 修改main中的data，但打印d的值，说明d是从range中这个隐式复制的data中取值的
		if i == 0 {
			data[1] += 1
			data[2] += 1
			data[3] += 1
		}
		fmt.Println(i, d, data[i])
	}
	fmt.Println(data)
}
