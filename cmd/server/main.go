package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	database "github.com/vasiliyantufev/go-advanced-devops/internal/db"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {

	cfg := config.NewConfigServer()

	if err := database.ConnectDB(cfg); err == nil {
		database.CreateTablesMigration()
	}

	mem := storage.NewMemStorage()
	hashServer := &config.HashServer{}

	log.SetLevel(cfg.GetConfigDebugLevelServer())
	//log.SetLevel(config.GetConfigDebugLevelServer())

	//srv := app.NewServer(mem, cfg, hashServer)
	srv := app.NewServer(mem, cfg, hashServer)

	srv.RestoreMetricsFromFile()

	r := chi.NewRouter()
	//r.Use(middleware.Logger)
	r.Use(app.GzipHandle)

	r.Mount("/", srv.Route())

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	go app.StartServer(r, srv.GetConfig())
	if cfg.GetConfigStoreIntervalServer() > 0 {
		//if config.GetConfigStoreIntervalServer() > 0 {
		go srv.StoreMetricsToFile()
	}
	<-ctx.Done()
	app.FileStore(srv.GetMem(), srv.GetConfig())
	log.Info("server shutdown on signal with:", ctx.Err())
}
