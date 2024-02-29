package main

func main() {
	// slice of integers
	sl := []int{1, 2, 3, 4, 5}

	var result = addition(sl...)
	println(result)
}

func addition(values ...int) int {
	var result int
	for _, value := range values {
		result += value
	}

	return result
}