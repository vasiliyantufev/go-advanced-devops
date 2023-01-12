package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"math/rand"
	"runtime"
	"strconv"
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

	var wg sync.WaitGroup

	wg.Add(2) // в группе две горутины
	go PutMetrics(cfg.Poll_interval)
	go SentMetrics(cfg.Report_interval)
	wg.Wait() // ожидаем завершения обоих горутин
}

func SentMetrics(interval time.Duration) {

	log.Info(interval)

	// Create a Resty Client
	client := resty.New()

	for range time.Tick(interval) {

		log.Info("Sent metrics")

		for name, val := range MemAgent.DataMetricsGauge {
			str := strconv.FormatFloat(val, 'f', 5, 64)
			_, err := client.R().
				SetHeader("Content-Type", "text/plain").
				Post("http://127.0.0.1:8080/update/gauge/" + name + "/" + str)

			if err != nil {
				log.Error(err)
			}
		}

		for name, val := range MemAgent.DataMetricsCount {
			str := strconv.FormatInt(val, 10)
			_, err := client.R().
				SetHeader("Content-Type", "text/plain").
				Post("http://127.0.0.1:8080/update/counter/" + name + "/" + str)

			if err != nil {
				log.Error(err)
			}
		}

	}
}

func PutMetrics(interval time.Duration) {

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	for range time.Tick(interval) {

		MemAgent.PutMetricsGauge("alloc", float64(memStats.Alloc))
		MemAgent.PutMetricsGauge("buck_hash_sys", float64(memStats.BuckHashSys))
		MemAgent.PutMetricsGauge("frees", float64(memStats.Frees))
		MemAgent.PutMetricsGauge("gc_cpu_fraction", float64(memStats.GCCPUFraction))
		MemAgent.PutMetricsGauge("gc_sys", float64(memStats.GCSys))
		MemAgent.PutMetricsGauge("heap_alloc", float64(memStats.HeapAlloc))
		MemAgent.PutMetricsGauge("heap_idle", float64(memStats.HeapIdle))
		MemAgent.PutMetricsGauge("heap_inuse", float64(memStats.HeapInuse))
		MemAgent.PutMetricsGauge("heap_objects", float64(memStats.HeapObjects))
		MemAgent.PutMetricsGauge("heap_released", float64(memStats.HeapReleased))
		MemAgent.PutMetricsGauge("heap_sys", float64(memStats.HeapSys))
		MemAgent.PutMetricsGauge("last_gc", float64(memStats.LastGC))
		MemAgent.PutMetricsGauge("lookups", float64(memStats.Lookups))
		MemAgent.PutMetricsGauge("mcache_inuse", float64(memStats.MCacheInuse))
		MemAgent.PutMetricsGauge("mcache_sys", float64(memStats.MCacheSys))
		MemAgent.PutMetricsGauge("mspan_inuse", float64(memStats.MSpanInuse))
		MemAgent.PutMetricsGauge("mspan_sys", float64(memStats.MSpanSys))
		MemAgent.PutMetricsGauge("mallocs", float64(memStats.Mallocs))
		MemAgent.PutMetricsGauge("next_gc", float64(memStats.NextGC))
		MemAgent.PutMetricsGauge("num_forced_gc", float64(memStats.NumForcedGC))
		MemAgent.PutMetricsGauge("num_gc", float64(memStats.NumGC))
		MemAgent.PutMetricsGauge("other_sys", float64(memStats.OtherSys))
		MemAgent.PutMetricsGauge("pause_total_ns", float64(memStats.PauseTotalNs))
		MemAgent.PutMetricsGauge("stack_inuse", float64(memStats.StackInuse))
		MemAgent.PutMetricsGauge("stack_sys", float64(memStats.StackSys))
		MemAgent.PutMetricsGauge("sys", float64(memStats.Sys))
		MemAgent.PutMetricsGauge("total_alloc", float64(memStats.TotalAlloc))

		pollCount, _ := MemAgent.GetMetricsCount("poll_count")
		MemAgent.PutMetricsCount("poll_count", pollCount+1)
		MemAgent.PutMetricsCount("random_value", rand.Int63())

		log.Info("Put metrics")
	}
}
