package main

import "encoding/json"

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
}

func main() {
	// ________________________________________________________
	// first example
	p1 := Person{
		Name:   "John",
		Age:    30,
		Height: 1.75,
	}

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(p1)
	if err != nil {
		println(err)
		return
	}

	// Print the JSON data
	println(string(jsonData))

	// ________________________________________________________
	// second example
	p2 := Person{
		Name:   "Jane",
		Age:    25,
		Height: 2.0,
	}

	// Marshal the struct to JSON
	jsonData, err = json.Marshal(p2)
	if err != nil {
		println(err)
		return
	}

	// Print the JSON data
	println(string(jsonData))
}
