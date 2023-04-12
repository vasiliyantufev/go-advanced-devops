package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

// GetMetricURLParamsHandler - getting metric using url parameters
func (s Handler) GetMetricURLParamsHandler(w http.ResponseWriter, r *http.Request) {

	typeMetrics := chi.URLParam(r, "type")
	nameMetrics := chi.URLParam(r, "name")

	var param string
	if typeMetrics == "gauge" {
		val, _, exists := s.memStorage.GetMetricsGauge(nameMetrics)
		if !exists {
			log.Error("The name " + nameMetrics + " incorrect")
			http.Error(w, "The name "+nameMetrics+" incorrect", http.StatusNotFound)
			return
		}
		param = fmt.Sprintf("%.3f", val)
	}
	if typeMetrics == "counter" {
		val, _, exists := s.memStorage.GetMetricsCount(nameMetrics)
		if !exists {
			log.Error("The name " + nameMetrics + " incorrect")
			http.Error(w, "The name "+nameMetrics+" incorrect", http.StatusNotFound)
			return
		}
		param = fmt.Sprint(val)
	}

	log.Debug("Request completed successfully " + nameMetrics + "=" + fmt.Sprint(param))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(param))
}
