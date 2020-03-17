package main

// 执行过程：
// deferprocStack(println, "abc")
// deferprocStack(main.func1)
// 把两个deferprocStack加入到一个链表中去
// set result
// 使用deferReturn执行所有链表中的func，FILO的顺序
// return

func main() {
	defer println("abc")
	defer func() {}()
}

// // 延迟调用的误用
// func main() {
// 	for i := 0; i < 10000; i++ {
// 		path := fmt.Sprintf("./log/%d.txt", i)
// 		f, err := os.Open(path)
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 		defer f.Close() // 相当于链表中加入了10000个deferProc
// 	}
// }

// // 正确方式
// func main() {
// 	for i := 0; i < 10000; i++ {
// 		func() {
// 			path := fmt.Sprintf("./log/%d.txt", i)
// 			f, err := os.Open(path)
// 			if err != nil {
// 				log.Println(err)
// 				continue
// 			}
// 			defer f.Close()
// 		}()
// 	}
// }

// 大多数情况可以确保defer一定会被执行，例如panic、return等
// 但例如os.Exit会杀掉进程，defer不会执行
// func main() {
// 	defer println("exit")
// 	//os.Exit(0)
// 	panic("abc")
// }
