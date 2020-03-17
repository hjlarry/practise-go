package main

// func test() func() {
// 	x := 100
// 	return func() {
// 		println(x)
// 	}
// }

// func main() {
// 	closure := test()
// 	closure()
// }

/*
11行中的test函数调用结束，那么test中的局部变量应该是被消除的
test返回的是一个函数，未执行的，按常理应该没有办法打印出x，这就涉及到编译器对闭包的处理

本质上test返回的是一个结构体
closure struct{
	func : 指针指向匿名函数的地址，
	x: 环境变量（已逃逸到堆上）
}
本质上调用的时候closure.func(closure.x)

go build -gcflags "-N -l -S" 2>a.txt closure.go
核心编译指令：
"".test STEXT size=126 args=0x8 locals=0x28
	0x0026 00038 (/root/.mac/gocode/closure.go:4)	MOVQ	$100, "".x+16(SP)
	0x002f 00047 (/root/.mac/gocode/closure.go:5)	LEAQ	type.noalg.struct { F uintptr; "".x int }(SB), AX
	0x003a 00058 (/root/.mac/gocode/closure.go:5)	CALL	runtime.newobject(SB)  //堆上创建函数对象
	0x0049 00073 (/root/.mac/gocode/closure.go:5)	LEAQ	"".test.func1(SB), CX  //函数地址写到CX寄存器，再到AX+0的位置
	0x0050 00080 (/root/.mac/gocode/closure.go:5)	MOVQ	CX, (AX)
	0x005a 00090 (/root/.mac/gocode/closure.go:5)	MOVQ	"".x+16(SP), CX
	0x005f 00095 (/root/.mac/gocode/closure.go:5)	MOVQ	CX, 8(AX)			   //100最终到了AX+8的位置
	0x0068 00104 (/root/.mac/gocode/closure.go:5)	MOVQ	AX, "".~r0+48(SP)	   //写到AX中返回


"".main STEXT size=65 args=0x0 locals=0x18
	0x0022 00034 (/root/.mac/gocode/closure.go:11)	MOVQ	(SP), DX               //(SP)就是test函数的返回
	0x0026 00038 (/root/.mac/gocode/closure.go:11)	MOVQ	DX, "".closure+8(SP)
	0x002b 00043 (/root/.mac/gocode/closure.go:12)	MOVQ	(DX), AX
	0x002e 00046 (/root/.mac/gocode/closure.go:12)	CALL	AX                     //调用匿名函数

"".test.func1 STEXT size=84 args=0x0 locals=0x18
	0x001d 00029 (/root/.mac/gocode/closure.go:5)	MOVQ	8(DX), AX
	0x0021 00033 (/root/.mac/gocode/closure.go:5)	MOVQ	AX, "".x+8(SP)
	0x0026 00038 (/root/.mac/gocode/closure.go:6)	CALL	runtime.printlock(SB)
	0x002b 00043 (/root/.mac/gocode/closure.go:6)	MOVQ	"".x+8(SP), AX
	0x0030 00048 (/root/.mac/gocode/closure.go:6)	MOVQ	AX, (SP)
	0x0034 00052 (/root/.mac/gocode/closure.go:6)	CALL	runtime.printint(SB)

*/

// 闭包的使用场景一
// 计数器，类似class的效果，给函数绑定了属性
// func getCounter() func() int {
// 	x := 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }

// func main() {
// 	c := getCounter()
// 	println(c())
// 	println(c())
// 	println(c())
// }

// 闭包的使用场景二
// 构成相对封闭的区间，这个里面的三个函数共享同一个变量
// func test() (func(), func(), func()) {
// 	x := 100
// 	fmt.Printf("%v\n", &x)
// 	return func() {
// 			fmt.Printf("f1.x:%v\n", &x)
// 		}, func() {
// 			fmt.Printf("f2.x:%v\n", &x)
// 		}, func() {
// 			fmt.Printf("f3.x:%v\n", &x)
// 		}
// }

// func main() {
// 	f1, f2, f3 := test()
// 	f1()
// 	f2()
// 	f3()
// }

// 闭包的使用场景三
// 把数据和逻辑绑定在一起传输
// func main() {
// 	db := NewDatabase("localhost:5432")
// 	http.HandleFunc("/hello", hello(db))
// 	http.ListenAndServe(":3000", nil)
// }
// func hello(db Database) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, db.Url)
// 	}
// }

// 闭包的使用场景四
// 改变函数的签名，类似partial或proxy签名

func test(x int) {
	println(x)
}

func partial(f func(int), x int) func() {
	return func() {
		f(x)
	}
}

func main() {
	f := partial(test, 100)
	f()
}
