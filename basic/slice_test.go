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

	var ss1 = make([]string, 3)
	var ss2 = make([]string, 3, 5)
	t.Log(len(ss1), cap(ss1))
	t.Log(len(ss2), cap(ss2))

	// 说明append必然会赋值给新的切片，不会对原切片有影响
	_ = append(ss1, "a")
	_ = append(ss1, "b")
	_ = append(ss1, "c")
	_ = append(ss1, "d")
	t.Log(ss1)

	// 任一切片中的值变化会导致底层数组和其他引用该数组的值的变化
	var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ss3 := arr[2:5]
	ss4 := arr[2:]
	ss3[2] = 0
	t.Log(ss3)
	t.Log(ss4)
	t.Log(arr)

}
