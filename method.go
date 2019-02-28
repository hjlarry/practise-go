package main

import (
	"fmt"
	"math"
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
func TestMethod() {
	r1 := rectangle{12, 2}
	c1 := circle{10}
	fmt.Println(r1.area())
	fmt.Println(c1.area())
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

func (b box) volume() float64 {
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
