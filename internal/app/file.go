package app

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"io"
	"os"
)

type metric struct {
	file    *os.File
	encoder *json.Encoder
	decoder *json.Decoder
}

func NewMetricReadWriter(fileName string) (*metric, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &metric{
		file:    file,
		encoder: json.NewEncoder(file),
		decoder: json.NewDecoder(file),
	}, nil
}

func (m *metric) WriteMetric(event *storage.JSONMetrics) error {
	return m.encoder.Encode(event)
}

func (m *metric) ReadMetric() (*storage.JSONMetrics, error) {

	mr := new(storage.JSONMetrics)
	if err := m.decoder.Decode(mr); err == io.EOF {
		return nil, err
	} else if err != nil {
		log.Error(err)
	}
	return mr, nil
}

func (m *metric) Close() error {
	return m.file.Close()
}

func FileStore(config storage.Config, agent *storage.MemStorage) {

	mWrite, err := NewMetricReadWriter(config.StoreFile)
	if err != nil {
		log.Fatal(err)
	}
	defer mWrite.Close()

	metrics := agent.GetAllMetrics()
	if len(metrics) == 0 {
		return
	}
	if err := mWrite.file.Truncate(0); err != nil {
		log.Fatalln("can't truncate file, cause:", err)
	}
	if _, err := mWrite.file.Seek(0, 0); err != nil {
		log.Fatal("failed to seek:", err)
	}
	for _, val := range agent.GetAllMetrics() {
		if err := mWrite.WriteMetric(&val); err != nil {
			log.Error("write to file failed with", err)
		}
	}
}

func FileRestore(config storage.Config, agent *storage.MemStorage) {

	mRead, err := NewMetricReadWriter(config.StoreFile)
	if err != nil {
		log.Fatal(err)
	}
	defer mRead.Close()
	for {
		mr, err := mRead.ReadMetric()

		if err == io.EOF {
			log.Info("File end")
			return
		}
		if err != nil {
			log.Fatal(err)
			return
		}
		if mr.MType == "counter" {
			agent.PutMetricsCount(mr.ID, *mr.Delta)
		}
		if mr.MType == "gauge" {
			agent.PutMetricsGauge(mr.ID, *mr.Value)
		}
	}
}
