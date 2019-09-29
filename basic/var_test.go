package basic

import (
	"errors"
	"testing"
)

func TestString(t *testing.T) {
	var s = "嘻哈china"
	for i := 0; i < len(s); i++ {
		t.Log(s[i])
	}

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
}

func TestVar(t *testing.T) {
	// 有符号整数，可以表示正负
	var a int8 = 1  // 1 字节
	var b int16 = 2 // 2 字节
	var c int32 = 3 // 4 字节
	var d int64 = 4 // 8 字节
	t.Log(a, b, c, d)

	// 无符号整数，只能表示非负数
	var ua uint8 = 1
	var ub uint16 = 2
	var uc uint32 = 3
	var ud uint64 = 4
	t.Log(ua, ub, uc, ud)

	// int 类型，在32位机器上占4个字节，在64位机器上占8个字节
	var e int = 5
	var ue uint = 5
	t.Log(e, ue)

	// bool 类型
	var f bool = true
	t.Log(f)

	// 字节类型
	var j byte = 'a'
	t.Log(j)

	// 字符串类型
	var g string = "abcdefg"
	t.Log(g)

	// 浮点数
	var h float32 = 3.14
	var i float64 = 3.141592653
	t.Log(h, i)

	var aArray [9]int
	t.Log(aArray)

	var aa = [10]int{1, 3, 4}
	var bb [10]int = [10]int{1, 3, 4}
	cc := [10]int{1, 3, 4}
	t.Log(aa)
	t.Log(bb)
	t.Log(cc)

	var seq [9]int
	for i := 0; i < len(seq); i++ {
		seq[i] = i * i
	}
	t.Log(seq)

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		t.Log(err)
	}
}
