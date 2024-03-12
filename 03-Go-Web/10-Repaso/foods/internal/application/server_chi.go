package application

import (
	"go-web/repaso/foods/internal"
	"go-web/repaso/foods/internal/handlers/httphd"
	"go-web/repaso/foods/internal/repository"
	"go-web/repaso/foods/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewServerChi(address string) *ServerChi {
	return &ServerChi{Address: address}
}

type ServerChi struct {
	// Address is a string that represents the address of the server
	Address string
}

func (s *ServerChi) Start() (err error) {
	// initialize dependencies
	// - repository
	db := []internal.Food{
		{
			ID:          1,
			Name:        "apple",
			Weight:      0.2,
			Ingridients: []string{"water", "sugar", "fiber"},
			Nationality: "USA",
			Category:    "fruit",
		},
		{
			ID:          2,
			Name:        "banana",
			Weight:      0.3,
			Ingridients: []string{"water", "sugar", "fiber"},
			Nationality: "Ecuador",
			Category:    "fruit",
		},
	}
	rp := repository.NewFoodSlice(db)
	// - service
	sv := service.NewFoodDefault(rp)
	// - handler
	hd := httphd.NewFoodDefault(sv)
	// - router
	rt := chi.NewRouter()
	rt.Get("/foods/{name}", hd.GetByName)

	// - server
	err = http.ListenAndServe(s.Address, rt)
	return
}

// func (s *ServerChi) RegisterFood(rt *chi.Mux) {
// 	// dependencies
// 	// - repository
// 	rp := repository.NewFoodSlice(nil)
// 	// - service
// 	sv := service.NewFoodDefault(rp)
// 	// - handler
// 	hd := httphd.NewFoodDefault(sv)

// 	rt.Get("/foods/{name}", hd.GetByName)
// }