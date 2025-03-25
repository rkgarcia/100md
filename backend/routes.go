package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	return r
}
