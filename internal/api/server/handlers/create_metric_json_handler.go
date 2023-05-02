package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/converter"
	"github.com/vasiliyantufev/go-advanced-devops/internal/models"
)

// CreateMetricJSONHandler - create metric using json
func (s Handler) CreateMetricJSONHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value := models.Metric{}
	if err = json.Unmarshal([]byte(string(resp)), &value); err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rawValue := models.Metric{
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
		s.memStorage.PutMetricsGauge(value.ID, *value.Value, hashServer)
		rawValue.MType = value.MType
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
		if oldVal, _, exists := s.memStorage.GetMetricsCount(value.ID); exists {
			sum = oldVal + *value.Delta
		} else {
			sum = *value.Delta
		}
		// calculate new hash
		hashSumServer := s.hashServer.GenerateHash(models.Metric{ID: value.ID, MType: value.MType, Delta: converter.Int64ToInt64Pointer(sum), Value: value.Value})
		// store new metric
		s.memStorage.PutMetricsCount(value.ID, sum, hashSumServer)

		rawValue.MType = value.MType
		rawValue.Delta = &sum
		rawValue.Hash = hashSumServer

	}

	resp, err = json.Marshal(rawValue)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if s.database != nil {
		err = s.database.InsertOrUpdateMetrics(s.memStorage)
		if err != nil {
			log.Error(err)
		}
	}
	if s.config.StoreInterval == 0 {
		s.fileStorage.FileStore(s.memStorage)
	}

	log.Debug("Request completed successfully metric:" + value.ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		log.Error(err)
	}
}
