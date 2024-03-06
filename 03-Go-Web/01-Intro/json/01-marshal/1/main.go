package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title       string   `json:"title,omitempty"`
	ReleaseYear int      `json:"release_year,omitempty"`
	Director    string   `json:"director,omitempty"`
	Actors      []string `json:"actors,omitempty"`
	Password	string   `json:"-"`
	field       string
}

func main() {
	movie := Movie{
		Title:       "Inception",
		ReleaseYear: 2010,
		Director:    "Christopher Nolan",
		Actors:      []string{"Leonardo DiCaprio", "Joseph Gordon-Levitt", "Elliot Page"},
		Password:	"pass",
		field:		"field",
	}

	// Marshal the movie struct into a JSON object
	bytes, err := json.Marshal(movie)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the JSON object
	fmt.Println(string(bytes))
}
