package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title       string   `json:"title,omitempty"`
	ReleaseYear int      `json:"ReLease_year,omitempty"`
	Director    string   `json:"director,omitempty"`
	Actors      []string `json:"actors,omitempty"`
	Password    string   `json:"-"`
}

func main() {
	jsonData := []byte(`{
		"title": "",
		"ReLease_year": 2010,
		"actors": ["Leonardo DiCaprio", "Joseph Gordon-Levitt", "Elliot Page"],
		"Password": "pass"
	}`)

	var movie Movie = Movie{
		Title:       "Inception",
		Director:   "Christopher Nolan",
	}
	if err := json.Unmarshal(jsonData, &movie); err != nil {
		println(err)
		return
	}

	println(movie.Title)
	println(movie.ReleaseYear)
	println(movie.Director)
	fmt.Println(movie.Actors)
	println(movie.Password)
}
