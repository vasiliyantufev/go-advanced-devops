package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

var metricsGauge = make(map[string]float64)
var metricsCounter = make(map[string]int64)

func metricsPolling() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("Update Metrics")
		go getMetrics()
	}
}

func reportPolling() {

	for {
		<-time.After(10 * time.Second)
		fmt.Println("Set Metrics")
		go sentMetrics()

	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	metricsCounter["PollCount"] = 0

	go metricsPolling()
	go reportPolling()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

// https://golang.org/pkg/runtime/#MemStats
func getMetrics() (memStats runtime.MemStats) {

	runtime.ReadMemStats(&memStats)
	metricsGauge["Alloc"] = float64(memStats.Alloc)
	metricsGauge["BuckHashSys"] = float64(memStats.BuckHashSys)
	metricsGauge["Frees"] = float64(memStats.Frees)
	metricsGauge["GCCPUFraction"] = float64(memStats.GCCPUFraction)
	metricsGauge["GCSys"] = float64(memStats.GCSys)
	metricsGauge["HeapAlloc"] = float64(memStats.HeapAlloc)
	metricsGauge["HeapIdle"] = float64(memStats.HeapIdle)
	metricsGauge["HeapInuse"] = float64(memStats.HeapInuse)
	metricsGauge["HeapObjects"] = float64(memStats.HeapObjects)
	metricsGauge["HeapReleased"] = float64(memStats.HeapReleased)
	metricsGauge["HeapSys"] = float64(memStats.HeapSys)
	metricsGauge["LastGC"] = float64(memStats.LastGC)
	metricsGauge["Lookups"] = float64(memStats.Lookups)
	metricsGauge["MCacheInuse"] = float64(memStats.MCacheInuse)
	metricsGauge["MCacheSys"] = float64(memStats.MCacheSys)
	metricsGauge["MSpanInuse"] = float64(memStats.MSpanInuse)
	metricsGauge["MSpanSys"] = float64(memStats.MSpanSys)
	metricsGauge["Mallocs"] = float64(memStats.Mallocs)
	metricsGauge["NextGC"] = float64(memStats.NextGC)
	metricsGauge["NumForcedGC"] = float64(memStats.NumForcedGC)
	metricsGauge["NumGC"] = float64(memStats.NumGC)
	metricsGauge["OtherSys"] = float64(memStats.OtherSys)
	metricsGauge["PauseTotalNs"] = float64(memStats.PauseTotalNs)
	metricsGauge["StackInuse"] = float64(memStats.StackInuse)
	metricsGauge["StackSys"] = float64(memStats.StackSys)
	metricsGauge["Sys"] = float64(memStats.Sys)
	metricsGauge["TotalAlloc"] = float64(memStats.TotalAlloc)

	metricsCounter["PollCount"] = metricsCounter["PollCount"] + 1
	metricsCounter["RandomValue"] = metricsCounter["PollCount"] + 1

	return
}

func sentMetrics() {

	log.Print(metricsGauge)

}
