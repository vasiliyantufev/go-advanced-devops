package handlers

import (
	"time"

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

// NewHandler - creates a new server instance
func NewHandler(mem *storage.MemStorage, cfg *configserver.ConfigServer, db *database.DB, hash *hashservicer.HashServer) *Handler {
	return &Handler{mem: mem, config: cfg, database: db, hashServer: hash}
}

// RestoreMetricsFromFile - restores metrics from a file
func (s Handler) RestoreMetricsFromFile() {
	if s.config.GetConfigRestoreServer() {
		log.Info("Restore metrics")
		file.FileRestore(s.mem, s.config)
	}
}

// StoreMetricsToFile - saves metrics to a file
func (s Handler) StoreMetricsToFile() {
	if s.config.GetConfigStoreFileServer() != "" && s.config.GetConfigDBServer() == "" {
		ticker := time.NewTicker(s.config.GetConfigStoreIntervalServer())
		for range ticker.C {
			log.Info("Store metrics")
			file.FileStore(s.mem, s.config)
		}
	}
}

// GetMem - get metrics from memory
func (s Handler) GetMem() *storage.MemStorage {
	return s.mem
}

// GetConfig - get the application configuration
func (s Handler) GetConfig() *configserver.ConfigServer {
	return s.config
}
