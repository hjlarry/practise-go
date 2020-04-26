package main

import "fmt"

func main(){
	a := 12345
	b := a
	fmt.Println(&a,&b)
}