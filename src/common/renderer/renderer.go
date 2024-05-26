package renderer

import (
	"html/template"
	"io"
)

// Renders a view template from "/views"
func RenderPageView(w io.Writer, name string, data ...interface{}) error {
	t, _ := template.ParseFiles("src/views/" + name)
	if len(data) == 0 {
		return t.Execute(w, nil)
	}
	return t.Execute(w, data[0])
}
