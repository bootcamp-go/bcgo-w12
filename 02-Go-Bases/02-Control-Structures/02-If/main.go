package main

func main() {
	var age int = 25

	// if age > 18 {
	// 	println("You are an adult")
	// } else {
	// 	println("You are a minor")
	// }

	// unique path condition
	// if age > 16 {
	// 	println("You can drive a car")
	// } else if age > 18 {
	// 	println("You can vote")
	// } else if age > 21 {
	// 	println("You can drink")
	// } else {
	// 	println("Unknown age")
	// }

	age = 20

	// multiple path conditions
	if age > 16 {
		println("You can drive a car")
	}
	if age > 18 {
		println("You can vote")
	}
	if age > 21 {
		println("You can drink")
	}
}