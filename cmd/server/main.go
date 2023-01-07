package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"net/http"
)

func main() {

	//storage.InitMap()

	r := chi.NewRouter()

	r.Get("/", app.IndexHandler)
	r.Get("/value/{type}/{name}", app.GetMetricsHandler)

	r.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", app.MetricsHandler)
	})

	http.ListenAndServe(":8080", r)
}
