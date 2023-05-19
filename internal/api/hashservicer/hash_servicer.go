// Package hashservicer
package hashservicer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/model"
)

type HashServices interface {
	ValidHashServer(clientMetric model.Metric) bool
	GenerateHash(metric model.Metric) string
	IsEnabled() bool
}

type HashServer struct {
	key string
}

func NewHashServer(key string) *HashServer {
	return &HashServer{key: key}
}

// Checks if the hash is available
func (hs HashServer) IsEnabled() bool {
	return hs.key != ""
}

// Compares the hash received from the client with the hash stored on the server
func (hs HashServer) ValidHashServer(clientMetric model.Metric) bool {
	if hs.IsEnabled() {
		return clientMetric.Hash == hs.GenerateHash(clientMetric)
	}
	return true
}

func (hs HashServer) GenerateHash(metric model.Metric) string {
	var data string
	switch metric.MType {
	case "counter":
		data = fmt.Sprintf("%s:%s:%d", metric.ID, metric.MType, *metric.Delta)
	case "gauge":
		data = fmt.Sprintf("%s:%s:%f", metric.ID, metric.MType, *metric.Value)
	}
	h := hmac.New(sha256.New, []byte(hs.key))
	_, err := h.Write([]byte(data))
	if err != nil {
		log.Error(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}
