package intro

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeightCheck(t *testing.T) {
	t.Run("Can ride the roller coaster", func(t *testing.T) {
		// Arrange
		height := 175.0
		expectedMsg := "Can ride the roller coaster"

		// Act
		actualMsg := HeightCheck(height)

		// Assert
		require.Equal(t, expectedMsg, actualMsg)
	})

	// t.Run("Can't ride the roller coaster", func(t *testing.T) {
	// 	// Arrange
	// 	height := 145.0
	// 	expectedMsg := "Can't ride the roller coaster"

	// 	// Act
	// 	actualMsg := HeightCheck(height)

	// 	// Assert
	// 	require.Equal(t, expectedMsg, actualMsg)
	// })
}

func TestHeightCheck2(t *testing.T) {
	t.Run("Can't ride the roller coaster", func(t *testing.T) {
		// Arrange
		height := 145.0
		expectedMsg := "Can't ride the roller coaster"

		// Act
		actualMsg := HeightCheck(height)

		// Assert
		// panic("test")
		require.Equal(t, expectedMsg, actualMsg)
	})
}