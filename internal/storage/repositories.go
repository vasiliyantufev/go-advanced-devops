package storage

import "sync"

type MemStorage struct {
	mx               sync.RWMutex
	DataMetricsGauge map[string]float64
	DataMetricsCount map[string]int64
}

//Создаём map
func NewMemStorage() *MemStorage {
	return &MemStorage{
		DataMetricsGauge: make(map[string]float64),
		DataMetricsCount: make(map[string]int64),
	}
}

func (data *MemStorage) PutMetricsGauge(id string, o float64) {
	data.mx.Lock()
	defer data.mx.Unlock()
	data.DataMetricsGauge[id] = o
}

func (data *MemStorage) GetMetricsGauge(id string) (o float64, b bool) {
	data.mx.RLock()
	defer data.mx.RUnlock()
	o, b = data.DataMetricsGauge[id]
	return
}

func (data *MemStorage) PutMetricsCount(id string, o int64) {
	data.mx.Lock()
	defer data.mx.Unlock()
	data.DataMetricsCount[id] = o
}

func (data *MemStorage) GetMetricsCount(id string) (o int64, b bool) {
	data.mx.RLock()
	defer data.mx.RUnlock()
	o, b = data.DataMetricsCount[id]
	return
}
