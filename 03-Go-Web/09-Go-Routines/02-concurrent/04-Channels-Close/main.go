package main

import (
	"math/rand"
	"time"
)

func NumberGenerator(ch chan<- int) {
	defer close(ch)
	for i := 0; i < rand.Intn(10); i++ {
		randomNumber := rand.Intn(100)
		ch <- randomNumber
	}
}

func main() {
	// create a channel
	ch := make(chan int)

	// start the goroutine
	go NumberGenerator(ch)

	// receive from the channel
	for {
		i, ok := <-ch
		if !ok {
			break
		}

		println(i)
	}

	go func (ch chan<- int) {
		ch <- 100
	}(ch)

	time.Sleep(250 * time.Millisecond)

	println("Done")
}