package storage

import (
	"sync"
)

type MemStorage struct {
	Mx               *sync.RWMutex
	DataMetricsGauge map[string]float64
	DataMetricsCount map[string]int64
}

// Создаём map
func NewMemStorage() *MemStorage {
	return &MemStorage{
		Mx:               new(sync.RWMutex),
		DataMetricsGauge: make(map[string]float64),
		DataMetricsCount: make(map[string]int64),
	}
}

func (data *MemStorage) PutMetricsGauge(id string, o float64) {
	data.Mx.Lock()
	defer data.Mx.Unlock()
	data.DataMetricsGauge[id] = o
}

func (data *MemStorage) GetMetricsGauge(id string) (o float64, b bool) {
	data.Mx.RLock()
	defer data.Mx.RUnlock()
	o, b = data.DataMetricsGauge[id]
	return
}

func (data *MemStorage) PutMetricsCount(id string, o int64) {
	data.Mx.Lock()
	defer data.Mx.Unlock()
	data.DataMetricsCount[id] = o
}

func (data *MemStorage) GetMetricsCount(id string) (o int64, b bool) {
	data.Mx.RLock()
	defer data.Mx.RUnlock()
	o, b = data.DataMetricsCount[id]
	return
}
