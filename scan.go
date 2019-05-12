package main

import "fmt"

func main() {
	data := "100%"
	var num int
	var percent string
	fmt.Sscanf(data, "%d%s", &num, &percent)
	fmt.Printf("number is %d%s\n", num, percent)
}
