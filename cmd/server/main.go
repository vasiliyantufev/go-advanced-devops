// module main
package main

import (
	"context"
	_ "net/http/pprof" // подключаем пакет pprof
	"os/signal"
	"syscall"

	"github.com/vasiliyantufev/go-advanced-devops/internal/api/file"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/server"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/server/handlers"
	routerdevops "github.com/vasiliyantufev/go-advanced-devops/internal/api/server/router"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	database "github.com/vasiliyantufev/go-advanced-devops/internal/db"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

// main server
func main() {
	configServer := configserver.NewConfigServer()

	db, err := database.NewDB(configServer)
	if err != nil {
		log.Error(err)
	} else {
		defer db.Close()
		db.CreateTablesMigration(configServer)
	}

	mem := storage.NewMemStorage()

	hashServer := hashservicer.NewHashServer(configServer.GetConfigKeyServer())

	log.SetLevel(configServer.GetConfigDebugLevelServer())

	srv := handlers.NewHandler(mem, configServer, db, hashServer)
	router := routerdevops.Route(srv)

	srv.RestoreMetricsFromFile()

	r := chi.NewRouter()
	r.Mount("/", router)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	go server.StartServer(r, srv.GetConfig())
	if configServer.GetConfigStoreIntervalServer() > 0 {
		go srv.StoreMetricsToFile()
	}
	<-ctx.Done()
	file.FileStore(srv.GetMem(), srv.GetConfig())
	log.Info("server shutdown on signal with:", ctx.Err())
}
