package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int	`json:"age"`
}

func main() {
	// writer
	var writter io.Writer = os.Stdout

	// encoder
	encoder := json.NewEncoder(writter)

	// person
	p := Person{
		Name: "John",
		Age: 30,
	}
	if err := encoder.Encode(p); err != nil {
		fmt.Println(err)
		return
	}
}