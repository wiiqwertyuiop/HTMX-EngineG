package my_middleware

import (
	"htmx-engineg/src/common/renderer"
	"net/http"
	"os"
	"path/filepath"
)

func FileServer() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Path
		if file == "/" {
			file = "/index.html"
		}
		file = filepath.Clean(file)
		if dat, err := os.ReadFile("public" + file); err == nil {
			file_type := filepath.Ext(file)
			if file_type == ".css" {
				w.Header().Add("Content-Type", "text/css")
			}
			w.Write(dat)
		} else {
			w.WriteHeader(404)
			renderer.RenderView(w, "404.html")
		}
	})
}
