package main

import "fmt"

type PointA struct {
	X, Y int
}

func (p PointA) pp() {
	fmt.Println(p.X)
}

type PointB struct {
	X int
}

func (p PointB) pp() {
	fmt.Println(p.X)
}

type Circle struct {
	PointA
	PointB
}

func main() {
	c := new(Circle)
	c.PointA.X = 10
	c.PointB.X = 20
	fmt.Printf("%#v\n", c)
	c.PointA.pp()
	// fmt.Println((*c).PointA.pp())
}
