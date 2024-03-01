package main

func main() {
	// represent json in go with a map
	m := map[string]any{
		"Int":    1,
		"Float":  3.14,
		"String": "Hello, world!",
		"Bool":   true,
		"Array":  []any{1, 2, 3},
	}

	// access map
	value, ok := m["Float"]
	if !ok {
		println("Key not found")
		return
	}

	// type assertion
	f, ok := value.(float64)
	if !ok {
		println("Value is not a float64")
		return
	}

	result := f * 2
	println(result)
}