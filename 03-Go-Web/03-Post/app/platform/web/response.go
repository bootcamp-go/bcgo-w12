package web

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, code int, data any) {
	// Set the content type
	w.Header().Set("Content-Type", "application/json")

	// Set the status code
	w.WriteHeader(code)

	// Encode the data
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
}