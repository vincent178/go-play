package main

import "fmt"

func main() {
	done := make(chan struct{})
	close(done)
	select {
	case <-done:
		fmt.Println("get value from closed chan")
	default:
		fmt.Println("nothing from closed chan")
	}
}
