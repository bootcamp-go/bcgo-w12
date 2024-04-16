package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-testing/http-api/internal"
	"io"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

type SongJSON struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

// NewSongHandler creates a new SongHandler.
func NewSongHandler(songService internal.SongService) *SongHandler {
	return &SongHandler{SongService: songService}
}

// SongHandler provides HTTP handlers for working with songs.
type SongHandler struct {
	// SongService provides access to the music library.
	SongService internal.SongService
}

// Create adds a new song to the music library.
func (h *SongHandler) Create(w http.ResponseWriter, r *http.Request) {
	// request
	// - bytes
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to read request body")
		return
	}

	// - map
	var m map[string]any
	if err := json.Unmarshal(bytes, &m); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}
	// - validate
	if err := ValidateExistanceField(m, "title"); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// - struct
	var body SongJSON
	if err := json.Unmarshal(bytes, &body); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}
	song := internal.Song{
		ID:     body.ID,
		Title:  body.Title,
		Artist: body.Artist,
		Album:  body.Album,
	}

	// service
	if err := h.SongService.AddSong(&song); err != nil {
		switch {
		case errors.Is(err, internal.ErrSongExists):
			response.Error(w, http.StatusConflict, "song already exists")
		default:
			response.Error(w, http.StatusInternalServerError, "failed to add song")
		}
		return
	}

	// response
	data := SongJSON{
		ID:     song.ID,
		Title:  song.Title,
		Artist: song.Artist,
		Album:  song.Album,
	}
	response.JSON(w, http.StatusCreated, map[string]any{
		"data":    data,
		"message": "song created successfully",
	})
}

func ValidateExistanceField(m map[string]any, field string) error {
	if _, ok := m[field]; !ok {
		return fmt.Errorf("missing field %s", field)
	}
	return nil
}
