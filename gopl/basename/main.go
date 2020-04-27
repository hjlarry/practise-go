package main

import (
	"fmt"
	"strings"
)

// basename(s) 将看起来像是系统路径的前缀删除，同时将看似文件类型的后缀名部分删除
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
// 方案一、手工硬编码
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// 方案二、使用strings.LastIndex库函数
func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	fmt.Println(basename("a/b/c.go"))  // "c"
	fmt.Println(basename("c.d.go"))    // "c.d"
	fmt.Println(basename("abc"))       // "abc"
	fmt.Println(basename2("a/b/c.go")) // "c"
	fmt.Println(basename2("c.d.go"))   // "c.d"
	fmt.Println(basename2("abc"))      // "abc"
}
