package calculator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrchestrator(t *testing.T) {
	t.Run("success - returns Add function", func(t *testing.T) {
		// arrange
		op := OperatorSum
		expectedPtr := reflect.ValueOf(Add).Pointer()
		expectedErr := ""

		// act
		fn, err := Orchestrator(op)
		fnPtr := reflect.ValueOf(fn).Pointer()

		// assert
		require.Equal(t, expectedErr, err)
		require.Equal(t, expectedPtr, fnPtr)
	})

	t.Run("failure - returns error message", func(t *testing.T) {
		// arrange
		op := MathOperator(5)
		expectedErr := ErrMsgInvalidOperator

		// act
		fn, err := Orchestrator(op)

		// assert
		require.Equal(t, expectedErr, err)
		require.Nil(t, fn)
	})
}