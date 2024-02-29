package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDivide(t *testing.T) {
	t.Run("success - happy path", func(t *testing.T) {
		// arrange - given
		n1 := 10
		n2 := 5
		expectedResult := 2
		expectedError := ""

		// act - when
		result, err := Divide(n1, n2)

		// assert - then
		if err != expectedError {
			// t.Fatalf("expected error %s, got %s", expectedError, err)
			t.Errorf("expected error %s, got %s", expectedError, err)
			return
		}
		if result != expectedResult {
			t.Errorf("expected result %d, got %d", expectedResult, result)
			return
		}
	})

	t.Run("success - both numbers are negative", func(t *testing.T) {
		// arrange - given

		// act - when

		// assert - then

	})

	t.Run("success - both numbers are positive", func(t *testing.T) {
		// arrange - given

		// act - when

		// assert - then

	})

	t.Run("success - not integer result", func(t *testing.T) {
		// arrange - given
		n1 := 9
		n2 := 10
		expectedResult := 0
		expectedError := ""

		// act - when
		result, err := Divide(n1, n2)

		// assert - then
		if err != expectedError {
			t.Errorf("expected error %s, got %s", expectedError, err)
			return
		}
		if result != expectedResult {
			t.Errorf("expected result %d, got %d", expectedResult, result)
			return
		}
	})

	t.Run("error - zero division", func(t *testing.T) {
		// arrange - given
		n1 := 10
		n2 := 0
		expectedResult := 0
		expectedError := ErrMsgDivisionByZero

		// act - when
		result, err := Divide(n1, n2)

		// assert - then
		if err != expectedError {
			t.Errorf("expected error %s, got %s", expectedError, err)
			return
		}
		if result != expectedResult {
			t.Errorf("expected result %d, got %d", expectedResult, result)
			return
		}
	})

	t.Run("error - indeterminate division", func(t *testing.T) {
		// arrange - given
		n1 := 0
		n2 := 0
		expectedResult := 0
		expectedError := ErrMsgDivisionIndeterminate

		// act - when
		result, err := Divide(n1, n2)

		// assert - then
		if err != expectedError {
			t.Errorf("expected error %s, got %s", expectedError, err)
			return
		}
		if result != expectedResult {
			t.Errorf("expected result %d, got %d", expectedResult, result)
			return
		}
	})
}

// func TestDivide_HappyPath(t *testing.T) {

// }

// func TestDivide_ErrZeroDivision(t *testing.T) {
	
// }

// func TestDivide_ErrIndeterminateDivision(t *testing.T) {
	
// }

func TestMultiply(t *testing.T) {
	t.Run("success - multiply two positive numbers", func(t *testing.T) {
		// arrange - given
		n1 := 10
		n2 := 5
		expectedResult := 50
		expectedError := ""

		// act - when
		result, err := Multiply(n1, n2)

		// assert - then
		require.Equal(t, expectedResult, result)
		require.Equal(t, expectedError, err)
	})

	t.Run("success - multiply two negative numbers", func(t *testing.T) {
		// arrange - given
		n1 := -10
		n2 := -10
		expectedResult := 100
		expectedError := ""

		// act - when
		result, err := Multiply(n1, n2)

		// assert - then
		require.Equal(t, expectedResult, result)
		require.Equal(t, expectedError, err)
	})

	t.Run("success - multiply by zero", func(t *testing.T) {
		// arrange - given
		n1 := 10
		n2 := 0
		expectedResult := 0
		expectedError := ""

		// act - when
		result, err := Multiply(n1, n2)

		// assert - then
		require.Equal(t, expectedResult, result)
		require.Equal(t, expectedError, err)
	})
}
