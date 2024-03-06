package main

import "net/http"

func main() {
	// handler
	handler := func (w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}

	// Start the server
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}