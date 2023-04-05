// Package server - server module
package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
)

// StartServer - starts the server
func StartServer(r *chi.Mux, config *configserver.ConfigServer) {
	log.Infof("Starting application %v\n", config.GetConfigAddressServer())
	if con := http.ListenAndServe(config.GetConfigAddressServer(), r); con != nil {
		log.Fatal(con)
	}
}
