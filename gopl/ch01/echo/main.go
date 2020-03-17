package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println(i, arg)
	}
	fmt.Println(s)
	fmt.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
