// Module hash server
package app

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
)

type HashServices interface {
	ValidHashServer(metric storage.JSONMetrics, clientHash string) bool
	GenerateHash(metric storage.JSONMetrics) string
	isEnabled()
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
func (hs HashServer) ValidHashServer(clientMetric storage.JSONMetrics) bool {
	if hs.IsEnabled() {
		return clientMetric.Hash == hs.GenerateHash(clientMetric)
	}
	return true
}

func (hs HashServer) GenerateHash(metric storage.JSONMetrics) string {
	var data string
	switch metric.MType {
	case "counter":
		data = fmt.Sprintf("%s:%s:%d", metric.ID, metric.MType, *metric.Delta)
	case "gauge":
		data = fmt.Sprintf("%s:%s:%f", metric.ID, metric.MType, *metric.Value)
	}
	h := hmac.New(sha256.New, []byte(hs.key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
