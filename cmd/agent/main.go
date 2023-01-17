package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"math/rand"
	"os"
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

	RestoreMetrics(cfg)

	wg := new(sync.WaitGroup)
	wg.Add(3)
	go StoreMetrics(cfg)
	go PutMetrics(cfg)
	go SentMetrics(cfg)
	wg.Wait()
}

func RestoreMetrics(config storage.Config) {

	if config.Restore {
		log.Info("Restore metrics")

		file, err := os.OpenFile(config.StoreFile, os.O_RDONLY|os.O_CREATE, 0777)
		if err != nil {
			fmt.Print(err)
			return
		}
		reader := bufio.NewReader(file)
		data, err := reader.ReadBytes('\n')

		metric := storage.OuterMetrics{}
		err = json.Unmarshal(data, &metric)
		if err != nil {
			fmt.Print(err)
			return
		}
		file.Close()

		MemAgent.PutMetricsGauge("Alloc", float64(metric.DataMetricsGauge.Alloc))
		MemAgent.PutMetricsGauge("BuckHashSys", float64(metric.DataMetricsGauge.BuckHashSys))
		MemAgent.PutMetricsGauge("Frees", float64(metric.DataMetricsGauge.Frees))
		MemAgent.PutMetricsGauge("GCCPUFraction", float64(metric.DataMetricsGauge.GCCPUFraction))
		MemAgent.PutMetricsGauge("GCSys", float64(metric.DataMetricsGauge.GCSys))
		MemAgent.PutMetricsGauge("HeapAlloc", float64(metric.DataMetricsGauge.HeapAlloc))
		MemAgent.PutMetricsGauge("HeapIdle", float64(metric.DataMetricsGauge.HeapIdle))
		MemAgent.PutMetricsGauge("HeapInuse", float64(metric.DataMetricsGauge.HeapInuse))
		MemAgent.PutMetricsGauge("HeapObjects", float64(metric.DataMetricsGauge.HeapObjects))
		MemAgent.PutMetricsGauge("HeapReleased", float64(metric.DataMetricsGauge.HeapReleased))
		MemAgent.PutMetricsGauge("HeapSys", float64(metric.DataMetricsGauge.HeapSys))
		MemAgent.PutMetricsGauge("LastGC", float64(metric.DataMetricsGauge.LastGC))
		MemAgent.PutMetricsGauge("Lookups", float64(metric.DataMetricsGauge.Lookups))
		MemAgent.PutMetricsGauge("MCacheInuse", float64(metric.DataMetricsGauge.MCacheInuse))
		MemAgent.PutMetricsGauge("MCacheSys", float64(metric.DataMetricsGauge.MCacheSys))
		MemAgent.PutMetricsGauge("MSpanInuse", float64(metric.DataMetricsGauge.MSpanInuse))
		MemAgent.PutMetricsGauge("MSpanSys", float64(metric.DataMetricsGauge.MSpanSys))
		MemAgent.PutMetricsGauge("Mallocs", float64(metric.DataMetricsGauge.Mallocs))
		MemAgent.PutMetricsGauge("NextGC", float64(metric.DataMetricsGauge.NextGC))
		MemAgent.PutMetricsGauge("NumForcedGC", float64(metric.DataMetricsGauge.NumForcedGC))
		MemAgent.PutMetricsGauge("NumGC", float64(metric.DataMetricsGauge.NumGC))
		MemAgent.PutMetricsGauge("OtherSys", float64(metric.DataMetricsGauge.OtherSys))
		MemAgent.PutMetricsGauge("PauseTotalNs", float64(metric.DataMetricsGauge.PauseTotalNs))
		MemAgent.PutMetricsGauge("StackInuse", float64(metric.DataMetricsGauge.StackInuse))
		MemAgent.PutMetricsGauge("StackSys", float64(metric.DataMetricsGauge.StackSys))
		MemAgent.PutMetricsGauge("Sys", float64(metric.DataMetricsGauge.Sys))
		MemAgent.PutMetricsGauge("TotalAlloc", float64(metric.DataMetricsGauge.TotalAlloc))
		MemAgent.PutMetricsGauge("RandomValue", float64(metric.DataMetricsGauge.RandomValue))
		MemAgent.PutMetricsCount("PollCount", int64(metric.DataMetricsCount.PollCount))

		fmt.Print(MemAgent)
	}
}

