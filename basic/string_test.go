package basic

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
	"unsafe"
)

/* 一、字符串基础 */
func TestString(t *testing.T) {
	// 字符串的零值为空
	var a string
	t.Log("s" + a + "b")

	// 字符串字面量允许十六进制、八进制和UTF编码格式
	var s = "嘻哈ch\x61\142\u0041"
	t.Log(s)
	// len返回底层字节数组的长度,RuneCountInString返回准确的unicode字符数量
	t.Log(len(s), utf8.RuneCountInString(s))
	// 截取并拼接一个不合法的字符串
	s = string(s[0:1] + s[3:4])
	t.Log(s, utf8.ValidString(s))

	// 可使用`不做转义处理
	s = `line\r\n
      line2`
	t.Log(s)
	// 允许索引号访问字节数组（非字符串），但不能使用&s[3]访问其地址，也不能s[3]='2'进行赋值
	t.Log(s[3])

	// 切片语法返回的子串，其内部依旧指向原字节数组
	s1 := s[:3]
	s2 := s[1:4]
	t.Log(s1, s2)
	t.Logf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	t.Logf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))

	// 两种遍历方式: byte和rune
	s = "哈a"
	for i := 0; i < len(s); i++ {
		t.Log(s[i])
	}
	for codepoint, runeValue := range s {
		t.Log(codepoint, int32(runeValue))
	}
}

func TestRune(t *testing.T) {
	//使用单引号的字面量，默认为rune
	r := '我'
	t.Logf("%T", r)

	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		t.Logf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

/* 二、字符串转换 */
func pp(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)
}

// 要修改字符串，可以将其转换为[]byte或[]rune，修改完后转回，但都会重新分配内存、复制数据
// 使用"非安全"方法可以改善
func TestStringConvert(t *testing.T) {
	s := "hello, world!"
	pp("s: %x\n", &s)

	bs := []byte(s)
	s2 := string(bs)
	pp("string to []byte,  bs: %x\n", &bs)
	pp("[]byte to string,  s2: %x\n", &s2)

	rs := []rune(s)
	s3 := string(bs)
	pp("string to []rune,  bs: %x\n", &rs)
	pp("[]rune to string,  s2: %x\n", &s3)

	bb := []byte("hello world!")
	ss := *(*string)(unsafe.Pointer(&bb))
	pp("bb: %x\n", &bb)
	pp("ss: %x\n", &ss)

}

/* 三、strings和strconv包 */
func TestStrings(t *testing.T) {
	s := "a,b,c,d"
	list := strings.Split(s, ",")
	t.Log(list)
	s = strings.Join(list, "-")
	t.Log(s)
}

func TestStrconv(t *testing.T) {
	i := 1
	s := strconv.Itoa(i)
	t.Logf("%T", s)
	// 字符串转数字时会返回值+错误
	b, _ := strconv.Atoi(s)
	t.Logf("%T", b)
	c, err := strconv.Atoi("a")
	t.Log(c)
	t.Log(err)
}

/* 四、直接拼接、使用join、使用bytes.buffer的性能比较 */
func plus() string {
	var s string
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	return s
}

//BenchmarkPlus-4   	   10000	    106563 ns/op
func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plus()
	}
}

func join(a []string, sep string) string {
	n := len(sep) * (len(a) - 1)  // 统计分隔符长度
	for i := 0; i < len(a); i++ { // 统计所有待拼接字符长度
		n += len(a[i])
	}
	b := make([]byte, n) //一次分配所需长度的数组空间
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

func testJoin() string {
	s := make([]string, 1000) //准备足够内存，避免中途扩张
	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}
	return join(s, "")
}

//BenchmarkJoin-4   	  200000	      9145 ns/op
func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testJoin()
	}
}

func buffer() string {
	var b bytes.Buffer
	b.Grow(1000) //准备足够内存，避免中途扩张
	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}
	return b.String()
}

//BenchmarkBuffer-4   	  200000	      6213 ns/op
func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buffer()
	}
}
