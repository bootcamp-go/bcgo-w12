package internal

type FoodCategory int

const (
	// FoodCategoryVegetable is a constant for food category
	FoodCategoryVegetable FoodCategory = iota
	// FoodCategoryFruit is a constant for food category
	FoodCategoryFruit
	// FoodCategoryMeat is a constant for food category
	FoodCategoryMeat
	// FoodCategoryDairy is a constant for food category
	FoodCategoryDairy
)

// Food is a struct that represents a food item
type Food struct {
	// ID is an int that represents the id of the food
	ID int
	// Name is a string that represents the name of the food
	Name string
	// Weight is a float64 that represents the weight of the food
	Weight float64
	// Ingridients is a slice of strings that represents the ingridients of the food
	Ingridients []string
	// Nationality is the country of origin of the food
	Nationality string
	// Category is the category of the food
	Category string
}