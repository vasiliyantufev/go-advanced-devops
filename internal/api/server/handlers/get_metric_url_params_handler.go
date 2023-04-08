package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
)

// GetMetricURLParamsHandler - getting metric using url parameters
func (s Handler) GetMetricURLParamsHandler(w http.ResponseWriter, r *http.Request) {

	typeMetrics := chi.URLParam(r, "type")
	nameMetrics := chi.URLParam(r, "name")

	status, err := storage.ValidURLParamGetMetrics(typeMetrics, nameMetrics)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), status)
		return
	}

	var param string
	if typeMetrics == "gauge" {
		val, _, exists := s.mem.GetMetricsGauge(nameMetrics)
		if !exists {
			log.Error("The name " + nameMetrics + " incorrect")
			http.Error(w, "The name "+nameMetrics+" incorrect", http.StatusNotFound)
			return
		}
		param = fmt.Sprintf("%.3f", val)
	}
	if typeMetrics == "counter" {
		val, _, exists := s.mem.GetMetricsCount(nameMetrics)
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
