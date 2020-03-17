package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 一、字符串基础
// 字符串的本质也是struct{ptr, len}
// 但它和slice不同的是，会把字符串的结构体和它的底层字节序列看做一体的，底层序列单独拿出是没有意义的
//func main() {
//	//由于字符串是只读的，所以不允许去拿字节序列元素的指针
//	//子串和原串可以共享底层字节数组
//	ss := "hello, world"
//	sub := ss[:6]
//	fmt.Printf("%#v \n", *(*reflect.StringHeader)(unsafe.Pointer(&ss)))
//	fmt.Printf("%#v \n", *(*reflect.StringHeader)(unsafe.Pointer(&sub)))
//
//	// 字符串的零值为空
//	var a string
//	fmt.Println("s" + a + "b")
//
//	// 字符串字面量允许十六进制、八进制和UTF编码格式
//	var s = "嘻哈ch\x61\142\u0041"
//	fmt.Println(s)
//	// len返回底层字节数组的长度,RuneCountInString返回准确的unicode字符数量
//	fmt.Println(len(s), utf8.RuneCountInString(s))
//	// 截取并拼接一个不合法的字符串
//	s = string(s[0:1] + s[3:4])
//	fmt.Println(s, utf8.ValidString(s))
//
//	// 可使用`不做转义处理
//	s = `line\r\n
//      line2`
//	fmt.Println(s)
//	// 允许索引号访问字节数组（非字符串），但不能使用&s[3]访问其地址，也不能s[3]='2'进行赋值
//	fmt.Println(s[3])
//
//	// 切片语法返回的子串，其内部依旧指向原字节数组
//	s1 := s[:3]
//	s2 := s[1:4]
//	fmt.Println(s1, s2)
//	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
//	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
//
//	// 两种遍历方式: byte和rune
//	s = "哈a"
//	for i := 0; i < len(s); i++ {
//		fmt.Println(s[i])
//	}
//	for codepoint, runeValue := range s {
//		fmt.Println(codepoint, int32(runeValue))
//	}
//
//	//使用单引号的字面量，默认为rune
//	r := '我'
//	fmt.Printf("%T", r)
//
//	sr := "中华人民共和国"
//	for _, c := range sr {
//		fmt.Printf("%[1]c %[1]d", c)
//	}
//
//	for i, w := 0, 0; i < len(sr); i += w {
//		runeValue, width := utf8.DecodeRuneInString(sr[i:])
//		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
//		w = width
//	}
//}

/* ------------------------------------------------------------------------------------- */
// 二、传递字符串
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
// 三、 字符串转换
// 1. 转换行为发生复制
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
//func main() {
//	s := "hello, world!"
//	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
//	bh := reflect.SliceHeader{
//		Data: sh.Data,
//		Len:  sh.Len,
//		Cap:  sh.Len,
//	}
//	bs := *(*[]byte)(unsafe.Pointer(&bh))
//	fmt.Printf("%#v \n", sh)
//	fmt.Printf("%#v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&bs)))
//}

// 2. 修改字符串
//func pp(format string, ptr interface{}) {
//	p := reflect.ValueOf(ptr).Pointer()
//	h := (*uintptr)(unsafe.Pointer(p))
//	fmt.Printf(format, *h)
//}
//
//// 要修改字符串，可以将其转换为[]byte或[]rune，修改完后转回，但都会重新分配内存、复制数据
//// 使用"非安全"方法可以改善
//func main() {
//	s := "hello, world!"
//	pp("s: %x\n", &s)
//
//	bs := []byte(s)
//	s2 := string(bs)
//	pp("string to []byte,  bs: %x\n", &bs)
//	pp("[]byte to string,  s2: %x\n", &s2)
//
//	rs := []rune(s)
//	s3 := string(bs)
//	pp("string to []rune,  bs: %x\n", &rs)
//	pp("[]rune to string,  s2: %x\n", &s3)
//
//	bb := []byte("hello world!")
//	ss := *(*string)(unsafe.Pointer(&bb))
//	pp("bb: %x\n", &bb)
//	pp("ss: %x\n", &ss)
//
//}

// 3. 使用strings和strconv包
func main() {
	s := "a,b,c,d"
	list := strings.Split(s, ",")
	fmt.Println(list)
	s = strings.Join(list, "-")
	fmt.Println(s)

	i := 1
	ss := strconv.Itoa(i)
	fmt.Printf("%T", ss)
	// 字符串转数字时会返回值+错误
	b, _ := strconv.Atoi(ss)
	fmt.Printf("%T", b)
	c, err := strconv.Atoi("a")
	fmt.Println(c)
	fmt.Println(err)
}

