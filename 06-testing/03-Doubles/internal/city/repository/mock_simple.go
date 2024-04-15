package repository

import (
	"go-testing/doubles/internal/city"
	"go-testing/doubles/platform/mock"
)

// NewMockSimple creates a new mock repository.
func NewMockSimple() *MockSimple {
	return &MockSimple{
		Mock: mock.NewMock(),
	}
}

type MockSimple struct {
	mock.Mock
}

func (m *MockSimple) SaveCity(c *city.City) (err error) {
	// call
	output := m.Called("SaveCity", mock.Values{*c})
	err = output[0].(error)
	return

	// // registry
	// m.Calls["SaveCity"] = mock.Call{Ok: true, Args: mock.Args{c}}

	// // expected call
	// args := m.ExpectedCalls["SaveCity"]
	// err = args[0].(error)
	// return
}

// func (m *MockSimple) GetCity(id int) (c city.City, err error) {
// 	// registry
// 	m.Calls["GetCity"] = mock.Call{Ok: true, Args: mock.Args{id}}

// 	// expected call
// 	args := m.ExpectedCalls["GetCity"]
// 	c = args[0].(city.City)
// 	err = args[1].(error)
// 	return
// }