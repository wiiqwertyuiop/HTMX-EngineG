package my_middleware

import (
	"htmx-engineg/src/common/renderer"
	"net/http"
)

// HTTP middleware setting a value on the request context
func RefererMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Referer() == "" {
			//http.Error(w, "404 page not found", 404)
			w.WriteHeader(404)
			renderer.RenderView(w, "404.html")
			return
		}
		next.ServeHTTP(w, r)
	})
}
