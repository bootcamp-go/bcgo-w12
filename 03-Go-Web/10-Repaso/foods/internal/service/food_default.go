package service

import "go-web/repaso/foods/internal"

var Categories = []string{"vegetable", "fruit", "meat", "dairy"}

func NewFoodDefault(repository internal.FoodRepository) internal.FoodService {
	return &FoodDefault{
		repository: repository,
	}
}

type FoodDefault struct {
	// repository is a FoodRepository that represents the repository
	repository internal.FoodRepository
	// apiMeLiFoodCategory is a map that represents the category of the food in the Mercado Libre API
	// apiMeLiFoodCategory internal.APIMeLiCategoryFood
}

func (f *FoodDefault) GetByName(name string) (food internal.Food, err error) {
	food, err = f.repository.GetByName(name)
	return
}