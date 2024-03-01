package internal

// MovieRepository is an interface that represents the movie repository
type MovieRepository interface {
	// Get returns all the movies
	Get() (mv []Movie)
}