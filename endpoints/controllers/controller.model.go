package controllers

import (
	"html/template"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Path func(r chi.Router)

type Controller struct {
	Route    string
	Children []Path
}

func RenderPageView(w io.Writer, name string, data ...interface{}) error {
	t, _ := template.ParseFiles("endpoints/views/page/" + name)
	if len(data) == 0 {
		return t.Execute(w, nil)
	}
	return t.Execute(w, data[0])
}

func Get(path string, handler func(w http.ResponseWriter, r *http.Request)) Path {
	return func(r chi.Router) {
		r.Get(path, handler)
	}
}
