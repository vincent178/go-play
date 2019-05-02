package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

func main() {

	num := 704912
	// num := 12

	c := make(chan string, num)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c <- "silly name #" + strconv.Itoa(rand.Intn(100))
		}()
	}

	f, err := os.Create("/tmp/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)

	go func() {
		wg.Wait()
		close(c)
	}()

	for name := range c {
		if err := w.Write([]string{name}); err != nil {
			log.Fatalln("err writing record to csv: ", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

}
