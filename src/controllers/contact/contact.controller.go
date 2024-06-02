package controller

import (
	"encoding/json"
	controller "htmx-engineg/src/common/controller-struct"
	renderer "htmx-engineg/src/common/renderer"
	contact "htmx-engineg/src/models/contact"
	"net/http"
	"net/url"
)

var Controller = controller.Controller{
	Route:        "/contact",
	NeedsReferer: true,
	Children:     paths,
}

var paths = []controller.Path{

	controller.Get("/", func(w http.ResponseWriter, r *http.Request) {
		renderer.RenderView(w, "contact/form.html", getContact(r))
	}),

	controller.Get("/edit", func(w http.ResponseWriter, r *http.Request) {
		renderer.RenderView(w, "contact/editForm.html", getContact(r))
	}),

	controller.Put("/", func(w http.ResponseWriter, r *http.Request) {
		// Read form data
		r.ParseForm()
		newContact := contact.Contact{
			FirstName: r.FormValue("firstName"),
			LastName:  r.FormValue("lastName"),
			Email:     r.FormValue("email"),
		}

		// Save data to cookie
		data, _ := json.Marshal(newContact)
		http.SetCookie(w, &http.Cookie{Name: "contact", Value: url.QueryEscape(string(data))})

		// Render page
		renderer.RenderView(w, "contact/form.html", newContact)
	}),
}

// For demo purposes
func getContact(r *http.Request) contact.Contact {
	// Cookie test
	if cookie, err := r.Cookie("contact"); err == nil {
		saved := contact.Contact{}
		decoded, _ := url.QueryUnescape(cookie.Value)
		json.Unmarshal([]byte(decoded), &saved)
		return saved
	}
	return contact.Contact{
		FirstName: "Joe", LastName: "Blow", Email: "joe@blow.com",
	}
}
