package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// handlers
	handlerHome := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}

	handlerHealth := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}

	// router
	router := chi.NewRouter()
	router.Get("/", handlerHome)
	router.Get("/health", handlerHealth)

	// Start the server
	if err := http.ListenAndServe(":8080", router); err != nil {
		println("Error starting server: ", err.Error())
		return
	}
}