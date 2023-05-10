// Package handlers - handler instance that all handlers use
package grpchandler

import (
	"context"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	grpcDevops "github.com/vasiliyantufev/go-advanced-devops/internal/api/proto"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	"github.com/vasiliyantufev/go-advanced-devops/internal/converter"
	database "github.com/vasiliyantufev/go-advanced-devops/internal/db"
	"github.com/vasiliyantufev/go-advanced-devops/internal/models"
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

func (h *Handler) CreateMetric(ctx context.Context, req *grpcDevops.CreateMetricRequest) (*grpcDevops.CreateMetricResponse, error) {
	var resp string
	if req.Type == "gauge" {
		val, err := strconv.ParseFloat(string(req.Value), 64)
		if err != nil {
			log.Fatal("The query parameter value " + req.Value + " is incorrect")
		}
		hashServer := h.hashServer.GenerateHash(models.Metric{ID: req.Name, MType: "gauge", Delta: nil, Value: converter.Float64ToFloat64Pointer(val)})
		h.memStorage.PutMetricsGauge(req.Name, val, hashServer)
		resp = "Request completed successfully " + req.Name + "=" + fmt.Sprint(val)
	}
	if req.Type == "counter" {
		val, err := strconv.ParseInt(string(req.Value), 10, 64)
		if err != nil {
			log.Fatal("The query parameter value " + req.Value + " is incorrect")
		}
		var sum int64
		if oldVal, _, exists := h.memStorage.GetMetricsCount(req.Name); exists {
			sum = oldVal + val
		} else {
			sum = val
		}
		hashServer := h.hashServer.GenerateHash(models.Metric{ID: req.Name, MType: "counter", Delta: converter.Int64ToInt64Pointer(val), Value: nil})
		h.memStorage.PutMetricsCount(req.Name, sum, hashServer)
		resp = "Request completed successfully " + req.Name + "=" + fmt.Sprint(sum)
	}
	return &grpcDevops.CreateMetricResponse{Resp: resp}, nil
}
