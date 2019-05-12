package main

import "fmt"

type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }

func main() {
	op := Point.Add

	p1 := Point{1, 2}
	p2 := Point{3, 4}

	fmt.Println(op(p1, p2).X)
}
