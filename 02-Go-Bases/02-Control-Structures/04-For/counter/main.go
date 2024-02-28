package main

import "fmt"

func main() {
	// implement explicit break condition on for loop
	var counter int
	limit := 5

	for counter < limit {
		println("Counter:", counter)

		// increment counter
		// counter = counter + 1
		// counter += 1
		counter++
	}

	// break keyword
	counter = 0
	limit = 10
	for {
		// display counter
		println("Counter:", counter)

		// break condition
		if counter >= limit {
			break
		}

		// increment counter
		counter++
	}

	// continue keyword
	// ...

	// for loop - controlled iteration
	// lifetime of "i" is limited to the loop
	for i := 0; i < 10; i++ {
		println("Counter:", i)
	}

	// println("Counter:", i) // error: undefined: i

	// for loop - controlled iteration
	// lifetime of "i" keeps after the loop
	var i int
	for i = 0; i < 20; i += 5 {
		println("Counter:", i)
	}

	fmt.Println(i)
}