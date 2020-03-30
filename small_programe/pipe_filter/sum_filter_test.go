package pipe_filter

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	sm := NewSumFilter()
	resp, err := sm.Process([]int{1, 3, 5})
	if err != nil {
		t.Fatal(err)
	}
	ret, ok := resp.(int)
	if !ok {
		t.Fatalf("Repsonse type is %T, but the expected type is int", ret)
	}
	if !reflect.DeepEqual(ret, 9) {
		t.Errorf("Expected value is 9, but actual is %v", ret)
	}
}

func TestWrongInputForSumFilter(t *testing.T) {
	sm := NewSumFilter()
	_, err := sm.Process(123)
	if err == nil {
		t.Fatal("An error is expected.")
	}
}
