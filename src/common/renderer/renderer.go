package renderer

import (
	"html/template"
	"io"
)

// Renders a view template from "/views"
func RenderView(w io.Writer, name string, data ...interface{}) error {
	if t, err := template.ParseFiles("src/views/" + name); err == nil {
		if len(data) == 0 {
			return t.Execute(w, nil)
		}
		return t.Execute(w, data[0])
	} else {
		// Potential infinite loop here, but can never realistically happen
		return RenderView(w, "404.html")
	}
}
