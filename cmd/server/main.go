package main

import (
	"context"
	_ "net/http/pprof" // подключаем пакет pprof
	"os/signal"
	"syscall"

	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/server"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/server/handlers"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/server/routers"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	database "github.com/vasiliyantufev/go-advanced-devops/internal/db"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/filestorage"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	log.Infof("Build version: %s", buildVersion)
	log.Infof("Build date: %s", buildDate)
	log.Infof("Build commit: %s", buildCommit)

	configServer := configserver.NewConfigServer()
	log.SetLevel(configServer.DebugLevel)

	db, err := database.NewDB(configServer)
	if err != nil {
		log.Error(err)
	} else {
		defer db.Close()
		db.CreateTablesMigration(configServer)
	}

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer(configServer.Key)

	fileStorage, err := filestorage.NewMetricReadWriter(configServer)
	if err != nil {
		log.Error(err)
	}
	defer fileStorage.Close()

	srv := handlers.NewHandler(memStorage, fileStorage, configServer, db, hashServer)

	fileStorage.RestoreMetricsFromFile(memStorage)

	routerService := routers.Route(srv)
	rs := chi.NewRouter()
	rs.Mount("/", routerService)

	routerPProfile := routers.RoutePProf(srv)
	rp := chi.NewRouter()
	rp.Mount("/", routerPProfile)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	go server.StartService(rs, configServer)
	go server.StartPProfile(rp, configServer)

	if configServer.StoreInterval > 0 {
		go fileStorage.StoreMetricsToFile(memStorage)
	}
	<-ctx.Done()
	fileStorage.FileStore(memStorage)
	log.Info("server shutdown on signal with:", ctx.Err())
}
