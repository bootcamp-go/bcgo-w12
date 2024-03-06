package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// reader
	rd := strings.NewReader(`{"name": "John", "age": 30}`)

	// decoder
	decoder := json.NewDecoder(rd)

	// person
	var p Person
	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p)
}