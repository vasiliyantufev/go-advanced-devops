// Package filestorage
package filestorage

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	"github.com/vasiliyantufev/go-advanced-devops/internal/models"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"

	log "github.com/sirupsen/logrus"
)

type FileStorages interface {
	WriteMetric(event *models.Metric) error
	ReadMetric() (*models.Metric, error)
	Close() error
	FileStore(mem *memstorage.MemStorage)
	FileRestore(mem *memstorage.MemStorage)
	RestoreMetricsFromFile(mem *memstorage.MemStorage)
	StoreMetricsToFile(mem *memstorage.MemStorage)
}

type FileStorage struct {
	file    *os.File
	config  *configserver.ConfigServer
	encoder *json.Encoder
	decoder *json.Decoder
}

// Creates a new file instance
func NewMetricReadWriter(config *configserver.ConfigServer) (*FileStorage, error) {
	file, err := os.OpenFile(config.StoreFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &FileStorage{
		file:    file,
		config:  config,
		encoder: json.NewEncoder(file),
		decoder: json.NewDecoder(file),
	}, nil
}

func (file *FileStorage) WriteMetric(event *models.Metric) error {
	return file.encoder.Encode(event)
}

func (file *FileStorage) ReadMetric() (*models.Metric, error) {
	mr := new(models.Metric)
	if err := file.decoder.Decode(mr); err == io.EOF {
		return nil, err
	} else if err != nil {
		log.Error(err)
	}
	return mr, nil
}

func (file *FileStorage) Close() error {
	return file.file.Close()
}

// Saves metrics from memory to a file
func (file *FileStorage) FileStore(mem *memstorage.MemStorage) {
	metrics := mem.GetAllMetrics()
	if len(metrics) == 0 {
		return
	}
	if err := file.file.Truncate(0); err != nil {
		log.Errorln("can't truncate file, cause:", err)
	}
	if _, err := file.file.Seek(0, 0); err != nil {
		log.Error("failed to seek:", err)
	}
	for _, val := range mem.GetAllMetrics() {
		if err := file.WriteMetric(&val); err != nil {
			log.Error("write to file failed with", err)
		}
	}
}

// Restores metrics from file to storage
func (file *FileStorage) FileRestore(mem *memstorage.MemStorage) {

	for {
		mr, err := file.ReadMetric()

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

// RestoreMetricsFromFile - restores metrics from a file
func (file *FileStorage) RestoreMetricsFromFile(mem *memstorage.MemStorage) {
	if file.config.Restore {
		log.Info("Restore metrics")
		file.FileRestore(mem)
	}
}

// StoreMetricsToFile - saves metrics to a file
func (file *FileStorage) StoreMetricsToFile(mem *memstorage.MemStorage) {
	if file.config.StoreFile != "" && file.config.DSN == "" {
		ticker := time.NewTicker(file.config.StoreInterval)
		for range ticker.C {
			log.Info("Store metrics")
			file.FileStore(mem)
		}
	}
}
