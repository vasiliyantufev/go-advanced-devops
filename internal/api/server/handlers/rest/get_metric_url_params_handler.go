package rest

import (
	"fmt"
	"net/http"
	"strconv"

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
		param = strconv.FormatFloat(val, 'f', -1, 64)
	}
	if typeMetrics == "counter" {
		val, _, exists := s.memStorage.GetMetricsCount(nameMetrics)
		if !exists {
			log.Error("The name " + nameMetrics + " incorrect")
			http.Error(w, "The name "+nameMetrics+" incorrect", http.StatusNotFound)
			return
		}
		param = strconv.FormatInt(val, 10)
	}

	log.Debug("Request completed successfully " + nameMetrics + "=" + fmt.Sprint(param))
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(param))
	if err != nil {
		log.Error(err)
	}
}
