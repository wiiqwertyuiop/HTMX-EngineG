package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RenderPageView(w io.Writer, name string, data ...interface{}) error {
	t, _ := template.ParseFiles("src/views/page/" + name)
	if len(data) == 0 {
		return t.Execute(w, nil)
	}
	return t.Execute(w, data[0])
}

type Route struct {
	route   string
	handler http.HandlerFunc
}

var routes = []Route{
	{
		route: "/hello",
		handler: func(w http.ResponseWriter, r *http.Request) {
			RenderPageView(w, "hello/hello.html", "world")
		}},
	{
		route: "/hello2",
		handler: func(w http.ResponseWriter, r *http.Request) {
			RenderPageView(w, "hello/hello.html", "world")
		}}}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Register routes
	for _, route := range routes {
		r.Get(route.route, route.handler)
	}

	fs := http.FileServer(http.Dir("public"))
	r.Handle("/*", fs)

	fmt.Println("Server is now running...")
	http.ListenAndServe("localhost:8080", r)
}
