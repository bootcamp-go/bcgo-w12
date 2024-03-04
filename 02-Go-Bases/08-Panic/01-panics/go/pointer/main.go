package main

func main() {
	var number *int

	println(*number) // panic: runtime error: invalid memory address or nil pointer dereference
}