package pipe_filter

import (
	"testing"
)

func TestPipeline(t *testing.T) {
	pipe := NewStraigtPipeline("mypipe", NewSplitFilter("|"), NewToIntFilter(), NewSumFilter())
	resp, err := pipe.Process("1|3|4")
	if err != nil {
		t.Fatal(err)
	}
	if resp != 8 {
		t.Fatalf("expected is 8, but the actual is %d", resp)
	}

}
