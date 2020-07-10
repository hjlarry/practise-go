package main

import (
	"context"
	"fmt"
	"time"
)

// 一、 通过上下文传值
// 查找值时实际上是递归查找，下例中b方法包装了新的值，最后就是新值
func a(ctx context.Context) {
	ctx = context.WithValue(ctx, "aa", "hello world!")
	b(ctx)
}

func b(ctx context.Context) {
	ctx = context.WithValue(ctx, "aa", 1234)
	c(ctx)
}

func c(ctx context.Context) {
	time.Sleep(time.Second * 2)
	fmt.Println("c exec.", ctx.Value("aa"))
}

func main() {
	a(context.Background())
}
