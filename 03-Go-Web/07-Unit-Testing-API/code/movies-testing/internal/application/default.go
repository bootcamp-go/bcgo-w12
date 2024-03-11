package application

import (
	"go-web/unit-testing/app/internal/handler"
	"go-web/unit-testing/app/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewDefault returns a new Default.
func NewDefault(addr string) *Default {
	if addr == "" {
		addr = ":8080"
	}

	return &Default{
		addr: addr,
	}
}

// Default is an implementation of the Default interface.
type Default struct {
	// addr is the address to listen on.
	addr string
}

// Run runs the application.
func (a *Default) Run() (err error) {
	// dependencies
	// - repository
	rp := repository.NewTaskMap(nil)
	// - handler
	hd := handler.NewTaskDefault(rp)
	// - router
	rt := chi.NewRouter()
	rt.Get("/tasks/{id}", hd.GetTask())

	// run the server
	err = http.ListenAndServe(a.addr, rt)
	return
}