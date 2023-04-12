// Package handlers - handler instance that all handlers use
package handlers

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	database "github.com/vasiliyantufev/go-advanced-devops/internal/db"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/filestorage"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"
)

type Handler struct {
	memStorage  *memstorage.MemStorage
	fileStorage *filestorage.FileStorage
	config      *configserver.ConfigServer
	database    *database.DB
	hashServer  *hashservicer.HashServer
}

// NewHandler - creates a new server instance
func NewHandler(mem *memstorage.MemStorage,
	file *filestorage.FileStorage,
	cfg *configserver.ConfigServer,
	db *database.DB,
	hash *hashservicer.HashServer) *Handler {
	return &Handler{memStorage: mem, fileStorage: file, config: cfg, database: db, hashServer: hash}
}

// RestoreMetricsFromFile - restores metrics from a file
func (s Handler) RestoreMetricsFromFile() {
	if s.config.Restore {
		log.Info("Restore metrics")
		s.fileStorage.FileRestore(s.memStorage, s.config)
	}
}

// StoreMetricsToFile - saves metrics to a file
func (s Handler) StoreMetricsToFile() {
	if s.config.StoreFile != "" && s.config.DSN == "" {
		ticker := time.NewTicker(s.config.StoreInterval)
		for range ticker.C {
			log.Info("Store metrics")
			s.fileStorage.FileStore(s.memStorage, s.config)
		}
	}
}
