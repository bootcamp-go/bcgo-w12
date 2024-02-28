package main

import "fmt"

func main() {
	// slice: initialize by default (zero value -> nil)
	// - length: 0
	// - capacity: 0
	var fruits []string
	if fruits == nil {
		fmt.Println("fruits is nil")
	}
	fmt.Printf("fruits: length=%d, capacity=%d\n", len(fruits), cap(fruits))

	// slice: initialize native (empty)
	// - length: 0
	// - capacity: 0
	var colors []string = []string{}
	fmt.Printf("colors: length=%d, capacity=%d\n", len(colors), cap(colors))

	// slice: initialize natively (some)
	// - length: 3
	// - capacity: 3
	var names []string = []string{
		"John",
		"Jane",
		"Joe",
	}
	fmt.Printf("names: length=%d, capacity=%d\n", len(names), cap(names))

	// slice: make
	// - length: 0
	// - capacity: 8
	var numbers []int = make([]int, 0, 5)
	fmt.Printf("numbers: length=%d, capacity=%d\n", len(numbers), cap(numbers))

	// slice: append
	// numbers[1] = 10 // panic: runtime error: index out of range [1] with length 0
	numbers = append(numbers, 1, 2, 3, 4, 5)
	fmt.Printf("numbers: length=%d, capacity=%d\n", len(numbers), cap(numbers))

	// slice: append (re-allocation) -> exceed capacity
	// - length: 8
	// - capacity: 10
	numbers = append(numbers, 6, 7, 8)
	fmt.Printf("numbers: length=%d, capacity=%d\n", len(numbers), cap(numbers))

	// slice: append (re-allocation) -> exceed capacity
	// - length: 12
	// - capacity: 20
	numbers = append(numbers, 9, 10, 11, 12)
	fmt.Printf("numbers: length=%d, capacity=%d\n", len(numbers), cap(numbers))

	// access slice
	fmt.Println("numbers from position 0 to 4:", numbers[0:5])
}