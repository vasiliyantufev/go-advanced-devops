package main

// @Title Metrics API
// @Description Metrics and alerting service
// @Version 1.0

// @Contact.email vasiliyantufev@gmail.com

// @Host http://127.0.0.1:8080/

import (
	"context"
	_ "net/http/pprof"
	"os/signal"
	"syscall"

	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/helpers"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/server"

	grpcHandler "github.com/vasiliyantufev/go-advanced-devops/internal/api/server/handlers/grpc"
	restHandler "github.com/vasiliyantufev/go-advanced-devops/internal/api/server/handlers/rest"

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
	helpers.PrintInfo(buildVersion, buildDate, buildCommit)

	configServer := configserver.NewConfigServer()

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
	fileStorage.RestoreMetricsFromFile(memStorage)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	srvRest := restHandler.NewHandler(memStorage, fileStorage, configServer, db, hashServer)
	routerService := routers.Route(ctx, srvRest)
	rs := chi.NewRouter()
	rs.Mount("/", routerService)

	routerPProfile := routers.RoutePProf(srvRest)
	rp := chi.NewRouter()
	rp.Mount("/", routerPProfile)

	if configServer.GRPC != "" {
		handlerGrpc := grpcHandler.NewHandler(memStorage, fileStorage, configServer, db, hashServer)
		go server.StartGRPCService(handlerGrpc, configServer)
	} else {
		go server.StartRestService(rs, configServer)
	}

	go server.StartPProfile(rp, configServer)

	if configServer.StoreInterval > 0 {
		go fileStorage.StoreMetricsToFile(memStorage)
	}
	<-ctx.Done()
	fileStorage.FileStore(memStorage)
	log.Info("server shutdown on signal with:", ctx.Err())
}
