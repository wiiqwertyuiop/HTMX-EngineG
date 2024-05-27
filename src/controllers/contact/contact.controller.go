package controller

import (
	"encoding/json"
	controller "htmx-engineg/src/common/controller"
	renderer "htmx-engineg/src/common/renderer"
	contact "htmx-engineg/src/models/contact"
	"net/http"
	"net/url"
)

// Example
var exampleContact = contact.Contact{
	FirstName: "Joe", LastName: "Blow", Email: "joe@blow.com",
}

var Controller = controller.Controller{
	Route:        "/contact",
	NeedsReferer: true,
	Children: []controller.Path{

		controller.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// Cookie test
			// Read saved data from cookie
			if cookie, err := r.Cookie("contact"); err == nil {
				decoded, _ := url.QueryUnescape(cookie.Value)
				json.Unmarshal([]byte(decoded), &exampleContact)
			}
			renderer.RenderView(w, "contact/form.html", exampleContact)
		}),

		controller.Get("/edit", func(w http.ResponseWriter, r *http.Request) {
			renderer.RenderView(w, "contact/editForm.html", exampleContact)
		}),

		controller.Put("/", func(w http.ResponseWriter, r *http.Request) {
			// Read form data
			r.ParseForm()
			exampleContact = contact.Contact{
				FirstName: r.FormValue("firstName"),
				LastName:  r.FormValue("lastName"),
				Email:     r.FormValue("email"),
			}

			// Save data to cookie
			data, _ := json.Marshal(exampleContact)
			http.SetCookie(w, &http.Cookie{Name: "contact", Value: url.QueryEscape(string(data))})

			// Render page
			renderer.RenderView(w, "contact/form.html", exampleContact)
		}),
	},
}
