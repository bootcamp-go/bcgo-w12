package internal

import "errors"

var (
	ErrSongNotFound = errors.New("song not found")
	ErrSongExists   = errors.New("song already exists")
)

// SongService provides access to the music library.
type SongService interface {
	// GetSong returns the song with the given ID.
	GetSong(id int) (Song, error)
	// AddSong adds a new song to the music library.
	AddSong(song *Song) error
}
