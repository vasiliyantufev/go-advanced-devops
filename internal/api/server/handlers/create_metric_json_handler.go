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

// CreateMetricJSONHandler - create metric using json
func (s Handler) CreateMetricJSONHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value := storage.JSONMetrics{}
	if err := json.Unmarshal([]byte(string(resp)), &value); err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rawValue := storage.JSONMetrics{
		ID: value.ID,
	}

	if value.Value != nil {
		// expect true id hashServer is disabled
		if isValid := s.hashServer.ValidHashServer(value); isValid {
			log.Println("GAUGE METRIC", isValid)
		}

		if !s.hashServer.ValidHashServer(value) {
			log.Error("Хеш-сумма не соответствует расчетной")
			http.Error(w, "Хеш-сумма не соответствует расчетной", http.StatusBadRequest)
			return
		}

		hashServer := s.hashServer.GenerateHash(value)
		s.mem.PutMetricsGauge(value.ID, *value.Value, hashServer)
		rawValue.Value = value.Value
		rawValue.Hash = hashServer
	}
	if value.Delta != nil {
		// expect true id hashServer is disabled
		if isValid := s.hashServer.ValidHashServer(value); isValid {
			log.Println("GAUGE METRIC", isValid)
		}
		// compare hashes
		if !s.hashServer.ValidHashServer(value) {
			log.Error("Хеш-сумма не соответствует расчетной")
			http.Error(w, "Хеш-сумма не соответствует расчетной", http.StatusBadRequest)
			return
		}

		// counter summing logic
		var sum int64
		if oldVal, _, exists := s.mem.GetMetricsCount(value.ID); exists {
			sum = oldVal + *value.Delta
		} else {
			sum = *value.Delta
		}
		// calculate new hash
		hashSumServer := s.hashServer.GenerateHash(storage.JSONMetrics{ID: value.ID, MType: value.MType, Delta: converter.Int64ToInt64Pointer(sum), Value: value.Value})
		// store new metric
		s.mem.PutMetricsCount(value.ID, sum, hashSumServer)

		rawValue.Delta = &sum
		rawValue.Hash = hashSumServer

	}

	resp, err = json.Marshal(rawValue)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if s.database != nil {
		s.database.InsertOrUpdateMetrics(s.mem)
	}
	if s.config.StoreInterval == 0 {
		file.FileStore(s.mem, nil)
	}

	log.Debug("Request completed successfully metric:" + value.ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
