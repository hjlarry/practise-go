package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("1234567892"))
	fmt.Println(comma2("1234567892"))
	fmt.Println(comma3("-1234567892.9678"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// 练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。
func comma2(s string) string {
	var buf bytes.Buffer
	pre := len(s) % 3
	if pre == 0 {
		pre = 3
	}
	buf.WriteString(s[:pre])
	for i := pre; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

// 练习 3.11： 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。
func comma3(s string) string {
	var buf bytes.Buffer
	mantissaStart := 0
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		mantissaStart = 1
	}
	mantissaEnd := strings.Index(s, ".")
	if mantissaEnd == -1 {
		mantissaEnd = len(s)
	}
	mantissa := s[mantissaStart:mantissaEnd]
	pre := len(mantissa) % 3
	if pre > 0 {
		buf.WriteString(mantissa[:pre])
		if len(mantissa) > pre {
			buf.WriteByte(',')
		}
	}

	for i, c := range mantissa[pre:] {
		if i%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(c)
	}
	buf.WriteString(s[mantissaEnd:])
	return buf.String()
}
