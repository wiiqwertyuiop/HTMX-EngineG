package controller

import (
	c "htmx-engineg/src/common/controller-router"
	"htmx-engineg/src/common/renderer"
	"net/http"
)

var Controller = c.Controller{
	Route: "/hello",
	Children: []c.Path{

		c.Get("/", func(w http.ResponseWriter, r *http.Request) {
			renderer.RenderView(w, "hello/hello.html", "World")
		}),

		c.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello x2!"))
		}),
	},
}
