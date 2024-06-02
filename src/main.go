package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	myMiddleware "htmx-engineg/src/common/middleware"
	regiserty "htmx-engineg/src/controllers"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// Register routes
	for _, route := range regiserty.ControllerRegistery {
		r.Route(route.Route, func(r chi.Router) {
			// TODO: auth
			if route.NeedsReferer {
				r.Use(myMiddleware.RefererMiddleware)
			}
			for _, path := range route.Children {
				path(r)
			}
		})
	}

	// Serve static files
	r.Handle("/*", myMiddleware.FileServer())

	fmt.Println("Server is now running...")
	http.ListenAndServe(":8080", r)
}
