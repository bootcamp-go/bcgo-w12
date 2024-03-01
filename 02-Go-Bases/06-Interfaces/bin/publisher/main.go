package main

import "fmt"

// Publisher is a struct that contains the name of the publisher and the names of the founders
type Publisher struct {
	// ID is the unique identifier of the publisher
	ID int
	// Name is the name of the publisher
	Name string
	// Founders is a slice of the names of the founders
	Founders []string
}

// TotalFounders returns the number of founders of the publisher
func (p Publisher) TotalFounders() int {
	return len(p.Founders)
}

// Book is an struct that contains the title, author and pages of a book
type Book struct {
	// ID is the unique identifier of the book
	ID int
	// Title is the title of the book
	Title string
	// Author is the author of the book
	Author string
	// Pages is the number of pages of the book
	Pages int
	// Publisher is the publisher of the book (embedded struct)
	Publisher
}

func (b Book) ShowInfo() string {
	return fmt.Sprintf(
		"%d: %s was written by %s and has %d pages - total publisher founders: %d",
		b.ID,
		b.Title,
		b.Author,
		b.Pages,
		b.TotalFounders(),
	)
}

// Movie is an struct that contains the title, director, imdb rating
type Movie struct {
	// ID is the unique identifier of the movie
	ID int
	// Title is the title of the movie
	Title string
	// Director is the director of the movie
	Director string
	// IMDBRating is the IMDB rating of the movie
	IMDBRating float64
	// Publisher is the publisher of the movie (embedded struct)
	Publisher
}

func (m Movie) ShowInfo() string {
	return fmt.Sprintf(
		"%d: %s was directed by %s and has a rating of %.1f - total publisher founders: %d",
		m.ID,
		m.Title,
		m.Director,
		m.IMDBRating,
		m.TotalFounders(),
	)
}

func main() {
	// Create a new publisher
	publisher := Publisher{
		Name: "Penguin Random House",
		Founders: []string{
			"Richard Lane",
			"John Lane",
		},
	}

	// Create a new book
	book := Book{
		Title:     "The Hitchhiker's Guide to the Galaxy",
		Author:    "Douglas Adams",
		Pages:     224,
		Publisher: publisher,
	}

	// Create a new movie
	movie := Movie{
		Title:      "Titanic",
		Director:   "James Cameron",
		IMDBRating: 7.8,
		Publisher:  publisher,
	}

	// Show the book info
	fmt.Println("book:", book.ShowInfo())

	// Show the movie info
	fmt.Println("movie:", movie.ShowInfo())
}
