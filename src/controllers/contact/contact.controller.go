package controller

import (
	C "htmx-engineg/src/common/controller-router"
	Renderer "htmx-engineg/src/common/renderer"
	Contact "htmx-engineg/src/models/contact"
	"net/http"
)

var Controller = C.Controller{
	Route:        "/contact",
	NeedsReferer: true,
	Children: []C.Path{

		C.Get("/", func(w http.ResponseWriter, r *http.Request) {
			contact := Contact.GetContact(r)
			Renderer.RenderView(w, "contact/form.html", contact)
		}),

		C.Get("/edit", func(w http.ResponseWriter, r *http.Request) {
			contact := Contact.GetContact(r)
			Renderer.RenderView(w, "contact/editForm.html", contact)
		}),

		C.Put("/", func(w http.ResponseWriter, r *http.Request) {
			newContact := Contact.UpdateContact(w, r)
			Renderer.RenderView(w, "contact/form.html", newContact)
		}),
	},
}
