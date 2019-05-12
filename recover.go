package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	nesttest()
}

func test() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("internal error: %v\n", p)
			fmt.Printf("stack trace: %s", debug.Stack())
		}
	}()

	panic("something wrong")
}

func nesttest() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("internal error: %v\n", p)
			fmt.Printf("stack trace: %s", debug.Stack())
		}
	}()

	f := func() {
		panic("lala")
	}

	f()
}
