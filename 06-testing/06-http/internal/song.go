package internal

// Song represents a song in the music library.
type Song struct {
	// ID is a unique identifier for the song.
	ID     int
	// Title is the title of the song.
	Title  string
	// Artist is the name of the artist who performed the song.
	Artist string
	// Album is the name of the album on which the song appears.
	Album  string
}