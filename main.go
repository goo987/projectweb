package main

import (
	"net/http"

	"projectweb/views"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	PORT := ":8080"

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		views.Index().Render(r.Context(), w)
	})

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		views.Login().Render(r.Context(), w)
	})

	r.Get("/Register", func(w http.ResponseWriter, r *http.Request) {
		views.Register().Render(r.Context(), w)
	})

	r.Get("/pengaduan", func(w http.ResponseWriter, r *http.Request) {
		views.Pengaduan().Render(r.Context(), w)
	})

	r.Get("/listpengaduan", func(w http.ResponseWriter, r *http.Request) {
		views.ListPengaduan().Render(r.Context(), w)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		views.Notfound().Render(r.Context(), w)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		views.Unautorize().Render(r.Context(), w)
	})
}
