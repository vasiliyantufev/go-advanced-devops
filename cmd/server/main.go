package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	//storage.InitMap()

	r := chi.NewRouter()
	//app.RunRouter()

	r.Get("/", app.IndexHandler)
	r.Get("/value/{type}/{name}", app.GetMetricsHandler)

	r.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", app.MetricsHandler)
	})
	log.Printf("Starting application on port %v\n", portNumber)
	con := http.ListenAndServe(portNumber, r)
	if con != nil {
		log.Fatal(con)
	}

}
