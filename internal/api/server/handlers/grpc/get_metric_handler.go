package grpchandler

import (
	"context"
	"strconv"

	log "github.com/sirupsen/logrus"

	grpcDevops "github.com/vasiliyantufev/go-advanced-devops/internal/api/proto"
)

// GetMetricHandler - getting metric
func (h *Handler) GetMetricHandler(ctx context.Context, req *grpcDevops.GetMetricRequest) (*grpcDevops.GetMetricResponse, error) {
	var param string
	if req.Type == "gauge" {
		val, _, exists := h.memStorage.GetMetricsGauge(req.Name)
		if !exists {
			log.Fatal("The name " + req.Name + " incorrect")
		}
		param = strconv.FormatFloat(val, 'f', -1, 64)
	}
	if req.Type == "counter" {
		val, _, exists := h.memStorage.GetMetricsCount(req.Name)
		if !exists {
			log.Fatal("The name " + req.Name + " incorrect")
		}
		param = strconv.FormatInt(val, 10)
	}
	return &grpcDevops.GetMetricResponse{Val: param}, nil
}
