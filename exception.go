package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

// TestException ...
func TestException() {
	var f, err = os.Open("main.go")
	if err != nil {
		// 文件不存在、权限等原因
		fmt.Println("open file failed reason:" + err.Error())
	}
	// 推迟到函数尾部调用，确保文件会关闭
	defer f.Close()
	// 存储文件内容
	var content = []byte{}
	// 临时的缓冲，按块读取，一次最多读取 100 字节
	var buf = make([]byte, 100)
	for {
		// 读文件，将读到的内容填充到缓冲
		n, err := f.Read(buf)
		if n > 0 {
			//  … 操作符的作用是将切片参数的所有元素展开后传递给 append 函数
			content = append(content, buf[:n]...)
		}
		if err != nil {
			// 遇到流结束或者其它错误
			fmt.Println(err.Error())
			break
		}
	}
	fmt.Println(string(content))
}

// TestException1 ... 一个redis错误处理的例子，不轻易使用异常捕获
func TestException1() {
	var client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	var val1, val2, val3 int
	valstr1, err := client.Get("value1").Result()
	if err == nil {
		val1, err = strconv.Atoi(valstr1)
		if err != nil {
			fmt.Println("value1 not a valid integer")
			return
		}
		// redis.Nil 就是客户端专门为 key 不存在这种情况而定义的错误对象
	} else if err == redis.Nil {
		fmt.Println("redis access error:" + err.Error())
		// return
	}

	valstr2, err := client.Get("value2").Result()
	if err == nil {
		val2, err = strconv.Atoi(valstr2)
		if err != nil {
			fmt.Println("value1 not a valid integer")
			return
		}
	} else if err != redis.Nil {
		fmt.Println("redis access error:" + err.Error())
		return
	}

	val3 = val1 * val2
	ok, err := client.Set("value2", val3, 0).Result()
	if err != nil {
		fmt.Println("set value error:" + err.Error())
		return
	}
	fmt.Println(ok)
}

var errNeg = fmt.Errorf("not a positive number")

func fact(a int) int {
	if a <= 0 {
		panic(errNeg)
	}
	return a
}

// TestException2 ... 异常捕获，但panic可以抛出错误以外的其他对象
func TestException2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error cached:", err)
		}
	}()
	fmt.Println(fact(10))
	fmt.Println(fact(5))
	fmt.Println(fact(-5))
	fmt.Println(fact(15))
}
