package middlewares

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something before
		token := r.Header.Get("Authorization")
		if token != "123" {
			response.Text(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// call the next handler
		next.ServeHTTP(w, r)

		// do something after
		// ...
	})
}