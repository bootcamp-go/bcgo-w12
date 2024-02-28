package main

import "time"

func main() {
	// infinite loop
	// for true {
	for {
		// do something
		time.Sleep(time.Millisecond * 500)
		println("Infinite loop")
	}
}