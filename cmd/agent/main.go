package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
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

	//go metricsPolling()
	go reportPolling()

	initMap()

	r := chi.NewRouter()

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("chi"))
	})

	// http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>;

	r.Route("/update", func(r chi.Router) {
		for name, _ := range metricsGauge {
			r.Post("/gauge/"+name, MetricsGaugeHandler)
		}
		for name, _ := range metricsCounter {
			r.Post("/counter/"+name, MetricsCounterHandler)
		}
	})

	http.ListenAndServe(":8080", r)
}

func sentMetrics() {

	log.Print(metricsGauge)
	log.Print(metricsCounter)

}

func MetricsGaugeHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s := r.URL.String()
	key := s[strings.LastIndex(s, "/")+1:]

	f, err := strconv.ParseFloat(string(resp), 64)
	if err != nil {
		panic(err)
	}

	metricsGauge[key] = f

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>MetricsGaugeHandler</h1>"))
}

func MetricsCounterHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s := r.URL.String()
	key := s[strings.LastIndex(s, "/")+1:]

	f, err := strconv.ParseInt(string(resp), 10, 64)
	if err != nil {
		panic(err)
	}

	metricsCounter[key] = f

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>MetricsCounterHandler</h1>"))
}

//http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>;

func initMap() {

	metricsCounter["poll_count"] = 0
	metricsCounter["random_value"] = 0
	metricsGauge["alloc"] = 0
	metricsGauge["buck_hash_sys"] = 0

	metricsGauge["Frees"] = 0
	metricsGauge["GCCPUFraction"] = 0
	metricsGauge["GCSys"] = 0
	metricsGauge["HeapAlloc"] = 0
	metricsGauge["HeapIdle"] = 0
	metricsGauge["HeapInuse"] = 0
	metricsGauge["HeapObjects"] = 0
	metricsGauge["HeapReleased"] = 0
	metricsGauge["HeapSys"] = 0
	metricsGauge["LastGC"] = 0
	metricsGauge["Lookups"] = 0
	metricsGauge["MCacheInuse"] = 0
	metricsGauge["MCacheSys"] = 0
	metricsGauge["MSpanInuse"] = 0
	metricsGauge["MSpanSys"] = 0
	metricsGauge["Mallocs"] = 0
	metricsGauge["NextGC"] = 0
	metricsGauge["NumForcedGC"] = 0
	metricsGauge["NumGC"] = 0
	metricsGauge["OtherSys"] = 0
	metricsGauge["PauseTotalNs"] = 0
	metricsGauge["StackInuse"] = 0
	metricsGauge["StackSys"] = 0
	metricsGauge["Sys"] = 0
	metricsGauge["TotalAlloc"] = 0
}

// https://golang.org/pkg/runtime/#MemStats
func getMetrics() (memStats runtime.MemStats) {

	metricsCounter["poll_count"] = metricsCounter["poll_count"] + 1
	metricsCounter["random_value"] = 123

	runtime.ReadMemStats(&memStats)
	metricsGauge["alloc"] = float64(memStats.Alloc)
	metricsGauge["buck_hash_sys"] = float64(memStats.BuckHashSys)

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

	return
}
