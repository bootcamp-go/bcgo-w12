package repository

import "go-web/repaso/foods/internal"

func NewFoodSlice(db []internal.Food) *FoodSlice {
	// default values
	if db == nil {
		db = make([]internal.Food, 0)
	}
	
	return &FoodSlice{db: db}
}

type FoodSlice struct {
	// db is a slice of Food that represents the database
	db []internal.Food
	// lastId is an int that represents the last id
	lastID int
}

func (fs *FoodSlice) GetByName(name string) (food internal.Food, err error) {
	// Iterate over the database
	var exists bool
	for _, f := range fs.db {
		if f.Name == name {
			exists = true
			food = f
			break
		}
	}

	// check if the food exists
	if !exists {
		err = internal.ErrFoodNotFound
		return
	}

	return
}

func (fs *FoodSlice) Save(food *internal.Food) (err error) {
	// check if the food exists
	for _, f := range fs.db {
		if f.Name == food.Name {
			err = internal.ErrFoodAlreadyExists
			return
		}
	}

	// save the food
	// - increment the last id	
	fs.lastID++
	// - set the id
	food.ID = fs.lastID
	// - append the food to the database
	fs.db = append(fs.db, *food)

	return
}