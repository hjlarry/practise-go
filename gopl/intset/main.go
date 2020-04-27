package main

import (
	"bytes"
	"fmt"
)

//一个bit数组, 每一个元素的每一位都表示集合里的一个值。
// 当集合的第i位被设置时，我们才说这个集合包含元素i。
// 下面的这个程序展示了一个简单的bit数组类型，并且实现了三个函数来对这个bit数组来进行操作
type intSet struct {
	words []uint64
}

func (s *intSet) has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *intSet) add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *intSet) unionWith(t *intSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *intSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y intSet
	x.add(1)
	x.add(133)
	x.add(89)
	fmt.Println(x.String())

	y.add(89)
	y.add(43)
	fmt.Println(y.String())

	x.unionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.has(43), y.has(133))
}