func StoreMetrics(config storage.Config) {

	if config.StoreFile != "" {
		for range time.Tick(config.StoreInterval) {
			log.Info("Store metrics")

			file, err := os.OpenFile(config.StoreFile, os.O_WRONLY|os.O_CREATE, 0777)
			if err != nil {
				fmt.Print(err)
				return
			}
			writer := bufio.NewWriter(file)

			fileData, err := json.Marshal(MemAgent)
			if err != nil {
				fmt.Print(err)
				return
			}

			if _, err := writer.Write(fileData); err != nil {
				fmt.Print(err)
				return
			}
			writer.Flush()
			file.Close()
		}
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
				SetBody(storage.Metrics{ID: name, MType: "gauge", Value: &val}).
				Post("http://" + config.Address + "/update/")
			if err != nil {
				log.Error(err)
			}
		}
		//=============================================================================
		for name, val := range MemAgent.DataMetricsCount {
			_, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(storage.Metrics{ID: name, MType: "counter", Delta: &val}).
				Post("http://" + config.Address + "/update/")
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func PutMetrics(config storage.Config) {

	var memStats runtime.MemStats
	for range time.Tick(config.PollInterval) {
		runtime.ReadMemStats(&memStats)
		MemAgent.PutMetricsGauge("Alloc", float64(memStats.Alloc))
		MemAgent.PutMetricsGauge("BuckHashSys", float64(memStats.BuckHashSys))
		MemAgent.PutMetricsGauge("Frees", float64(memStats.Frees))
		MemAgent.PutMetricsGauge("GCCPUFraction", float64(memStats.GCCPUFraction))
		MemAgent.PutMetricsGauge("GCSys", float64(memStats.GCSys))
		MemAgent.PutMetricsGauge("HeapAlloc", float64(memStats.HeapAlloc))
		MemAgent.PutMetricsGauge("HeapIdle", float64(memStats.HeapIdle))
		MemAgent.PutMetricsGauge("HeapInuse", float64(memStats.HeapInuse))
		MemAgent.PutMetricsGauge("HeapObjects", float64(memStats.HeapObjects))
		MemAgent.PutMetricsGauge("HeapReleased", float64(memStats.HeapReleased))
		MemAgent.PutMetricsGauge("HeapSys", float64(memStats.HeapSys))
		MemAgent.PutMetricsGauge("LastGC", float64(memStats.LastGC))
		MemAgent.PutMetricsGauge("Lookups", float64(memStats.Lookups))
		MemAgent.PutMetricsGauge("MCacheInuse", float64(memStats.MCacheInuse))
		MemAgent.PutMetricsGauge("MCacheSys", float64(memStats.MCacheSys))
		MemAgent.PutMetricsGauge("MSpanInuse", float64(memStats.MSpanInuse))
		MemAgent.PutMetricsGauge("MSpanSys", float64(memStats.MSpanSys))
		MemAgent.PutMetricsGauge("Mallocs", float64(memStats.Mallocs))
		MemAgent.PutMetricsGauge("NextGC", float64(memStats.NextGC))
		MemAgent.PutMetricsGauge("NumForcedGC", float64(memStats.NumForcedGC))
		MemAgent.PutMetricsGauge("NumGC", float64(memStats.NumGC))
		MemAgent.PutMetricsGauge("OtherSys", float64(memStats.OtherSys))
		MemAgent.PutMetricsGauge("PauseTotalNs", float64(memStats.PauseTotalNs))
		MemAgent.PutMetricsGauge("StackInuse", float64(memStats.StackInuse))
		MemAgent.PutMetricsGauge("StackSys", float64(memStats.StackSys))
		MemAgent.PutMetricsGauge("Sys", float64(memStats.Sys))
		MemAgent.PutMetricsGauge("TotalAlloc", float64(memStats.TotalAlloc))
		MemAgent.PutMetricsGauge("RandomValue", rand.Float64())

		pollCount, _ := MemAgent.GetMetricsCount("PollCount")
		MemAgent.PutMetricsCount("PollCount", pollCount+1)

		storage.Store(*MemAgent)

		log.Info("Put metrics")
	}
}
