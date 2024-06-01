package renderer

import (
	"html/template"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
)

var templateMap = make(map[string]*template.Template)

func init() {
	// Parse all HTML template files in the templates directory
	basepath := "src/views/"
	err := filepath.Walk(basepath, func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			key := strings.TrimPrefix(strings.ReplaceAll(path, "\\", "/"), basepath)
			templateMap[key] = template.Must(template.ParseFiles(path))
		}
		return err
	})
	if err != nil {
		panic(err)
	}
}

// Renders a view template from "/views"
func RenderView(w io.Writer, name string, data ...interface{}) error {
	if len(data) == 0 {
		return templateMap[name].Execute(w, nil)
	}
	return templateMap[name].Execute(w, data[0])
}
