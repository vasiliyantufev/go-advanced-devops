package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	database "github.com/vasiliyantufev/go-advanced-devops/internal/db"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

type ViewData struct {
	MapG map[string]float64
	MapC map[string]int64
}

type Server struct {
	Mem *storage.MemStorage
}

func NewServer(param *storage.MemStorage) *Server {
	return &Server{Mem: param}
}

func (s Server) Route() *chi.Mux {

	r := chi.NewRouter()
	r.Get("/", s.indexHandler)
	r.Get("/ping", s.pingHandler)
	r.Route("/value", func(r chi.Router) {
		r.Get("/{type}/{name}", s.getMetricsHandler)
		r.Post("/", s.postValueMetricsHandler)
	})
	r.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", s.metricsHandler)
		r.Post("/", s.postMetricHandler)
	})
	r.Post("/updates/", s.postMetricsHandler)

	return r
}

func (s Server) indexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		log.Errorf("Parse failed: %s", err)
		http.Error(w, "Error loading index page", http.StatusInternalServerError)
		return
	}

	gauges := make(map[string]float64)
	counters := make(map[string]int64)

	metrics := s.Mem.GetAllMetrics()

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

func (s Server) metricsHandler(w http.ResponseWriter, r *http.Request) {

	typeMetrics := chi.URLParam(r, "type")
	nameMetrics := chi.URLParam(r, "name")
	valueMetrics := chi.URLParam(r, "value")

	status, err := storage.ValidURLParamMetrics(typeMetrics, nameMetrics, valueMetrics)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), status)
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
		s.Mem.PutMetricsGauge(nameMetrics, val, hashServer)
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
		if oldVal, _, exists := s.Mem.GetMetricsCount(nameMetrics); exists {
			sum = oldVal + val
		} else {
			sum = val
		}

		hashServer := config.GetHashServer(nameMetrics, "counter", val, 0)
		s.Mem.PutMetricsCount(nameMetrics, sum, hashServer)
		resp = "Request completed successfully " + nameMetrics + "=" + fmt.Sprint(sum)
	}

	log.Debug(resp)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (s Server) getMetricsHandler(w http.ResponseWriter, r *http.Request) {

	typeMetrics := chi.URLParam(r, "type")
	nameMetrics := chi.URLParam(r, "name")

	status, err := storage.ValidURLParamGetMetrics(typeMetrics, nameMetrics)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), status)
		return
	}

	var param string
	if typeMetrics == "gauge" {
		val, _, exists := s.Mem.GetMetricsGauge(nameMetrics)
		if !exists {
			log.Error("The name " + nameMetrics + " incorrect")
			http.Error(w, "The name "+nameMetrics+" incorrect", http.StatusNotFound)
			return
		}
		param = fmt.Sprint(val)
	}
	if typeMetrics == "counter" {
		val, _, exists := s.Mem.GetMetricsCount(nameMetrics)
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

func (s Server) postMetricHandler(w http.ResponseWriter, r *http.Request) {

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
		s.Mem.PutMetricsGauge(value.ID, *value.Value, hashServer)
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
		if oldVal, _, exists := s.Mem.GetMetricsCount(value.ID); exists {
			sum = oldVal + *value.Delta
		} else {
			sum = *value.Delta
		}
		// calculate new hash
		hashSumServer := config.GetHashServer(value.ID, "counter", sum, 0)
		// store new metric
		s.Mem.PutMetricsCount(value.ID, sum, hashSumServer)

		rawValue.Delta = &sum
		rawValue.Hash = hashSumServer
	}

	resp, err = json.Marshal(rawValue)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if config.GetConfigDBServer() != "" {
		database.InsertOrUpdateMetrics(s.Mem)
	}
	if config.GetConfigStoreIntervalServer() == 0 {
		FileStore(s.Mem)
	}

	log.Debug("Request completed successfully metric:" + value.ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (s Server) postMetricsHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var metrics []*storage.JSONMetrics
	if err := json.Unmarshal([]byte(string(resp)), &metrics); err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, metric := range metrics {

		if metric.Value != nil {

			hashServer := config.GetHashServer(metric.ID, "gauge", 0, *metric.Value)
			if hashServer != metric.Hash && config.GetConfigKeyServer() != "" && config.GetConfigKeyAgent() != "" {
				log.Error("Хеш-сумма не соответствует расчетной")
				http.Error(w, "Хеш-сумма не соответствует расчетной", http.StatusBadRequest)
				return
			}
			s.Mem.PutMetricsGauge(metric.ID, *metric.Value, hashServer)

		}
		if metric.Delta != nil {

			// calculate hash
			hashServer := config.GetHashServer(metric.ID, "counter", *metric.Delta, 0)
			// compare hashes
			if hashServer != metric.Hash && config.GetConfigKeyServer() != "" && config.GetConfigKeyAgent() != "" {
				log.Error("Хеш-сумма не соответствует расчетной")
				http.Error(w, "Хеш-сумма не соответствует расчетной", http.StatusBadRequest)
				return
			}

			// counter summing logic
			var sum int64
			if oldVal, _, exists := s.Mem.GetMetricsCount(metric.ID); exists {
				sum = oldVal + *metric.Delta
			} else {
				sum = *metric.Delta
			}
			// calculate new hash
			hashSumServer := config.GetHashServer(metric.ID, "counter", sum, 0)
			// store new metric
			s.Mem.PutMetricsCount(metric.ID, sum, hashSumServer)
		}
	}

	if config.GetConfigDBServer() != "" {
		database.InsertOrUpdateMetrics(s.Mem)
	}
	if config.GetConfigStoreIntervalServer() == 0 {
		FileStore(s.Mem)
	}

	log.Debug("Request completed successfully metric")
	w.WriteHeader(http.StatusOK)
}

func (s Server) postValueMetricsHandler(w http.ResponseWriter, r *http.Request) {

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
		val, hash, exists := s.Mem.GetMetricsGauge(value.ID)
		if !exists {
			log.Error("Element " + value.ID + " not exists")
			http.Error(w, "Element "+value.ID+" not exists", http.StatusNotFound)
			return
		}
		rawValue.Value = &val
		rawValue.Hash = hash
	}
	if value.MType == "counter" {
		val, hash, exists := s.Mem.GetMetricsCount(value.ID)
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

func (s Server) RestoreMetricsFromFile() {

	if config.GetConfigRestoreServer() {
		log.Info("Restore metrics")
		FileRestore(s.Mem)
	}
}

func (s Server) StoreMetricsToFile() {

	if config.GetConfigStoreFileServer() != "" && config.GetConfigDBServer() == "" {
		ticker := time.NewTicker(config.GetConfigStoreIntervalServer())
		//for range time.Tick(config.GetConfigStoreIntervalServer()) {
		for range ticker.C {
			log.Info("Store metrics")
			FileStore(s.Mem)
		}
	}
}

func StartServer(r *chi.Mux) {

	log.Infof("Starting application %v\n", config.GetConfigAddressServer())
	if con := http.ListenAndServe(config.GetConfigAddressServer(), r); con != nil {
		log.Fatal(con)
	}
}

func (s Server) pingHandler(w http.ResponseWriter, r *http.Request) {

	if err := database.Ping(); err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Info("ping")
	w.WriteHeader(http.StatusOK)
}
