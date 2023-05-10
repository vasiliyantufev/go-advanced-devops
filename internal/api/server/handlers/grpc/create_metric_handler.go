package grpchandler

import (
	"context"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"

	grpcDevops "github.com/vasiliyantufev/go-advanced-devops/internal/api/proto"
	"github.com/vasiliyantufev/go-advanced-devops/internal/converter"
	"github.com/vasiliyantufev/go-advanced-devops/internal/models"
)

// CreateMetricHandler - create metric
func (h *Handler) CreateMetricHandler(ctx context.Context, req *grpcDevops.CreateMetricRequest) (*grpcDevops.CreateMetricResponse, error) {
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
