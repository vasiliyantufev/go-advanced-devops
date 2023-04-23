package hashservicer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vasiliyantufev/go-advanced-devops/internal/converter"
	"github.com/vasiliyantufev/go-advanced-devops/internal/models"
)

var key = "secretKey"

func TestIsEnabled(t *testing.T) {
	hashServer := NewHashServer(key)
	assert.True(t, hashServer.IsEnabled())
}

func TestValidHashServer(t *testing.T) {
	hashServer := NewHashServer(key)
	hashAgent := NewHashServer(key)
	clientMetric := models.Metric{
		ID:    "Alloc",
		MType: "gauge",
		Delta: nil,
		Value: converter.Uint64ToFloat64Pointer(1),
	}
	clientMetric.Hash = hashAgent.GenerateHash(models.Metric{ID: clientMetric.ID, MType: clientMetric.MType, Delta: clientMetric.Delta, Value: clientMetric.Value})
	assert.True(t, hashServer.ValidHashServer(clientMetric))
}

func TestGenerateHash(t *testing.T) {
	hashAgent := NewHashServer(key)
	clientMetric := models.Metric{
		ID:    "Alloc",
		MType: "gauge",
		Delta: nil,
		Value: converter.Uint64ToFloat64Pointer(1),
	}
	clientMetric.Hash = hashAgent.GenerateHash(models.Metric{ID: clientMetric.ID, MType: clientMetric.MType, Delta: clientMetric.Delta, Value: clientMetric.Value})
	assert.True(t, len(clientMetric.Hash) > 0)
}
