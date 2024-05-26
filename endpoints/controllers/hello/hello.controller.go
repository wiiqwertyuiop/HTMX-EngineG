package hello_controller

import (
	controller "htmx-engineg/endpoints/controllers"
	"net/http"
)

var HelloController = controller.Controller{
	Route: "/hello",
	Children: []controller.Path{

		controller.Get("/", func(w http.ResponseWriter, r *http.Request) {
			controller.RenderPageView(w, "hello/hello.html", "World")
		}),

		controller.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello x2!"))
		}),
	},
}
