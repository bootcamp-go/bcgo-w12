package main

func main() {
	// explicit var declaration
	// var colorOne, colorTwo, colorThree string = "Red", "Green", "Blue"
	// var colorOne string = "Red"
	// var colorTwo string = "Green"
	// var colorThree string = "Blue"

	// implicit var declaration
	// - var
	// var colorOne, colorTwo, colorThree = "Red", "Green", "Blue"
	// var colorOne = "Red"
	// var colorTwo = "Green"
	// var colorThree = "Blue"

	// - :=
	// colorOne, colorTwo, colorThree := "Red", "Green", "Blue"
	colorOne := "Red"
	colorTwo := "Green"
	colorThree := "Blue"

	println(colorOne, colorTwo, colorThree)
}