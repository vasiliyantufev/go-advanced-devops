package grpchandler

import (
	"context"

	grpcDevops "github.com/vasiliyantufev/go-advanced-devops/internal/api/proto"
)

// PingHandler - checks is alive grpc server or not.
func (h *Handler) PingHandler(ctx context.Context, req *grpcDevops.PingRequest) (*grpcDevops.PingResponse, error) {
	return &grpcDevops.PingResponse{Resp: "Ping success"}, nil
}
