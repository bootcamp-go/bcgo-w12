package main

import "net/http"

func main() {
	// handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	// handler := func (w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Add("Content-Type", "text/plain")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte("Hello, World!"))
	// }

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		println("Error starting server: ", err.Error())
		return
	}
}