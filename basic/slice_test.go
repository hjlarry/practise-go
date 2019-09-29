package basic

import "testing"

func TestSlice(t *testing.T) {
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

	t.Log(names, scores)

	var names1 = make([]string, len(fruits))
	var scores1 = make([]int, len(fruits))
	for name, score := range fruits {
		names1 = append(names1, name)
		scores1 = append(scores1, score)
	}

	t.Log(names1, scores1)

}
