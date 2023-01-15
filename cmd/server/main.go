package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"net/http"
)

func main() {

	var cfg storage.Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.SetLevel(log.DebugLevel)
	log.Info(cfg)

	r := chi.NewRouter()
	//r.Use(middleware.Logger)

	r.Get("/", app.IndexHandler)
	r.Route("/value", func(r chi.Router) {
		r.Get("/{type}/{name}", app.GetMetricsHandler)
		r.Post("/", app.PostValueMetricsHandler)
	})
	r.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", app.MetricsHandler)
		r.Post("/", app.PostMetricsHandler)
	})

	log.Infof("Starting application on port %v\n", cfg.Port)
	if con := http.ListenAndServe(cfg.Port, r); con != nil {
		log.Fatal(con)
	}
}
