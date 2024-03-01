package internal

import "time"

// MovieAttributes is a struct that represents the attributes of a movie
type MovieAttributes struct {
	Title    string
	Year     int
	Director string
	Genre    string
	Duration time.Duration
	Rating   float64
}

// Movie is a struct that represents a movie
type Movie struct {
	// ID is the unique identifier of the movie
	ID int
	// MovieAttributes is the attributes of the movie
	MovieAttributes
}