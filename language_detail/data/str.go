package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 字符串的本质也是struct{ptr, len}
// 但它和slice不同的是，会把字符串的结构体和它的底层字节序列看做一体的，底层序列单独拿出是没有意义的

// 由于字符串是只读的，所以不允许去拿字节序列元素的指针
// 子串和原串可以共享底层字节数组
// func main() {
// 	s := "hello, world"
// 	sub := s[:6]
// 	fmt.Printf("%#v \n", *(*reflect.StringHeader)(unsafe.Pointer(&s)))
// 	fmt.Printf("%#v \n", *(*reflect.StringHeader)(unsafe.Pointer(&sub)))
// }

/* ------------------------------------------------------------------------------------- */

// go build -gcflags "-S" 2>a.txt str.go
// 字符串传递时也只传其字节数组地址和长度
// 0x001d 00029 (/root/.mac/gocode/str.go:22)	LEAQ	go.string."hello,world!"(SB), AX
// 0x0024 00036 (/root/.mac/gocode/str.go:22)	PCDATA	$0, $0
// 0x0024 00036 (/root/.mac/gocode/str.go:22)	MOVQ	AX, (SP)
// 0x0028 00040 (/root/.mac/gocode/str.go:22)	MOVQ	$12, 8(SP)
// 0x0031 00049 (/root/.mac/gocode/str.go:22)	CALL	"".test(SB)

//go:noinline
// func test(s string) {
// 	println(s)
// }

// func main() {
// 	a := "hello,world!"
// 	test(a)
// }

/* ------------------------------------------------------------------------------------- */
// func main() {
// 	s := "hello, world"

// 	bs := []byte(s)
// 	s2 := string(bs)
// 	rs := []rune(s)
// 	s3 := string(rs)

// 	// 每一次转换前后指针地址都不同，说明会发生复制。因为转换前是只读的、转换后是可写的，理应发生复制
// 	fmt.Printf("%#v \n", *(*reflect.StringHeader)(unsafe.Pointer(&s)))
// 	fmt.Printf("%#v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&bs)))
// 	fmt.Printf("%#v \n", *(*reflect.StringHeader)(unsafe.Pointer(&s2)))
// 	fmt.Printf("%#v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&rs)))
// 	fmt.Printf("%#v \n", *(*reflect.StringHeader)(unsafe.Pointer(&s3)))
// }
/* ------------------------------------------------------------------------------------- */
// func main() {
// 	// bs是struct{ptr, len, cap}，转为string忽略掉cap没有问题
// 	bs := []byte("hello, world!")
// 	fmt.Printf("%#v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&bs)))
// 	s := *(*string)(unsafe.Pointer(&bs))
// 	fmt.Printf("%#v \n", *(*reflect.StringHeader)(unsafe.Pointer(&s)))
// }

// 但是反过来转为slice时，少一个cap，必须构建一个头部
func main() {
	s := "hello, world!"
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	bs := *(*[]byte)(unsafe.Pointer(&bh))
	fmt.Printf("%#v \n", sh)
	fmt.Printf("%#v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&bs)))
}
