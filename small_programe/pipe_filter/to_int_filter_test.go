package pipe_filter

import (
	"reflect"
	"testing"
)

func TestToInt(t *testing.T) {
	ti := NewToIntFilter()
	resp, err := ti.Process([]string{"1", "2", "3"})
	if err != nil {
		t.Fatal(err)
	}
	ret, ok := resp.([]int)
	if !ok {
		t.Fatalf("Repsonse type is %T, but the expected type is []int", ret)
	}
	if !reflect.DeepEqual(ret, []int{1, 2, 3}) {
		t.Errorf("Expected value is []int, but actual is %v", ret)
	}
}

func TestWrongInputForToIntFilter(t *testing.T) {
	ti := NewToIntFilter()
	_, err := ti.Process(123)
	if err == nil {
		t.Fatal("An error is expected.")
	}
}
