package repository

import "go-bases/composition/internal"

// bad practice
// - complex to test
// - data is not encapsulated
// - not dynamic
// var movies []internal.Movie
// var limit int = 10
// var count int

// NewMoviesSlice creates a new movies slice repository
func NewMoviesSlice(movies []internal.Movie) MoviesSlice {
	return MoviesSlice{movies}
}

// MoviesSlice is a slice of movies
type MoviesSlice struct{
	// movies is a slice of movies
	movies []internal.Movie
	// // limit is the maximum number of movies
	// limit int
	// // count is the number of movies
	// count int
}

// Get returns all the movies
func (m MoviesSlice) Get() (mv []internal.Movie) {
	// make a copy of the movies slice
	// - deep copy
	mv = make([]internal.Movie, len(m.movies))
	copy(mv, m.movies)
	// for i, m := range movies {
	// 	mv[i] = m
	// }

	return
}