package main

import "fmt"

type P struct {
	Name string
}

func (p *P) updateName(name string) {
	p.Name = name
}

func main() {
	p := P{"Vincent"}
	p.updateName("Jacky")
	fmt.Println("name is ", p.Name)
}
