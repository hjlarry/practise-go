package concurrent

import (
	"github.com/pkg/errors"
	"testing"
	"time"
)

type ReusableObj struct {
}

type ObjPool struct {
	bufferChan chan *ReusableObj
}

func NewObjPool(number int) *ObjPool {
	bufferChan := make(chan *ReusableObj, number)
	objPool := ObjPool{bufferChan: bufferChan}
	for i := 0; i < number; i++ {
		bufferChan <- new(ReusableObj)
	}
	return &objPool
}

func (p *ObjPool) GetReusableObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufferChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufferChan <- obj:
		return nil
	default:
		return errors.New("full")
	}

}

func TestReusableObjPool(t *testing.T) {
	pool := NewObjPool(5)
	for i := 0; i < 5; i++ {
		obj, err := pool.GetReusableObj(time.Second)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%T", obj)
	}

	for i := 0; i < 6; i++ {
		t.Log(pool.ReleaseObj(new(ReusableObj)))
	}
}
