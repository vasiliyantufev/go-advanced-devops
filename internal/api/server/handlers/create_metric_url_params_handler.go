package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/converter"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
)

// CreateMetricURLParamsHandler - create metric using url parameters
func (s Handler) CreateMetricURLParamsHandler(w http.ResponseWriter, r *http.Request) {

	typeMetrics := chi.URLParam(r, "type")
	nameMetrics := chi.URLParam(r, "name")
	valueMetrics := chi.URLParam(r, "value")

	status, err := storage.ValidURLParamMetrics(typeMetrics, nameMetrics, valueMetrics)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), status)
		return
	}

	var resp string
	if typeMetrics == "gauge" {
		val, err := strconv.ParseFloat(string(valueMetrics), 64)
		if err != nil {
			log.Error("The query parameter value " + valueMetrics + " is incorrect")
			http.Error(w, "The query parameter value "+valueMetrics+" is incorrect", http.StatusBadRequest)
			return
		}
		hashServer := s.hashServer.GenerateHash(storage.JSONMetrics{ID: nameMetrics, MType: "gauge", Delta: nil, Value: converter.Float64ToFloat64Pointer(val)})
		s.mem.PutMetricsGauge(nameMetrics, val, hashServer)
		resp = "Request completed successfully " + nameMetrics + "=" + fmt.Sprint(val)
	}
	if typeMetrics == "counter" {
		val, err := strconv.ParseInt(string(valueMetrics), 10, 64)
		if err != nil {
			log.Error("The query parameter value " + valueMetrics + " is incorrect")
			http.Error(w, "The query parameter value "+valueMetrics+" is incorrect", http.StatusBadRequest)
			return
		}
		var sum int64
		if oldVal, _, exists := s.mem.GetMetricsCount(nameMetrics); exists {
			sum = oldVal + val
		} else {
			sum = val
		}
		hashServer := s.hashServer.GenerateHash(storage.JSONMetrics{ID: nameMetrics, MType: "counter", Delta: converter.Int64ToInt64Pointer(val), Value: nil})
		s.mem.PutMetricsCount(nameMetrics, sum, hashServer)
		resp = "Request completed successfully " + nameMetrics + "=" + fmt.Sprint(sum)
	}

	log.Debug(resp)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
