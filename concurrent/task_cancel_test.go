package concurrent

import (
	"testing"
	"time"
)

func isCanceled(cancelChan chan int) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

// 只能取消一个任务
func cancel1(cancelChan chan int) {
	cancelChan <- 0
}

// 可取消全部任务
func cancel2(cancelChan chan int) {
	close(cancelChan)
}

func TestTaskClose(t *testing.T) {
	cancelChan := make(chan int)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan int) {
			for {
				if isCanceled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			println(i, "cancelled")
		}(i, cancelChan)
	}
	// cancel1(cancelChan)
	cancel2(cancelChan)

	time.Sleep(time.Second)
}
