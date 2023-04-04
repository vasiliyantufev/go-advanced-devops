package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/file"
	"github.com/vasiliyantufev/go-advanced-devops/internal/converter"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
)

// Create metrics using json
func (s Handler) CreateMetricsJSONHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var metrics []*storage.JSONMetrics
	if err := json.Unmarshal([]byte(string(resp)), &metrics); err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, metric := range metrics {

		if metric.Value != nil {

			if !s.hashServer.ValidHashServer(*metric) {
				log.Error("Хеш-сумма не соответствует расчетной")
				http.Error(w, "Хеш-сумма не соответствует расчетной", http.StatusBadRequest)
				return
			}
			s.mem.PutMetricsGauge(metric.ID, *metric.Value, s.hashServer.GenerateHash(*metric))

		}
		if metric.Delta != nil {

			// compare hashes
			if !s.hashServer.ValidHashServer(*metric) {
				log.Error("Хеш-сумма не соответствует расчетной")
				http.Error(w, "Хеш-сумма не соответствует расчетной", http.StatusBadRequest)
				return
			}
			// counter summing logic
			var sum int64
			if oldVal, _, exists := s.mem.GetMetricsCount(metric.ID); exists {
				sum = oldVal + *metric.Delta
			} else {
				sum = *metric.Delta
			}
			// calculate new hash
			hashSumServer := s.hashServer.GenerateHash(storage.JSONMetrics{ID: metric.ID, MType: metric.MType, Delta: converter.Int64ToInt64Pointer(sum), Value: metric.Value})
			// store new metric
			s.mem.PutMetricsCount(metric.ID, sum, hashSumServer)
		}
	}

	if s.database != nil {
		s.database.InsertOrUpdateMetrics(s.mem)
	}
	if s.config.GetConfigStoreIntervalServer() == 0 {
		file.FileStore(s.mem, nil)
	}

	log.Debug("Request completed successfully metric")
	w.WriteHeader(http.StatusOK)
}
