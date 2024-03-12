package internal

import "errors"

var (
	// ErrFoodInvalidIngridients is an error that is returned when the ingridients are invalid
	ErrFoodInvalidIngridients = errors.New("food ingridients are invalid")
	// ErrFoodInvalidCategory is an error that is returned when the category is invalid
	ErrFoodInvalidCategory = errors.New("food category is invalid")
)

type FoodService interface {
	// GetByName is a method that returns a food item by its name
	GetByName(name string) (Food, error)
}