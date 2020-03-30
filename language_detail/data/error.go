package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

// 一、 一般的错误处理方式
//var errLess = errors.New("the number should not less than 2")
//var errBig = errors.New("the number should not big than 100")
//func GetFib(n int) ([]int, error) {
//	if n < 2 {
//		return nil, errLess
//	}
//	if n > 100 {
//		return nil, errBig
//	}
//	fiblist := []int{1, 1}
//	for i := 2; i < n; i++ {
//		fiblist = append(fiblist, fiblist[i-2]+fiblist[i-1])
//	}
//	return fiblist, nil
//}
//
//func main() {
//	list, err := GetFib(1)
//	// 检测错误可通过特定的一个变量
//	if err == errLess {
//		s := fmt.Errorf("need a bigger number \n")
//		fmt.Println(s)
//	}
//	if err == errBig {
//		_ = fmt.Errorf("need a smaller number")
//	}
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(list)
//}

// 二、 应避免错误的嵌套，有错误时提前短路如下面的GetFib2
//func GetFib1(str string) {
//	var (
//		i    int
//		err  error
//		list []int
//	)
//	if i, err = strconv.Atoi(str); err == nil {
//		if list, err = GetFib(i); err == nil {
//			println(list)
//		} else {
//			println(err)
//		}
//	} else {
//		println(err)
//	}
//}
//
//func GetFib2(str string) {
//	var (
//		i    int
//		err  error
//		list []int
//	)
//	if i, err = strconv.Atoi(str); err != nil {
//		println(err)
//		return
//	}
//
//	if list, err = GetFib(i); err != nil {
//		println(err)
//		return
//	}
//	println(list)
//}

// 三、 使用常用错误值
// 由于变量错误对象可被修改，可以变向的使用常量错误值
//type Error string
//
//func (e Error) Error() string {
//	return string(e)
//}
//
//const ErrEOF = Error("EOF")
//
//func test() error {
//	return ErrEOF
//}
//
//func main() {
//	err := test()
//	println(err == Error("EOF"))
//}

// 四、错误包装
// 通过对错误的包装，可以使其在不同的层级(上下文)呈现不同的内容
//type OrderError struct {
//	msg string
//}
//
//func (e *OrderError) Error() string {
//	return e.msg
//}
//
//func main() {
//	a := errors.New("a")
//	b := fmt.Errorf("b_%w", a)
//	c := fmt.Errorf("c_%w", b)
//	fmt.Println(c)
//
//	fmt.Println(errors.Unwrap(c) == b)
//	fmt.Println(errors.Is(c, a)) //递归匹配
//
//	x := &OrderError{"order"}
//	y := fmt.Errorf("y_%w", x)
//	z := fmt.Errorf("z_%w", y)
//	fmt.Println(z)
//
//	var x2 *OrderError
//	if errors.As(z, &x2) { // 按类型递归查找，并填写指针（指针的指针）
//		fmt.Println(x == x2)
//	}
//}

// 五、panic和普通退出
//func main() {
//	fmt.Println("start")
//	//os.Exit(-1)
//	panic(errors.New("sth wrong"))
//	fmt.Println("end")  // 执行不到这一行
//}

// 六、panic+recover
//var errNeg = fmt.Errorf("not a positive number")
//
//func fact(a int) int {
//	if a <= 0 {
//		panic(errNeg)
//	}
//	return a
//}
//
////异常捕获，但panic可以抛出错误以外的其他对象
//func main() {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println("error cached:", err)
//		}
//	}()
//	fmt.Println(fact(10))
//	fmt.Println(fact(5))
//	fmt.Println(fact(-5))
//	fmt.Println(fact(15))
//}

// 三、 redis的错误处理的例子，不轻易使用异常捕获
func main() {
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
