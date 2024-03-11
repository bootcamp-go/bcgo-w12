package application

import (
	"fmt"
	"go-web/middlewares/app/internal/handlers"
	"go-web/middlewares/app/internal/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewServerChi creates a new server application on chi
func NewServerChi(host, port string) *ServerChi {
	// default config
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "8080"
	}

	return &ServerChi{
		serverPort: port,
		serverHost: host,
	}
}

// ServerChi is the default configuration for the application
type ServerChi struct {
	// serverHost is the host of the server to listen on
	serverHost string
	// serverPort is the port of the server to listen on
	serverPort string
}

/*
	request
	| (w, r)
	router chi - pre-processing
	| (w, r)
	|
	| - detect endpoint: method + path
	|
	| - check middlewares: md_1, md_2
	|
	| - construct http.Handler (ServeHTTP)	<---------------------------------------|
	|   | handler -> http.HandlerFunc												|
	|   			 | - md_2(http.Handler) http.Handler							|
	|   			                        | - md_1(http.Handler) http.Handler (final)
	|   (w, r)																		|
	| - call the handler (ServeHTTP)												|
*/
func (d *ServerChi) Run() (err error) {
	// dependencies
	// - router
	rt := chi.NewRouter()
	// - middlewares
	rt.Use(middlewares.LoggerMiddleware, middlewares.AuthMiddleware)
	// - routes
	rt.Get("/health", handlers.HealthCheck)
	// rt.Route("/api", func(rt chi.Router) {
	// 	rt.Use(middlewares.AuthMiddleware)
	// 	rt.Get("/user", func(w http.ResponseWriter, r *http.Request) {
	// 		w.Write([]byte("user"))
	// 	})
	// })

	// start the server
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", d.serverHost, d.serverPort), rt)
	if err != nil {
		return
	}

	return
}