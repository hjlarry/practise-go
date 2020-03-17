package main

// 编译之后，和if语句编译后指令相同
// func main() {
// 	x := 0x11
// 	switch {
// 	case x > 0x2:
// 		println(123)
// 	case x < 0x2:
// 		println(456)
// 	default:
// 		println(789)
// 	}
// }

// func main() {
// 	switch x := 5; {
// 	case x > 5:
// 		println(123)
// 	case x > 0 && x < 5: // 不能写成case x>0,x<5  逗号表示||
// 		println(456)
// 	default:
// 		println(789)
// 	}
// }

func main() {
	switch x := 5; x {
	case 5:
		// x++
		x += 10
		if x == 15 {
			break //中断整个switch
		}
		println(x)
		fallthrough // 只能在case块的最后一句
	case 6:
		println(x, 22)
	}
}
