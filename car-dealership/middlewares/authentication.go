package middlewares

import (
	"net/http"
)

// Authentication is the middleware to authenticate the request
func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Api-Key") != "srijan-zs" {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}

		next.ServeHTTP(w, r)
	})
}
