package main

import (
	"fmt"
	"go-web/methods/app/internal/handlers"
	"go-web/methods/app/internal/repository"
	"go-web/methods/app/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// dependencies
	// - repository
	rp := repository.NewTaskMap(nil, 0)
	// - service
	sv := service.NewDefaultTask(rp)
	// - handler
	hd := handlers.NewDefaultTask(sv)
	// - router
	router := chi.NewRouter()
	router.Route("/tasks", func(r chi.Router) {
		// POST /tasks
		r.Post("/", hd.Create())
		// PUT /tasks/{id}
		r.Put("/{id}", hd.Update())
		// PATCH /tasks/{id}
		r.Patch("/{id}", hd.UpdatePartial())
		// DELETE /tasks/{id}
		r.Delete("/{id}", hd.Delete())
	})

	// server
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println(err)
		return
	}
}