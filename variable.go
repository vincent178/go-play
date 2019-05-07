package main

import (
	"flag"
	"fmt"
)

func main() {
	x := 1
	p := &x
	fmt.Println(*p)

	*p = 2
	fmt.Println(*p)

	n := flag.Bool("n", false, "omit trailing newline")
}
