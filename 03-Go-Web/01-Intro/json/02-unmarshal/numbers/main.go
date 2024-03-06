package main

import "encoding/json"

func main() {
	jsonData := []byte(`{
		"name": "John",
		"age": 30,
		"height": 1.75,
		"weight": 80.0
	}`)

	m := make(map[string]any)
	if err := json.Unmarshal(jsonData, &m); err != nil {
		println(err.Error())
		return
	}

	age, ok := m["age"]
	if !ok {
		println("age not found")
		return
	}

	ageFloat, ok := age.(float64)
	if !ok {
		println("age is not an int")
		return
	}

	ageInt := int(ageFloat)

	ageInt++
	println(ageInt)
}