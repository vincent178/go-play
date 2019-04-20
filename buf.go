package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {

	var b bytes.Buffer

	b.Write([]byte("Hello world"))

	p := make([]byte, 100)

	fmt.Println("p length ", len(p))

	n, err := b.Read(p)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes\n", n)

	fmt.Println(string(p))
}
