package rest

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/model"
)

// GetValueMetricJSONHandler - getting metric value using json
func (s Handler) GetValueMetricJSONHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("invalid request body: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value := model.Metric{}
	if err = json.Unmarshal([]byte(string(resp)), &value); err != nil {
		log.Errorf("invalid request body: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rawValue := model.Metric{
		ID:    value.ID,
		MType: value.MType,
	}
	log.Infoln("VALUE METRIC RECV", rawValue)
	if value.MType == "gauge" {
		val, hash, exists := s.memStorage.GetMetricsGauge(value.ID)
		if !exists {
			log.Error("Element " + value.ID + " not exists")
			http.Error(w, "Element "+value.ID+" not exists", http.StatusNotFound)
			return
		}
		rawValue.Value = &val
		rawValue.Hash = hash
	}
	if value.MType == "counter" {
		val, hash, exists := s.memStorage.GetMetricsCount(value.ID)
		if !exists {
			log.Error("Element " + value.ID + " not exists")
			http.Error(w, "Element "+value.ID+" not exists", http.StatusNotFound)
			return
		}
		rawValue.Delta = &val
		rawValue.Hash = hash
	}
	log.Infoln("VALUE METRIC RESPONSE", rawValue)
	resp, err = json.Marshal(rawValue)
	if err != nil {
		log.Errorf("invalid respounse body: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debug("Request completed successfully metric:" + value.ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		log.Error(err)
	}
}
