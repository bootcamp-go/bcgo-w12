package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pre-processing
		start := time.Now()

		// call the next handler / middleware
		next.ServeHTTP(w, r)

		// post-processing
		elapsed := time.Since(start)
		fmt.Printf("%s %s | time: %d\n", r.Method, r.URL, elapsed.Milliseconds())
	})
}