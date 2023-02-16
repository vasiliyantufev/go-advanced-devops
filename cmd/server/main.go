package main

import (
	"context"
	database "github.com/vasiliyantufev/go-advanced-devops/internal/db"
	"os/signal"
	"syscall"

	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {

	configServer := config.NewConfigServer()

	//if err := database.ConnectDB(cfg); err == nil {
	//	database.CreateTablesMigration()
	//}
	db, err := database.NewDB(configServer)
	if err != nil {
		log.Error(err)
	} else {
		defer db.Close()
		db.CreateTablesMigration()
	}

	mem := storage.NewMemStorage()
	hashServer := &config.HashServer{}

	log.SetLevel(configServer.GetConfigDebugLevelServer())
	srv := app.NewServer(mem, configServer, db, hashServer)

	srv.RestoreMetricsFromFile()

	r := chi.NewRouter()
	//r.Use(middleware.Logger)
	r.Use(app.GzipHandle)

	r.Mount("/", srv.Route())

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	go app.StartServer(r, srv.GetConfig())
	if configServer.GetConfigStoreIntervalServer() > 0 {
		go srv.StoreMetricsToFile()
	}
	<-ctx.Done()
	app.FileStore(srv.GetMem(), srv.GetConfig())
	log.Info("server shutdown on signal with:", ctx.Err())
}
