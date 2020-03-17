package main

// 反汇编后该函数内直接调用了CALL main.factorial(SB)说明没有进行尾递归优化
func factorial(n, total int) int {
	if n == 1 {
		return total
	}
	return factorial(n-1, n*total)
}

func main() {
	println(factorial(5, 1))
}
