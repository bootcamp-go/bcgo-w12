package main

func main() {
	var m map[string]string

	// value := m["name"] // no panic
	// println(value)     // prints an empty string

	m["name"] = "Alice" // panic: assignment to entry in nil map
}