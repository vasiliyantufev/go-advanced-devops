// Package server - server module
package server

import (
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	grpcDevops "github.com/vasiliyantufev/go-advanced-devops/internal/api/proto"
	grpcHandler "github.com/vasiliyantufev/go-advanced-devops/internal/api/server/handlers/grpc"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	"google.golang.org/grpc"
	//"google.golang.org/grpc"
)

// StartRestService - starts the rest devops server
func StartRestService(r *chi.Mux, config *configserver.ConfigServer) {
	if config.CryptoKey != "" && config.Certificate != "" {
		log.Infof("Starting tls application %v\n", config.Address)
		if con := http.ListenAndServeTLS(config.Address, config.Certificate, config.CryptoKey, r); con != nil {
			log.Fatal(con)
		}
	} else {
		log.Infof("Starting application %v\n", config.Address)
		if con := http.ListenAndServe(config.Address, r); con != nil {
			log.Fatal(con)
		}
	}
}

// StartPProfile - starts the pprof
func StartPProfile(r *chi.Mux, config *configserver.ConfigServer) {
	log.Infof("Starting pprofile %v\n", config.AddressPProfile)
	if con := http.ListenAndServe(config.AddressPProfile, r); con != nil {
		log.Fatal(con)
	}
}

// StartGRPCService - starts the gRPC devops server
func StartGRPCService(grpcHandler *grpcHandler.Handler, config *configserver.ConfigServer) {

	log.Infof("Starting gRPC application %v\n", config.GRPC)

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", config.GRPC)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcDevops.RegisterDevopsServer(grpcServer, grpcHandler)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer Serve: %v", err)
	}
}
