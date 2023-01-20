package app

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/jsonmetrics"
	"io"
	"os"
)

var metricGauge []*jsonmetrics.JSONMetricsToServer
var metricCounter []*jsonmetrics.JSONMetricsToServer

type metric struct {
	file    *os.File
	encoder *json.Encoder
}

func NewMetric(fileName string) (*metric, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}
	return &metric{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func (m *metric) WriteMetric(event *jsonmetrics.JSONMetricsToServer) error {
	return m.encoder.Encode(&event)
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

	mWrite, err := NewMetric(config.StoreFile)
	if err != nil {
		log.Error(err)
	}

	for name, val := range agent.DataMetricsGauge {
		m := new(jsonmetrics.JSONMetricsToServer)
		m.ID = name
		m.MType = "gauge"
		m.Value = &val

		if err := mWrite.WriteMetric(m); err != nil {
			log.Error(err)
		}
	}

	for name, val := range agent.DataMetricsCount {
		m := new(jsonmetrics.JSONMetricsToServer)
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

	mRead, err := NewMetric(config.StoreFile)
	if err != nil {
		log.Error(err)
	}

	dec := json.NewDecoder(mRead.file)
	for {
		mr := new(jsonmetrics.JSONMetricsToServer)
		if err := dec.Decode(mr); err == io.EOF {
			break
		} else if err != nil {
			log.Error(err)
		}

		if (mr.MType == "counter") {
			agent.PutMetricsCount(mr.ID, *mr.Delta)
		} else if mr.MType == "gauge" {
			agent.PutMetricsGauge(mr.ID, *mr.Value)
		}
	}

	mRead.Close()
}