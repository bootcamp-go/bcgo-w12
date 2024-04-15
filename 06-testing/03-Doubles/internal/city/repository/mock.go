package repository

import "go-testing/doubles/internal/city"

// NewMock creates a new mock repository.
func NewMock() *Mock {
	return &Mock{
		FuncSaveCity: func(c *city.City) (err error) {
			return
		},
	}
}

type Mock struct {
	// FuncSaveCity is the function to save a city.
	FuncSaveCity func(c *city.City) (err error)

	// Calls
	Calls struct {
		// register the n times the method is called
		SaveCityCount int
		// register args from save city
		SaveCityArgs city.City
	}
}

func (m *Mock) SaveCity(c *city.City) (err error) {
	// registry
	m.Calls.SaveCityCount++
	m.Calls.SaveCityArgs = *c

	// expected call
	err = m.FuncSaveCity(c)
	return
}