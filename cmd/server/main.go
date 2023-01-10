package main

import (
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"net/http"
)

const portNumber = ":8080"

func main() {

	r := chi.NewRouter()

	r.Get("/", app.IndexHandler)

	r.Route("/value", func(r chi.Router) {
		r.Get("/{type}/{name}", app.GetMetricsHandler)
	})
	r.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", app.MetricsHandler)
	})

	log.Infof("Starting application on port %v\n", portNumber)
	if con := http.ListenAndServe(portNumber, r); con != nil {
		log.Fatal(con)
	}

}
