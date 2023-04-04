package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/file"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	database "github.com/vasiliyantufev/go-advanced-devops/internal/db"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
)

type Handler struct {
	mem        *storage.MemStorage
	config     *configserver.ConfigServer
	database   *database.DB
	hashServer *hashservicer.HashServer
}

// Creates a new server instance
func NewHandler(mem *storage.MemStorage, cfg *configserver.ConfigServer, db *database.DB, hash *hashservicer.HashServer) *Handler {
	return &Handler{mem: mem, config: cfg, database: db, hashServer: hash}
}

func StartServer(r *chi.Mux, config *configserver.ConfigServer) {

	log.Infof("Starting application %v\n", config.GetConfigAddressServer())
	if con := http.ListenAndServe(config.GetConfigAddressServer(), r); con != nil {
		log.Fatal(con)
	}
}

func (s Handler) RestoreMetricsFromFile() {

	if s.config.GetConfigRestoreServer() {
		log.Info("Restore metrics")
		file.FileRestore(s.mem, s.config)
	}
}

func (s Handler) StoreMetricsToFile() {

	if s.config.GetConfigStoreFileServer() != "" && s.config.GetConfigDBServer() == "" {
		ticker := time.NewTicker(s.config.GetConfigStoreIntervalServer())
		for range ticker.C {
			log.Info("Store metrics")
			file.FileStore(s.mem, s.config)
		}
	}
}

func (s Handler) GetMem() *storage.MemStorage {
	return s.mem
}

func (s Handler) GetConfig() *configserver.ConfigServer {
	return s.config
}
