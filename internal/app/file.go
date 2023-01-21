package app

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"io"
	"os"
)

var metricGauge []*storage.JSONMetrics
var metricCounter []*storage.JSONMetrics

type metric struct {
	file    *os.File
	encoder *json.Encoder
	decoder *json.Decoder
}

func NewMetricW(fileName string) (*metric, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}
	return &metric{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func NewMetricR(fileName string) (*metric, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return &metric{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}

func (m *metric) WriteMetric(event *storage.JSONMetrics) error {
	return m.encoder.Encode(&event)
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

func FileCreate(config storage.Config) {

	file, err := os.Create(config.StoreFile) // создаем файл
	if err != nil {                          // если возникла ошибка
		fmt.Println("Unable to create file:", err)
		os.Exit(1) // выходим из программы
	}
	file.Close()             // закрываем файл
	fmt.Println(file.Name()) // hello.txt

}

func FileStore(config storage.Config, agent *storage.MemStorage) {

	mWrite, err := NewMetricW(config.StoreFile)
	if err != nil {
		log.Error(err)
	}

	for name, val := range agent.DataMetricsGauge {
		m := new(storage.JSONMetrics)
		m.ID = name
		m.MType = "gauge"
		m.Value = &val

		if err := mWrite.WriteMetric(m); err != nil {
			log.Error(err)
		}
	}

	for name, val := range agent.DataMetricsCount {
		m := new(storage.JSONMetrics)
		m.ID = name
		m.MType = "counter"
		m.Delta = &val

		if err := mWrite.WriteMetric(m); err != nil {
			log.Error(err)
		}
	}
	mWrite.Close()
}

func FileRestore(config storage.Config, agent *storage.MemStorage) {

	mRead, err := NewMetricR(config.StoreFile)
	if err != nil {
		log.Error(err)
	}

	for {
		mr, err := mRead.ReadMetric()

		if err == io.EOF {
			log.Info("File end")
			break
		} else if err != nil {
			log.Error(err)
			break
		}

		if mr.MType == "counter" {
			agent.PutMetricsCount(mr.ID, *mr.Delta)
		} else if mr.MType == "gauge" {
			agent.PutMetricsGauge(mr.ID, *mr.Value)
		}
	}
	mRead.Close()
}
