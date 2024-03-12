package internal

import "errors"

var (
	// ErrFoodNotFound is an error that is returned when the food is not found
	ErrFoodNotFound = errors.New("food not found")
	// ErrFoodAlreadyExists is an error that is returned when the food already exists
	ErrFoodAlreadyExists = errors.New("food already exists")
)

// FoodRepository is an interface that represents a food repository
type FoodRepository interface {
	// GetByName is a method that returns a food item by its name
	GetByName(name string) (Food, error)
}