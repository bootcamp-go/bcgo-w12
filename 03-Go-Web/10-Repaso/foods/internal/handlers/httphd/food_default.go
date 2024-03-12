package httphd

import (
	"errors"
	"go-web/repaso/foods/internal"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type FoodJSON struct {
	// ID is an int that represents the id of the food
	ID int `json:"id"`
	// Name is a string that represents the name of the food
	Name string `json:"name"`
	// Weight is a float64 that represents the weight of the food
	Weight float64 `json:"weight"`
	// Ingridients is a slice of strings that represents the ingridients of the food
	Ingridients []string `json:"ingridients"`
	// Nationality is the country of origin of the food
	Nationality string `json:"nationality"`
	// Category is the category of the food
	Category string `json:"category"`
}

func (f *FoodJSON) ToFood() internal.Food {
	return internal.Food{
		ID:          f.ID,
		Name:        f.Name,
		Weight:      f.Weight,
		Ingridients: f.Ingridients,
		Nationality: f.Nationality,
		Category:    f.Category,
	}
}

func NewFoodDefault(service internal.FoodService) *FoodDefault {
	return &FoodDefault{
		service: service,
	}
}

type FoodDefault struct {
	// service is a FoodService that represents the service
	service internal.FoodService
}

func (f *FoodDefault) GetByName(w http.ResponseWriter, r *http.Request) {
	// request
	name := chi.URLParam(r, "name")

	// process
	food, err := f.service.GetByName(name)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrFoodNotFound):
			response.Text(w, http.StatusNotFound, "couldn't find the food")
		default:
			response.JSON(w, http.StatusInternalServerError, nil)
		}

		return
	}

	// response
	response.JSON(w, http.StatusOK, FoodJSON{
		ID:          food.ID,
		Name:        food.Name,
		Weight:      food.Weight,
		Ingridients: food.Ingridients,
		Nationality: food.Nationality,
		Category:    food.Category,
	})
}
