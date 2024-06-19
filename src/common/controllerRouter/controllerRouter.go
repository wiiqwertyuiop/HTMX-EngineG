package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Path func(r chi.Router)

type Controller struct {
	Route        string
	NeedsReferer bool
	Children     []Path
}

// Setup GET path for route
func Get(path string, handler func(w http.ResponseWriter, r *http.Request)) Path {
	return func(r chi.Router) {
		r.Get(path, handler)
	}
}

// Setup POST path for route
func Post(path string, handler func(w http.ResponseWriter, r *http.Request)) Path {
	return func(r chi.Router) {
		r.Post(path, handler)
	}
}

// Setup PUT path for route
func Put(path string, handler func(w http.ResponseWriter, r *http.Request)) Path {
	return func(r chi.Router) {
		r.Put(path, handler)
	}
}

// Setup DELETE path for route
func Delete(path string, handler func(w http.ResponseWriter, r *http.Request)) Path {
	return func(r chi.Router) {
		r.Delete(path, handler)
	}
}
