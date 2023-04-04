package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Checking if the database is available
func (s Handler) PingHandler(w http.ResponseWriter, r *http.Request) {

	if s.database != nil {
		if err := s.database.Ping(); err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Info("ping")
		w.WriteHeader(http.StatusOK)
		return
	}
	log.Error("db is nil")
	w.WriteHeader(http.StatusInternalServerError)
}
