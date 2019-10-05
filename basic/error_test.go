package basic

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/go-redis/redis"
)

var lessError = errors.New("the number should not less than 2")
var bigError = errors.New("the number should not big than 100")

// 一、 一般的错误处理方式
func GetFib(n int) ([]int, error) {
	if n < 2 {
		return nil, lessError
	}
	if n > 100 {
		return nil, bigError
	}
	fiblist := []int{1, 1}
	for i := 2; i < n; i++ {
		fiblist = append(fiblist, fiblist[i-2]+fiblist[i-1])
	}
	return fiblist, nil
}

func TestFib(t *testing.T) {
	list, err := GetFib(101)
	if err == lessError {
		t.Error("need a bigger number")
	}
	if err == bigError {
		t.Error("need a smaller number")
	}
	if err != nil {
		t.Log(err)
	}
	t.Log(list)
}

// 二、 应避免错误的嵌套，有错误时提前短路如下面的GetFib2
func GetFib1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFib(i); err == nil {
			println(list)
		} else {
			println(err)
		}
	} else {
		println(err)
	}
}

func GetFib2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		println(err)
		return
	}

	if list, err = GetFib(i); err != nil {
		println(err)
		return
	}
	println(list)
}

// 三、 redis的错误处理的例子，不轻易使用异常捕获
func TestRedisConn(t *testing.T) {
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

// 四、panic和普通退出
func TestPanicExit(t *testing.T) {
	t.Log("start")
	// os.Exit(-1)
	panic(errors.New("sth wrong"))
	t.Log("end")
}

// 五、panic+recover
var errNeg = fmt.Errorf("not a positive number")

func fact(a int) int {
	if a <= 0 {
		panic(errNeg)
	}
	return a
}

//异常捕获，但panic可以抛出错误以外的其他对象
func TestPanicRecover(t *testing.T) {
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
