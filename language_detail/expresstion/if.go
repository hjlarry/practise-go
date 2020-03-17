package main

//go:noinline
func count() int{
	return 6
}

func main(){
	if x := count(); x > 5 {
		println(123)
	}else if x > 7 { //死代码，编译器自动剔除
		println(456)
	}
}