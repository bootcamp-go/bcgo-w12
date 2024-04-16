package handler

import (
	"go-testing/http-api/internal"
	"go-testing/http-api/internal/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSongHandler_Create(t *testing.T) {
	t.Run("success - create song", func(t *testing.T) {
		// given
		// - service
		svMock := service.NewMockSong()
		svMock.On("AddSong", &internal.Song{
			ID:	 0,
			Title:  "Higher Power",
			Artist: "Coldplay",
			Album:  "Music of the Spheres",
		}).Return(error(nil))
		// mock.AnythingOfType("*internal.Song") // to check only by type
		// mock.Anything // joker for the arg
		// - handler
		hd := NewSongHandler(svMock)
		
		// when
		request := httptest.NewRequest(http.MethodPost, "/songs", strings.NewReader(
			`{"title": "Higher Power", "artist": "Coldplay", "album": "Music of the Spheres"}`,
		))
		response := httptest.NewRecorder()
		hd.Create(response, request)

		// then
		expectedCode := http.StatusCreated
		expectedBody := `{"data":{"id":0,"title":"Higher Power","artist":"Coldplay","album":"Music of the Spheres"}, "message":"song created successfully"}`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		svMock.AssertExpectations(t)
	})

	t.Run("fail - missing title", func(t *testing.T) {
		// given

		// when

		// then

	})

	t.Run("fail - song already exists", func(t *testing.T) {
		// given

		// when

		// then

	})
}