package concurrent

import (
	"sync"
	"testing"
	"time"
	"unsafe"
)

type Singleten struct{}

var single *Singleten

var once sync.Once

func getSingleton() *Singleten {
	once.Do(func() {
		single = new(Singleten)
		println("create success")
	})
	return single
}

func TestOnceTask(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			sin := getSingleton()
			println(unsafe.Pointer(sin))
		}()
	}
	time.Sleep(time.Second)
}
