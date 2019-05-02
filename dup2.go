package main

import "os"

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, args := range files {

		}
	}
}

func countLines(f *os.File, counts map[string]int) {

}
