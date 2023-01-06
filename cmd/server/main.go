package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"net/http"
)

func main() {

	storage.InitMap()

	r := chi.NewRouter()
	r.Get("/", app.IndexHandler)
	r.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", app.MetricsHandler)
	})

	http.ListenAndServe(":8080", r)
}
