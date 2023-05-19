// storage - repositories
package memstorage

import (
	"sync"

	"github.com/vasiliyantufev/go-advanced-devops/internal/model"
)

type MemStorages interface {
	PutMetricsGauge(id string, o float64, h string)
	GetMetricsGauge(id string) (o float64, h string, b bool)
	PutMetricsCount(id string, o int64, h string)
	GetMetricsCount(id string) (o int64, h string, b bool)
	GetAllMetrics() []model.Metric
}

type MemStorage struct {
	mx   *sync.RWMutex
	data map[string]model.Metric
}

// NewMemStorage - creates a new store instance
func NewMemStorage() *MemStorage {
	return &MemStorage{
		mx:   new(sync.RWMutex),
		data: make(map[string]model.Metric),
	}
}

// PutMetricsGauge - puts the metric in storage with type Gauge
func (data *MemStorage) PutMetricsGauge(id string, o float64, h string) {
	data.mx.Lock()
	defer data.mx.Unlock()
	data.data[id] = model.Metric{
		ID:    id,
		MType: "gauge",
		Delta: nil,
		Value: &o,
		Hash:  h,
	}
}

// GetMetricsGauge - get metrics from storage of type Gauge
func (data *MemStorage) GetMetricsGauge(id string) (o float64, h string, b bool) {
	data.mx.RLock()
	defer data.mx.RUnlock()
	if metric, exists := data.data[id]; exists {
		return *metric.Value, metric.Hash, exists
	} else {
		return 0, "", exists
	}
}

// PutMetricsCount - puts the metric in storage with type Count
func (data *MemStorage) PutMetricsCount(id string, o int64, h string) {
	data.mx.Lock()
	defer data.mx.Unlock()
	data.data[id] = model.Metric{
		ID:    id,
		MType: "counter",
		Delta: &o,
		Value: nil,
		Hash:  h,
	}
}

// GetMetricsCount - gets metrics from storage of type Count
func (data *MemStorage) GetMetricsCount(id string) (o int64, h string, b bool) {
	data.mx.RLock()
	defer data.mx.RUnlock()
	if metric, exists := data.data[id]; exists {
		return *metric.Delta, metric.Hash, exists
	} else {
		return 0, "", exists
	}
}

// GetAllMetrics - gets all metrics from storage
func (data *MemStorage) GetAllMetrics() []model.Metric {
	data.mx.RLock()
	defer data.mx.RUnlock()
	result := make([]model.Metric, 0)
	for _, metric := range data.data {
		result = append(result, metric)
	}
	return result
}
