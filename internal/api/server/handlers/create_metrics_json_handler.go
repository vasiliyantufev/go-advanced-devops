package handlers

import (
	"encoding/json"
	"io"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/converter"
	"github.com/vasiliyantufev/go-advanced-devops/internal/models"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/errors"
)

// CreateMetricsJSONHandler - create metrics using json
func (s Handler) CreateMetricsJSONHandler(w http.ResponseWriter, r *http.Request) {

	if s.config.TrustedSubnet != "" {
		// look at the X-Real-IP request header
		IPAddressAgent := net.ParseIP(r.Header.Get("X-Real-IP"))
		_, subnet, _ := net.ParseCIDR(s.config.TrustedSubnet)

		if !subnet.Contains(IPAddressAgent) {
			log.Error(errors.ErrAddressNotTrustedSubnet)
			http.Error(w, errors.ErrAddressNotTrustedSubnet.Error(), http.StatusForbidden)
			w.WriteHeader(http.StatusForbidden)
		}
	}

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var metrics []*models.Metric
	if err = json.Unmarshal([]byte(string(resp)), &metrics); err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, metric := range metrics {
		if metric.Value != nil {
			if !s.hashServer.ValidHashServer(*metric) {
				log.Error(errors.ErrHashSum)
				http.Error(w, errors.ErrHashSum.Error(), http.StatusBadRequest)
				return
			}
			s.memStorage.PutMetricsGauge(metric.ID, *metric.Value, s.hashServer.GenerateHash(*metric))
		}
		if metric.Delta != nil {
			// compare hashes
			if !s.hashServer.ValidHashServer(*metric) {
				log.Error(errors.ErrHashSum)
				http.Error(w, errors.ErrHashSum.Error(), http.StatusBadRequest)
				return
			}
			// counter summing logic
			var sum int64
			if oldVal, _, exists := s.memStorage.GetMetricsCount(metric.ID); exists {
				sum = oldVal + *metric.Delta
			} else {
				sum = *metric.Delta
			}
			// calculate new hash
			hashSumServer := s.hashServer.GenerateHash(models.Metric{ID: metric.ID, MType: metric.MType, Delta: converter.Int64ToInt64Pointer(sum), Value: metric.Value})
			// store new metric
			s.memStorage.PutMetricsCount(metric.ID, sum, hashSumServer)
		}
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

	resp, err = json.Marshal(s.memStorage.GetAllMetrics())
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debug("Request completed successfully metrics")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		log.Error(err)
	}
}
