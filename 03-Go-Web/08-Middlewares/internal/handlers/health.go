package handlers

import (
	"net/http"
	"time"
)

// HealthCheck is a simple health check handler
// - indicates if the server is up and running
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	time.Sleep(500 * time.Millisecond)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}