package main

import "testing"

func add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		t.Fatal("add(1,2)!=3")
	}
}
