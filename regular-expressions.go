package main

import (
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p[a-z]+", "peach")

	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch pinch")) // find  first match
	fmt.Println(r.FindAllString("peach punch pinch", -1))  // find all match
	fmt.Println(r.FindAllString("peach punch pinch", 2))

}
