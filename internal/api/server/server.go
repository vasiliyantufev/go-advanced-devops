// Package server - server module
package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
)

// StartService - starts the devops server
func StartService(r *chi.Mux, config *configserver.ConfigServer) {
	log.Infof("Starting application %v\n", config.Address)
	if con := http.ListenAndServe(config.Address, r); con != nil {
		log.Fatal(con)
	}
}

// StartPProfile - starts the pprof
func StartPProfile(r *chi.Mux, config *configserver.ConfigServer) {
	log.Infof("Starting pprofile %v\n", config.AddressPProfile)
	if con := http.ListenAndServe(config.AddressPProfile, r); con != nil {
		log.Fatal(con)
	}
}
