package main

import "fmt"

func main() {
	// unique path condition
	var age int = 25

	// multiple path condition
	switch {
	case age > 21:
		println("You can drink")
		// println("You can vote")
		// println("You can drive a car")
		fallthrough
	case age > 18:
		println("You can vote")
		// println("You can drive a car")
		fallthrough
	case age > 16:
		println("You can drive a car")
	default:
		println("not sure what you can do with that age")
	}

	// example of multiples numbers
	number := 60
	switch {
	case number%21 == 0:
		fmt.Printf("%d is multiple of 20\n", number)
	case number%30 == 0:
		fmt.Printf("%d is multiple of 15\n", number)
		fallthrough
	case number%15 == 0:
		fmt.Printf("%d is multiple of 5\n", number)
		fallthrough
	case number%5 == 0:
		fmt.Printf("%d is multiple of 3\n", number)
	default:
		fmt.Println("No condition met")
	}
}