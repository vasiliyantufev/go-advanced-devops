// Package filestorage
package filestorage

import (
	"encoding/json"
	"io"
	"os"

	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	"github.com/vasiliyantufev/go-advanced-devops/internal/models"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"

	log "github.com/sirupsen/logrus"
)

type FileStorage struct {
	file    *os.File
	encoder *json.Encoder
	decoder *json.Decoder
}

// Creates a new file instance
func NewMetricReadWriter(fileName string) (*FileStorage, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &FileStorage{
		file:    file,
		encoder: json.NewEncoder(file),
		decoder: json.NewDecoder(file),
	}, nil
}

func (m *FileStorage) WriteMetric(event *models.Metric) error {
	return m.encoder.Encode(event)
}

func (m *FileStorage) ReadMetric() (*models.Metric, error) {

	mr := new(models.Metric)
	if err := m.decoder.Decode(mr); err == io.EOF {
		return nil, err
	} else if err != nil {
		log.Error(err)
	}
	return mr, nil
}

func (m *FileStorage) Close() error {
	return m.file.Close()
}

// Saves metrics from memory to a file
func (m *FileStorage) FileStore(mem *memstorage.MemStorage, config *configserver.ConfigServer) {

	//mWrite, err := NewMetricReadWriter(config.StoreFile)
	//if err != nil {
	//	log.Error(err)
	//}
	//defer mWrite.Close()

	metrics := mem.GetAllMetrics()
	if len(metrics) == 0 {
		return
	}
	if err := m.file.Truncate(0); err != nil {
		log.Errorln("can't truncate file, cause:", err)
	}
	if _, err := m.file.Seek(0, 0); err != nil {
		log.Error("failed to seek:", err)
	}
	for _, val := range mem.GetAllMetrics() {
		if err := m.WriteMetric(&val); err != nil {
			log.Error("write to file failed with", err)
		}
	}
}

// Restores metrics from file to storage
func (m *FileStorage) FileRestore(mem *memstorage.MemStorage, config *configserver.ConfigServer) {

	//mRead, err := NewMetricReadWriter(config.StoreFile)
	//if err != nil {
	//	log.Error(err)
	//}
	//defer mRead.Close()

	for {
		mr, err := m.ReadMetric()

		if err == io.EOF {
			log.Info("File end")
			return
		}
		if err != nil {
			log.Error(err)
			return
		}
		if mr.MType == "counter" {
			mem.PutMetricsCount(mr.ID, *mr.Delta, mr.Hash)
		}
		if mr.MType == "gauge" {
			mem.PutMetricsGauge(mr.ID, *mr.Value, mr.Hash)
		}
	}
}
