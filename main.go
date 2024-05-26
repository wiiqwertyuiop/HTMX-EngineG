package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	registery "htmx-engineg/endpoints"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// Register routes
	for _, route := range registery.ControllerRegistery {
		r.Route(route.Route, func(r chi.Router) {
			for _, path := range route.Children {
				path(r)
			}
		})
	}

	// Serve static files
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/*", fs)

	fmt.Println("Server is now running...")
	http.ListenAndServe("localhost:8080", r)
}