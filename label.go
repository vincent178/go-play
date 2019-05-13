package main

import "fmt"

func main() {
	fmt.Printf("test\n")

ll:
	fmt.Println("jj")
	for i := 0; i < 3; i++ {
		if i == 2 {
			goto ll
		}
		fmt.Println(i)
	}
	fmt.Println("Hello")
}
