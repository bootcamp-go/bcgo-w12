package main

import "fmt"

func main() {
	// array: initialize by default (zero value)
	var names [3]string // ["", "", ""]
	// var active [3]bool  // [false, false, false]
	fmt.Println(names)

	// array: initialize with value
	var temps [5]float64 = [5]float64{
		1.2,
		1.3,
		1.4,
		1.5,
	}
	temps[4] = 1.6
	fmt.Println(temps)
}