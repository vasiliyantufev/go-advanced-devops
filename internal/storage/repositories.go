package storage

import (
	"sync"
)

type MemStorage struct {
	mx   *sync.RWMutex
	data map[string]JSONMetrics
}

// Создаём map
func NewMemStorage() *MemStorage {
	return &MemStorage{
		mx:   new(sync.RWMutex),
		data: make(map[string]JSONMetrics),
	}
}
func (data *MemStorage) PutMetricsGauge(id string, o float64) {
	data.mx.Lock()
	defer data.mx.Unlock()
	data.data[id] = JSONMetrics{
		ID:    id,
		MType: "gauge",
		Delta: nil,
		Value: &o,
	}
}

func (data *MemStorage) GetMetricsGauge(id string) (o float64, b bool) {
	data.mx.RLock()
	defer data.mx.RUnlock()
	if metric, exists := data.data[id]; exists {
		return *metric.Value, exists
	} else {
		return 0, exists
	}
}

func (data *MemStorage) PutMetricsCount(id string, o int64) {
	data.mx.Lock()
	defer data.mx.Unlock()
	data.data[id] = JSONMetrics{
		ID:    id,
		MType: "counter",
		Delta: &o,
		Value: nil,
	}
}

func (data *MemStorage) GetMetricsCount(id string) (o int64, b bool) {
	data.mx.RLock()
	defer data.mx.RUnlock()
	if metric, exists := data.data[id]; exists {
		return *metric.Delta, exists
	} else {
		return 0, exists
	}
}

func (data *MemStorage) GetAllMetrics() []JSONMetrics {
	data.mx.RLock()
	defer data.mx.RUnlock()
	result := make([]JSONMetrics, 0)
	for _, metric := range data.data {
		result = append(result, metric)
	}
	return result
}
