package main

import (
	"encoding/json"
	"fmt"
)

type D struct {
	A string
	a string
}

func main() {
	data := "{\"a\":\"a\"}"

	var d D
	json.Unmarshal([]byte(data), &d)

	fmt.Println(d.a)
}
