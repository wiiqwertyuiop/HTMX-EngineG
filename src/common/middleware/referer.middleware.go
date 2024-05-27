package my_middleware

import (
	"net/http"
)

// HTTP middleware setting a value on the request context
func RefererMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Referer() == "" {
			http.Error(w, "404 page not found", 404)
			return
		}
		next.ServeHTTP(w, r)
	})
}
