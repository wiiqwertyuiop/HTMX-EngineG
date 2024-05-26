package controller

import (
	controller "htmx-engineg/src/common/controller"
	"net/http"
)

var Controller = controller.Controller{
	Route:        "/edit",
	NeedsReferer: true,
	Children: []controller.Path{
		controller.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("TODO"))
		}),
	},
}
