package app

import (
	_ "compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	"html/template"
	"io"
	"net/http"
	"strconv"
	_ "strings"
	"time"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
)

var MemServer = storage.NewMemStorage()

type ViewData struct {
	MapG map[string]float64
	MapC map[string]int64
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	//log.Error(MemServer.GetAllMetrics())

	tmpl, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		log.Errorf("Parse failed: %s", err)
		http.Error(w, "Error loading index page", http.StatusInternalServerError)
		return
	}

	gauges := make(map[string]float64)
	counters := make(map[string]int64)

	metrics := MemServer.GetAllMetrics()

	for _, metric := range metrics {
		if metric.MType == "gauge" {
			gauges[metric.ID] = *metric.Value
		}
		if metric.MType == "counter" {
			counters[metric.ID] = *metric.Delta
		}
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	data := ViewData{MapG: gauges, MapC: counters}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Errorf("Execution failed: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {

	typeMetrics := chi.URLParam(r, "type")
	if typeMetrics == "" {
		log.Error("The query parameter type is missing")
		http.Error(w, "The query parameter type is missing", http.StatusBadRequest)
		return
	}
	if typeMetrics != "gauge" && typeMetrics != "counter" {
		log.Error("The type incorrect " + typeMetrics)
		http.Error(w, "The type incorrect", http.StatusNotImplemented)
		return
	}
	nameMetrics := chi.URLParam(r, "name")
	if nameMetrics == "" {
		log.Error("The query parameter name is missing")
		http.Error(w, "The query parameter name is missing", http.StatusBadRequest)
		return
	}

	valueMetrics := chi.URLParam(r, "value")
	if valueMetrics == "" {
		log.Error("The query parameter value is missing")
		http.Error(w, "The query parameter value is missing", http.StatusBadRequest)
		return
	}

	var resp string
	if typeMetrics == "gauge" {
		val, err := strconv.ParseFloat(string(valueMetrics), 64)
		if err != nil {
			log.Error("The query parameter value " + valueMetrics + " is incorrect")
			http.Error(w, "The query parameter value "+valueMetrics+" is incorrect", http.StatusBadRequest)
			return
		}

		hashServer := config.GetHashServer(nameMetrics, "gauge", 0, val)
		MemServer.PutMetricsGauge(nameMetrics, val, hashServer)
		resp = "Request completed successfully " + nameMetrics + "=" + fmt.Sprint(val)
	}
	if typeMetrics == "counter" {
		val, err := strconv.ParseInt(string(valueMetrics), 10, 64)
		if err != nil {
			log.Error("The query parameter value " + valueMetrics + " is incorrect")
			http.Error(w, "The query parameter value "+valueMetrics+" is incorrect", http.StatusBadRequest)
			return
		}
		var sum int64
		if oldVal, _, exists := MemServer.GetMetricsCount(nameMetrics); exists {
			sum = oldVal + val
		} else {
			sum = val
		}

		hashServer := config.GetHashServer(nameMetrics, "counter", val, 0)
		MemServer.PutMetricsCount(nameMetrics, sum, hashServer)
		resp = "Request completed successfully " + nameMetrics + "=" + fmt.Sprint(sum)
	}

	log.Debug(resp)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func GetMetricsHandler(w http.ResponseWriter, r *http.Request) {

	typeMetrics := chi.URLParam(r, "type")
	if typeMetrics == "" {
		log.Error("The query parameter type is missing")
		http.Error(w, "The query parameter type is missing", http.StatusBadRequest)
		return
	}
	if typeMetrics != "gauge" && typeMetrics != "counter" {
		log.Error("The type incorrect " + typeMetrics)
		http.Error(w, "The type incorrect", http.StatusNotImplemented)
		return
	}
	nameMetrics := chi.URLParam(r, "name")
	if nameMetrics == "" {
		log.Error("The query parameter name is missing")
		http.Error(w, "The query parameter name is missing", http.StatusBadRequest)
		return
	}

	var param string
	if typeMetrics == "gauge" {
		val, _, exists := MemServer.GetMetricsGauge(nameMetrics)
		if !exists {
			log.Error("The name " + nameMetrics + " incorrect")
			http.Error(w, "The name "+nameMetrics+" incorrect", http.StatusNotFound)
			return
		}
		param = fmt.Sprint(val)
	}
	if typeMetrics == "counter" {
		val, _, exists := MemServer.GetMetricsCount(nameMetrics)
		if !exists {
			log.Error("The name " + nameMetrics + " incorrect")
			http.Error(w, "The name "+nameMetrics+" incorrect", http.StatusNotFound)
			return
		}
		param = fmt.Sprint(val)
	}

	log.Debug("Request completed successfully " + nameMetrics + "=" + fmt.Sprint(param))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(param))
}

func PostMetricsHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value := storage.JSONMetrics{}
	if err := json.Unmarshal([]byte(string(resp)), &value); err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rawValue := storage.JSONMetrics{
		ID:    value.ID,
		MType: value.MType,
	}
	if value.Value != nil {

		hashServer := config.GetHashServer(value.ID, "gauge", 0, *value.Value)
		if hashServer != value.Hash && config.GetConfigKeyServer() != "" && config.GetConfigKeyAgent() != "" {
			log.Error("Хеш-сумма не соответствует расчетной")
			http.Error(w, "Хеш-сумма не соответствует расчетной", http.StatusBadRequest)
			return
		}
		MemServer.PutMetricsGauge(value.ID, *value.Value, hashServer)
		rawValue.Value = value.Value
		rawValue.Hash = hashServer

	}
	if value.Delta != nil {

		// calculate hash
		hashServer := config.GetHashServer(value.ID, "counter", *value.Delta, 0)
		// compare hashes
		if hashServer != value.Hash && config.GetConfigKeyServer() != "" && config.GetConfigKeyAgent() != "" {
			log.Error("Хеш-сумма не соответствует расчетной")
			http.Error(w, "Хеш-сумма не соответствует расчетной", http.StatusBadRequest)
			return
		}

		// counter summing logic
		var sum int64
		if oldVal, _, exists := MemServer.GetMetricsCount(value.ID); exists {
			sum = oldVal + *value.Delta
		} else {
			sum = *value.Delta
		}
		// calculate new hash
		hashSumServer := config.GetHashServer(value.ID, "counter", sum, 0)
		// store new metric
		MemServer.PutMetricsCount(value.ID, sum, hashSumServer)

		rawValue.Delta = &sum
		rawValue.Hash = hashSumServer
	}

	resp, err = json.Marshal(rawValue)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if config.GetConfigStoreIntervalServer() == 0 {
		FileStore(MemServer)
	}

	log.Debug("Request completed successfully metric:" + value.ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func PostValueMetricsHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value := storage.JSONMetrics{}
	if err := json.Unmarshal([]byte(string(resp)), &value); err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rawValue := storage.JSONMetrics{
		ID:    value.ID,
		MType: value.MType,
	}
	if value.MType == "gauge" {
		val, hash, exists := MemServer.GetMetricsGauge(value.ID)
		if !exists {
			log.Error("Element " + value.ID + " not exists")
			http.Error(w, "Element "+value.ID+" not exists", http.StatusNotFound)
			return
		}
		rawValue.Value = &val
		rawValue.Hash = hash
	}
	if value.MType == "counter" {
		val, hash, exists := MemServer.GetMetricsCount(value.ID)
		if !exists {
			log.Error("Element " + value.ID + " not exists")
			http.Error(w, "Element "+value.ID+" not exists", http.StatusNotFound)
			return
		}
		rawValue.Delta = &val
		rawValue.Hash = hash
	}
	resp, err = json.Marshal(rawValue)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debug("Request completed successfully metric:" + value.ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func RestoreMetricsFromFile() {

	if config.GetConfigRestoreServer() {
		log.Info("Restore metrics")
		FileRestore(MemServer)
	}
}

func StoreMetricsToFile() {

	if config.GetConfigStoreFileServer() != "" {
		ticker := time.NewTicker(config.GetConfigStoreIntervalServer())
		//for range time.Tick(config.GetConfigStoreIntervalServer()) {
		for range ticker.C {
			log.Info("Store metrics")
			FileStore(MemServer)
		}
	}
}

func StartServer(r *chi.Mux) {

	log.Infof("Starting application %v\n", config.GetConfigAddressServer())
	if con := http.ListenAndServe(config.GetConfigAddressServer(), r); con != nil {
		log.Fatal(con)
	}
}
