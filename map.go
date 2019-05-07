package main

import "fmt"

func main() {
	ages := make(map[string]int)

	ages["bob"] += 1

	fmt.Println(ages["bob"])
}
