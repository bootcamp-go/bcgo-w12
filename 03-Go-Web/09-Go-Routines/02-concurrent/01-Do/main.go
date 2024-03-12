package main

import "time"

func Do(n1, n2 int) {
	time.Sleep(250 * time.Millisecond)
	result := n1 + n2
	println(result)
}

func main() {
	// set start time
	start := time.Now()

	// operation 1
	n1, n2 := 2, 3
	go Do(n1, n2)

	// operation 2
	n1, n2 = 5, 7
	go Do(n1, n2)

	// operation 3
	n1, n2 = 11, 13
	go Do(n1, n2)

	// lets wait for the goroutines to finish
	time.Sleep(260 * time.Millisecond)

	// set end time
	elapsed := time.Since(start)
	println("Elapsed time: ", elapsed.Milliseconds(), "ms")
}