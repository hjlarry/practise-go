package basic

import (
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
	"unsafe"
)

func TestString(t *testing.T) {
	// 字符串的零值为空
	var a string
	t.Log("s" + a + "b")

	var s = "嘻哈china"
	for i := 0; i < len(s); i++ {
		t.Log(s[i])
	}
	// 长度为底层字节数组的长度
	t.Log(len(s))
	// s[3] = '2' string是不可变的byte slice

	for codepoint, runeValue := range s {
		t.Log(codepoint, int32(runeValue))
	}

	var s1 = s[0:3]
	t.Log(s1)
	var s2 = s[6:9]
	t.Log(s2)

	var b = []byte(s) // 字符串转字节切片
	t.Log(b)
	b[6] = 100
	var s3 = string(b) // 字节切片转字符串
	t.Log(s3)

	var c = []rune(s)
	t.Log(c)
	t.Log("c[0] size:", unsafe.Sizeof(c[0]))
	t.Logf("%x", c[0])
	// 字符串可以是任意二进制数
	s = "\xE4\xBA\xB5\xFF"
	t.Log(s)
}

func TestStringToRune(t *testing.T) {
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
