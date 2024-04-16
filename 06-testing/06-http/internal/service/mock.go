package service

import (
	"go-testing/http-api/internal"

	"github.com/stretchr/testify/mock"
)

// NewMockSong creates a new MockSong.
func NewMockSong() *MockSong {
	return &MockSong{}
}

// MockSong is a mock implementation of the SongService interface.
type MockSong struct {
	mock.Mock
}

// GetSong returns the values stored in the mock.
func (m *MockSong) GetSong(id int) (internal.Song, error) {
	args := m.Called(id)
	return args.Get(0).(internal.Song), args.Error(1)
}

// AddSong returns the values stored in the mock.
func (m *MockSong) AddSong(song *internal.Song) error {
	args := m.Called(song)
	return args.Error(0)
}