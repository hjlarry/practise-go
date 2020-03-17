package main

import "sync"

var m sync.Mutex

func A() {
	m.Lock()
	defer m.Unlock()
	B()
}
func B() {
	m.Lock()
	defer m.Unlock()
	println("hello")
}
func main() {
	A()
}

/*
不支持递归锁(可重入锁)

[ubuntu] ~/.mac/gocode $ go run lock.go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_SemacquireMutex(0x4e1014, 0x4c6600, 0x1)
	/usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0x4e1010)
	/usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
	/usr/local/go/src/sync/mutex.go:81
main.B()
	/root/.mac/gocode/lock.go:13 +0xbf
main.A()
	/root/.mac/gocode/lock.go:10 +0x71
main.main()
	/root/.mac/gocode/lock.go:18 +0x20
exit status 2

*/
