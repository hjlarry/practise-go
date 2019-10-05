package basic

import (
	"fmt"
	"math"
	"testing"
)

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

// TestMethod ...
func TestMethod(t *testing.T) {
	r1 := rectangle{12, 2}
	c1 := circle{10}
	t.Log(r1.area())
	t.Log(c1.area())
}

// color
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type color byte
type box struct {
	width, height, depth float64
	color                color
}
type boxlist []box

func (b *box) volume() float64 {
	return b.width * b.height * b.depth
}

func (b *box) setColor(c color) {
	b.color = c
}

func (bl boxlist) biggestColor() color {
	v := 0.00
	k := color(WHITE)
	for _, b := range bl {
		if bv := b.volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl boxlist) PaintItBlack() {
	for i := range bl {
		bl[i].setColor(BLACK)
	}
}

func (c color) string() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

// TestMethod2 ...
func TestMethod2(t *testing.T) {
	boxes := boxlist{
		box{4, 4, 4, RED},
		box{10, 2, 10, BLUE},
		box{5, 5, 20, YELLOW},
		box{1, 4, 4, BLACK},
		box{4, 30, 4, WHITE},
	}
	t.Logf("we have %d boxes", len(boxes))
	t.Log("first box volumn:", boxes[0].volume(), "cm3")
	t.Log("last box color:", boxes[len(boxes)-1].color.string())
	t.Log("biggest:", boxes.biggestColor().string())

	t.Log("paint all to black:")
	boxes.PaintItBlack()
	t.Log("last box color:", boxes[len(boxes)-1].color.string())
}

type human struct {
	name string
	age  int
}

type student1 struct {
	human
	school string
}

type employee1 struct {
	human
	company string
}

func (h human) sayHi() {
	fmt.Println("this is human", h.name)
}

func (e employee1) sayHi() {
	fmt.Println("this is employee", e.name, e.company)
}

// TestMethod3 ... method继承和重写
func TestMethod3(t *testing.T) {
	mark := student1{human{"mark", 30}, "MIT"}
	sam := employee1{human{"sam", 30}, "Golang Inc."}

	mark.sayHi()
	sam.sayHi()
}