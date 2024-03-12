package main

import "time"

func Do(n1, n2 int, ch chan<- int) {
	time.Sleep(250 * time.Millisecond)
	result := n1 + n2
	ch <- result
}

// func Read(ch <-chan int) {
// 	result := <-ch
// 	println(result)
// }

func main() {
	// set start time
	start := time.Now()

	// create a channel
	// - capacity: 0 (unbuffered)
	ch := make(chan int, )

	// create a channel
	// - capacity: 3 (buffered)
	// ch := make(chan int, 3)

	// operation 1
	n1, n2 := 2, 3
	go Do(n1, n2, ch)

	// operation 2
	n1, n2 = 5, 7
	go Do(n1, n2, ch)

	// operation 3
	n1, n2 = 11, 13
	go Do(n1, n2, ch)

	time.Sleep(time.Second)

	// lets wait for the goroutines to finish
	// - receive from the channel
	// Read(ch)
	// Read(ch)
	// Read(ch)
	r1 := <-ch
	r2 := <-ch
	r3 := <-ch
	// r4 := <-ch  // fatal error: all goroutines are asleep - deadlock!

	println(r1, r2, r3)

	// set end time
	elapsed := time.Since(start)
	println("Elapsed time: ", elapsed.Milliseconds(), "ms")
}