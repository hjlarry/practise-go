package main

import "fmt"

// TestSlice ...
func TestSlice() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}

	var names = make([]string, 0, len(fruits))
	var scores = make([]int, 0, len(fruits))

	for name, score := range fruits {
		names = append(names, name)
		scores = append(scores, score)
	}

	fmt.Println(names, scores)

	var names1 = make([]string, len(fruits))
	var scores1 = make([]int, len(fruits))
	for name, score := range fruits {
		names1 = append(names1, name)
		scores1 = append(scores1, score)
	}

	fmt.Println(names1, scores1)

}
