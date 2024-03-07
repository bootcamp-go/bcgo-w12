package main

import (
	"go-web/get/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	h := handler.NewHandler()

	r.Get("/", h.Get())
	r.Get("/users/{userId}", h.GetById())
	r.Get("/users", h.GetByQuery())

	http.ListenAndServe(":8080", r)
}
