package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"math/rand"
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

	wg.Add(2) // в группе две горутины
	go PutMetrics(cfg)
	go SentMetrics(cfg)
	wg.Wait() // ожидаем завершения обоих горутин
}

func SentMetrics(config storage.Config) {

	// Create a Resty Client
	client := resty.New()

	for range time.Tick(config.ReportInterval) {

		log.Info("Sent metrics")

		for name, val := range MemAgent.DataMetricsGauge {

			//str := strconv.FormatFloat(val, 'f', 5, 64)
			//_, err := client.R().
			//	SetHeader("Content-Type", "text/plain").
			//	Post("http://localhost:8080/update/gauge/" + name + "/" + str)
			//if err != nil {
			//	log.Error(err)
			//}

			_, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(storage.Metrics{ID: name, MType: "gauge", Value: &val}).
				Post("http://localhost" + config.Port + "/update/")
			if err != nil {
				log.Error(err)
			} /**/

		}

		//=============================================================================

		for name, val := range MemAgent.DataMetricsCount {

			//str := strconv.FormatInt(val, 10)
			//_, err := client.R().
			//	SetHeader("Content-Type", "text/plain").
			//	Post("http://localhost:8080/update/counter/" + name + "/" + str)
			//if err != nil {
			//	log.Error(err)
			//}

			_, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(storage.Metrics{ID: name, MType: "counter", Delta: &val}).
				Post("http://localhost" + config.Port + "/update/")
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

		log.Info("Put metrics")
	}
}
