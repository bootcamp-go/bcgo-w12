package service

import "go-bases/composition/internal"

var (
	// ErrMsgNoMovies is the error message when there are no movies
	ErrMsgNoMovies = "there are no movies"
)

// NewMovieDefault creates a new movie default service
func NewMovieDefault(rp internal.MovieRepository) MovieDefault {
	return MovieDefault{rp}
}

// MovieDefault is the default movie service
type MovieDefault struct {
	// rp is the movie repository
	rp internal.MovieRepository
}

func (m MovieDefault) AverageRating() (avg float64, err string) {
	// get all the movies
	movies := m.rp.Get()
	
	// check if there are no movies
	if len(movies) == 0 {
		err = ErrMsgNoMovies
		return
	}

	// calculate the average rating
	var sum float64
	for _, movie := range movies {
		sum += movie.Rating
	}

	avg = sum / float64(len(movies))
	return
}