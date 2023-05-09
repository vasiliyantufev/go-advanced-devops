// Package handlers - handler instance that all handlers use
package grpchandler

import (
	"context"

	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	grpcDevops "github.com/vasiliyantufev/go-advanced-devops/internal/api/proto"
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
	grpcDevops.UnimplementedDevopsServer
}

// NewHandler - creates a new grpc server instance
func NewHandler(
	mem *memstorage.MemStorage,
	file *filestorage.FileStorage,
	cfg *configserver.ConfigServer,
	db *database.DB,
	hash *hashservicer.HashServer) *Handler {
	return &Handler{memStorage: mem, fileStorage: file, config: cfg, database: db, hashServer: hash /**/}
}

// Ping checks is alive db or not.
func (h *Handler) Ping(ctx context.Context, req *grpcDevops.PingRequest) (*grpcDevops.PingResponse, error) {
	return &grpcDevops.PingResponse{}, nil
}
