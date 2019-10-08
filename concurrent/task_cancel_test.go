package concurrent

import (
	"context"
	"testing"
	"time"
)

func isCancelled(cancelChan chan int) bool {
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
				if isCancelled(cancelChan) {
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

// 使用context取消任务，则可以取消该任务的子任务
func isCancelledByCtx(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestTaskCtxClose(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelledByCtx(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			println(i, "cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second)
}
