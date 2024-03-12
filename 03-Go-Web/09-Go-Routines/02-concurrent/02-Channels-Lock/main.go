package main

func main() {
	// create a channel
	ch := make(chan int)

	// send a value to the channel
	ch <- 1

	// receive from the channel
	println(<-ch)
}