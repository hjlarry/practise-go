package main

import "testing"

func TestAdd2(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{1, 2, 3},
		{2, 3, 6},
	}

	for _, tt := range tests {
		o := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()
			got := add(o.x, o.y)
			if got != o.want {
				t.Errorf("add(%d, %d): want %d, got %d", o.x, o.y, o.want, got)
			}
		})
	}
}
