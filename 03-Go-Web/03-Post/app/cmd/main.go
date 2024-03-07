package main

import (
	"fmt"
	"go-web/post/app/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// dependencies
	// - handler
	handler := handlers.NewDefaultTask(nil, 0)
	// - router
	router := chi.NewRouter()
	router.Post("/tasks", handler.Create())

	// server
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println(err)
		return
	}
}