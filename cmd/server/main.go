package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	database "github.com/vasiliyantufev/go-advanced-devops/internal/db"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"os/signal"
	"syscall"
)

func main() {

	config.SetConfigServer()

	if err := database.ConnectDB(); err == nil {
		database.CreateTablesMigration()
	}

	log.SetLevel(config.GetConfigDebugLevelServer())

	mem := storage.NewMemStorage()
	srv := app.NewServer(mem)

	srv.RestoreMetricsFromFile()

	r := chi.NewRouter()
	//r.Use(middleware.Logger)
	r.Use(app.GzipHandle)

	r.Mount("/", srv.Route())

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	go app.StartServer(r)
	if config.GetConfigStoreIntervalServer() > 0 {
		go srv.StoreMetricsToFile()
	}
	<-ctx.Done()
	app.FileStore(srv.Mem)
	log.Info("server shutdown on signal with:", ctx.Err())
}
