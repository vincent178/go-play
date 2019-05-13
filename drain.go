package main

import (
	"fmt"
	"time"
)

func main() {
	sem := make(chan int, 10)

	for i := 0; i < 10; i++ {
		sem <- i
	}

	go func() {
		for i := range sem {
			fmt.Println(i)
		}
	}()

	go func() {
		// tick := time.Tick(1 * time.Second)
		for {
			sem <- 100
			// select {
			// case <-tick:
			// }
		}
	}()

	select {
	case <-time.After(1 * time.Millisecond):
	}
}
