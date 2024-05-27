package controller

import (
	controller "htmx-engineg/src/common/controller"
	"htmx-engineg/src/common/renderer"
	"net/http"
)

var Controller = controller.Controller{
	Route: "/hello",
	Children: []controller.Path{

		controller.Get("/", func(w http.ResponseWriter, r *http.Request) {
			renderer.RenderView(w, "hello/hello.html", "World")
		}),

		controller.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello x2!"))
		}),
	},
}
