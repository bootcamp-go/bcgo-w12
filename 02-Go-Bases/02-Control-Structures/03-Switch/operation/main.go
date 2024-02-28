package main

func main() {
	// command
	var number int = 10

	// check if number is odd or even
	switch r := number % 2; r {
	case 0:
		println("Even")
	default:
		println("Odd")
	}
}