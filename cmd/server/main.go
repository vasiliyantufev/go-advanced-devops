package main

import (
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	database "github.com/vasiliyantufev/go-advanced-devops/internal/db"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/flags"
	"os/signal"
	"syscall"
	_ "time"

	"context"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	flags.SetFlagsServer()
	config.SetConfigServer()

	//database.ConnectDB()
	if err := database.ConnectDB(); err == nil {
		database.CreateTables()
	}

	log.SetLevel(config.GetConfigDebugLevelServer())

	app.RestoreMetricsFromFile()

	r := chi.NewRouter()
	//r.Use(middleware.Logger)
	r.Use(app.GzipHandle)
	r.Get("/", app.IndexHandler)
	r.Get("/ping", app.PingHandler)
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

	go app.StartServer(r)
	if config.GetConfigStoreIntervalServer() > 0 {
		go app.StoreMetricsToFile()
	}
	<-ctx.Done()
	app.FileStore(app.MemServer)
	log.Info("server shutdown on signal with:", ctx.Err())
}
