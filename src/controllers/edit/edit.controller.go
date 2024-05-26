package controller

import (
	controller "htmx-engineg/src/models/controller"
	"net/http"
)

var Controller = controller.Controller{
	Route: "/edit",
	Children: []controller.Path{

		controller.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("TODO"))
		}),
	},
}
