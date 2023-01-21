package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"os/signal"
	"syscall"
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

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	go app.StartServer(cfg, r)
	if cfg.StoreInterval > 0 {
		go app.StoreMetricsToFile(cfg)
	}
	<-ctx.Done()
	app.FileStore(cfg, app.MemServer)
	log.Println("server shutdown on signal with:", ctx.Err())
}
