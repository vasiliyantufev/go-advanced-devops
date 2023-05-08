// Package handlers - handler instance that all handlers use
package rest

import (
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	database "github.com/vasiliyantufev/go-advanced-devops/internal/db"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/filestorage"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"
)

type DBS interface {
	Ping() error
	InsertOrUpdateMetrics(metrics *memstorage.MemStorage) error
}

type Handler struct {
	memStorage  *memstorage.MemStorage
	fileStorage *filestorage.FileStorage
	config      *configserver.ConfigServer
	database    *database.DB
	//database   DBS
	hashServer *hashservicer.HashServer
}

// NewHandler - creates a new server instance
func NewHandler(mem *memstorage.MemStorage,
	file *filestorage.FileStorage,
	cfg *configserver.ConfigServer,
	db *database.DB,
	hash *hashservicer.HashServer) *Handler {
	return &Handler{memStorage: mem, fileStorage: file, config: cfg, database: db, hashServer: hash}
}
