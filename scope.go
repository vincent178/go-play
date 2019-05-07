package main

import (
	"bytes"
	"fmt"
)

const a = func() string {
	return "Hello world"
}

func main() {
	x := "hello"

	for i := 0; i < len(x); i++ {
		i := 9
		fmt.Println(i)
		// x := x[i]
		// if x != '!' {
		// 	x := x + 'A' - 'a'
		// 	fmt.Printf("%c", x)
		// }
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	// b := true
	// fmt.Println(b ? 1 : 0)

	// s := "Hello world"
	// s[0] = "L"

	const GoUsage = `Go is a tool for managing Go source code.
Usage:
	go command [arguments]
`
	fmt.Println(GoUsage)

	fmt.Println(string("Hello"[1]))             // ASCII only
	fmt.Println(string([]rune("Hello, 世界")[1])) // UTF-8
	fmt.Println(string([]rune("Hello, 世界")[8])) // UTF-8

	fmt.Println(intsToString([]int{1, 2, 3}))

	a()
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
