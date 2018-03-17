package main

import "fmt"

// Structs are mutable.
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Hello world")

	p := Person{"Vincent", 30}
	fmt.Println(p.Age)

	p.Age = 40
	fmt.Println(p.Age)

	sp := &p
	sp.Age = 50
	fmt.Println(sp.Age)

	p.updateAge(70)
	fmt.Println(p.Age)

	p.updateAge2(80)
	fmt.Println(p.Age)
}

func (p *Person) updateAge(newAge int) {
	p.Age = newAge
}

// this will not change struct, because it use copy for the caller
func (p Person) updateAge2(newAge int) {
	p.Age = newAge
}
