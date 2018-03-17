package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // without initial value should be zero-valued
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}
