package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"runtime"
	"sync"
	"time"
)

var MemAgent = storage.NewMemStorage()

func main() {

	var cfg storage.Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go PutMetrics(cfg)
	go SentMetrics(cfg)
	wg.Wait()
}

func PutMetrics(config storage.Config) {

	var memStats runtime.MemStats
	for range time.Tick(config.PollInterval) {
		log.Info("Put metrics")
		runtime.ReadMemStats(&memStats)
		app.DataFromRuntime(MemAgent, memStats)
	}
}

func SentMetrics(config storage.Config) {

	// Create a Resty Client
	client := resty.New()
	for range time.Tick(config.ReportInterval) {
		log.Info("Sent metrics")
		for name, val := range MemAgent.DataMetricsGauge {
			_, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(storage.JSONMetrics{ID: name, MType: "gauge", Value: &val}).
				Post("http://" + config.Address + "/update/")
			if err != nil {
				log.Error(err)
			}
		}
		//=============================================================================
		for name, val := range MemAgent.DataMetricsCount {
			_, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(storage.JSONMetrics{ID: name, MType: "counter", Delta: &val}).
				Post("http://" + config.Address + "/update/")
			if err != nil {
				log.Error(err)
			}
		}
	}
}
