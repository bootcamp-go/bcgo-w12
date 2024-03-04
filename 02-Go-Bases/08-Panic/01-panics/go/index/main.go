package main

func main() {
	people := []string{"Alice", "Bob", "Charlie"}
	println(people[4]) // panic: runtime error: index out of range
}