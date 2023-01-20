package main

import (
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"sync"
	_ "time"
)

func main() {

	cfg := storage.GetConfig()

	log.SetLevel(log.DebugLevel)
	log.Info(cfg)

	app.RestoreMetricsFromFile(cfg)

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

	app.FileCreate(cfg)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go app.StartServer(cfg, r)
	go app.StoreMetricsToFile(cfg)
	wg.Wait()
}
