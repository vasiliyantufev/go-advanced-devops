package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // в группе две горутины
	go PutMetrics()
	go SentMetrics()
	wg.Wait() // ожидаем завершения обоих горутин
}

func SentMetrics() {

	for range time.Tick(10 * time.Second) {

		log.Print("SentMetrics")

		// Create a Resty Client
		client := resty.New()

		for name, val := range storage.MetricsGauge {
			str := strconv.FormatFloat(val, 'f', 5, 64)
			resp, err := client.R().
				SetHeader("Content-Type", "text/plain").
				//SetBody(storage.MetricsGauge["alloc"]).
				Post("http://127.0.0.1:8080/update/gauge/" + name + "/" + str)

			if err != nil {
				log.Fatal(err)
			}
			log.Print(resp)
		}

		for name, val := range storage.MetricsCounter {
			str := strconv.FormatInt(val, 10)
			resp, err := client.R().
				SetHeader("Content-Type", "text/plain").
				//SetBody(storage.MetricsCounter["alloc"]).
				Post("http://127.0.0.1:8080/update/counter/" + name + "/" + str)

			if err != nil {
				log.Fatal(err)
			}
			log.Print(resp)
		}

	}
}

func PutMetrics() {

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	for range time.Tick(2 * time.Second) {

		storage.MetricsGauge["alloc"] = float64(memStats.Alloc)
		storage.MetricsGauge["buck_hash_sys"] = float64(memStats.BuckHashSys)
		storage.MetricsGauge["frees"] = float64(memStats.Frees)
		storage.MetricsGauge["gc_cpu_fraction"] = float64(memStats.GCCPUFraction)
		storage.MetricsGauge["gc_sys"] = float64(memStats.GCSys)
		storage.MetricsGauge["heap_alloc"] = float64(memStats.HeapAlloc)
		storage.MetricsGauge["heap_idle"] = float64(memStats.HeapIdle)
		storage.MetricsGauge["heap_inuse"] = float64(memStats.HeapInuse)
		storage.MetricsGauge["heap_objects"] = float64(memStats.HeapObjects)
		storage.MetricsGauge["heap_released"] = float64(memStats.HeapReleased)
		storage.MetricsGauge["heap_sys"] = float64(memStats.HeapSys)
		storage.MetricsGauge["last_gc"] = float64(memStats.LastGC)
		storage.MetricsGauge["lookups"] = float64(memStats.Lookups)
		storage.MetricsGauge["mcache_inuse"] = float64(memStats.MCacheInuse)
		storage.MetricsGauge["mcache_sys"] = float64(memStats.MCacheSys)
		storage.MetricsGauge["mspan_inuse"] = float64(memStats.MSpanInuse)
		storage.MetricsGauge["mspan_sys"] = float64(memStats.MSpanSys)
		storage.MetricsGauge["mallocs"] = float64(memStats.Mallocs)
		storage.MetricsGauge["next_gc"] = float64(memStats.NextGC)
		storage.MetricsGauge["num_forced_gc"] = float64(memStats.NumForcedGC)
		storage.MetricsGauge["num_gc"] = float64(memStats.NumGC)
		storage.MetricsGauge["other_sys"] = float64(memStats.OtherSys)
		storage.MetricsGauge["pause_total_ns"] = float64(memStats.PauseTotalNs)
		storage.MetricsGauge["stack_inuse"] = float64(memStats.StackInuse)
		storage.MetricsGauge["stack_sys"] = float64(memStats.StackSys)
		storage.MetricsGauge["sys"] = float64(memStats.Sys)
		storage.MetricsGauge["total_alloc"] = float64(memStats.TotalAlloc)

		storage.MetricsCounter["poll_count"] = storage.MetricsCounter["poll_count"] + 1
		storage.MetricsCounter["random_value"] = rand.Int63()

		//log.Print(storage.MetricsGauge)
		//log.Print(storage.MetricsCounter)
		log.Print("PutMetrics")
	}
}
