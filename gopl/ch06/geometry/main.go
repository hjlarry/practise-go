package main

import (
	"fmt"
	"math"
)

// Point ...
type Point struct {
	X, Y float64
}

// 使用函数
func distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 使用方法
func (p Point) distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 5}
	q := Point{6, 9}
	fmt.Println(distance(p, q))
	fmt.Println(p.distance(q))
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.distance())
}

// Path ...
type Path []Point

func (path Path) distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].distance(path[i])
		}
	}
	return sum
}
