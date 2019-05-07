package main

import "fmt"

type a struct {
	bt   *b
	name string
}

type b struct {
	name string
}

func main() {
	ch := make(chan a, 1)

	aa := a{
		bt: &b{
			name: "that",
		},
		name: "yes",
	}

	ch <- aa

	ab := <-ch

	ab.name = "no"
	ab.bt.name = "this"

	fmt.Println(ab.name)
	fmt.Println(ab.bt.name)

	fmt.Println(aa.name)
	fmt.Println(aa.bt.name)

	// ch <- change(aa)
	// fmt.Println(aa.bt.name)
	// fmt.Println(aa.name)
}

func change(aa a) {
	aa.bt.name = "this"
	aa.name = "no"
}
