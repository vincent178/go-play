package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string, 2)

	go func() {
		messages <- "ping"
		time.Sleep(1 * time.Second)
		messages <- "pong"
	}()

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
