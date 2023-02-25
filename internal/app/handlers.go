package app

import (
	"encoding/json"
	"fmt"
	"github.com/vasiliyantufev/go-advanced-devops/internal/converter"
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
	mem        *storage.MemStorage
	config     *config.ConfigServer
	database   *database.DB
	hashServer *HashServer
}

func NewServer(mem *storage.MemStorage, cfg *config.ConfigServer, db *database.DB, hash *HashServer) *Server {
	return &Server{mem: mem, config: cfg, database: db, hashServer: hash}
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

	metrics := s.mem.GetAllMetrics()

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
		hashServer := s.hashServer.GenerateHash(storage.JSONMetrics{ID: nameMetrics, MType: "gauge", Delta: nil, Value: converter.Float64ToFloat64Pointer(val)})
		s.mem.PutMetricsGauge(nameMetrics, val, hashServer)
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
		if oldVal, _, exists := s.mem.GetMetricsCount(nameMetrics); exists {
			sum = oldVal + val
		} else {
			sum = val
		}
		hashServer := s.hashServer.GenerateHash(storage.JSONMetrics{ID: nameMetrics, MType: "counter", Delta: converter.Int64ToInt64Pointer(val), Value: nil})
		s.mem.PutMetricsCount(nameMetrics, sum, hashServer)
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
		val, _, exists := s.mem.GetMetricsGauge(nameMetrics)
		if !exists {
			log.Error("The name " + nameMetrics + " incorrect")
			http.Error(w, "The name "+nameMetrics+" incorrect", http.StatusNotFound)
			return
		}
		param = fmt.Sprintf("%.3f", val)
	}
	if typeMetrics == "counter" {
		val, _, exists := s.mem.GetMetricsCount(nameMetrics)
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
		ID: value.ID,
	}
	if value.Value != nil {
		// expect true id hashServer is disabled
		if isValid := s.hashServer.ValidHashServer(value); isValid {
			log.Println("GAUGE METRIC", isValid)
		}

		if !s.hashServer.ValidHashServer(value) {
			log.Error("Хеш-сумма не соответствует расчетной")
			http.Error(w, "Хеш-сумма не соответствует расчетной", http.StatusBadRequest)
			return
		}

		hashServer := s.hashServer.GenerateHash(value)
		s.mem.PutMetricsGauge(value.ID, *value.Value, hashServer)
		rawValue.Value = value.Value
		rawValue.Hash = hashServer
	}
	if value.Delta != nil {
		// expect true id hashServer is disabled
		if isValid := s.hashServer.ValidHashServer(value); isValid {
			log.Println("GAUGE METRIC", isValid)
		}
		// compare hashes
		if !s.hashServer.ValidHashServer(value) {
			log.Error("Хеш-сумма не соответствует расчетной")
			http.Error(w, "Хеш-сумма не соответствует расчетной", http.StatusBadRequest)
			return
		}

		// counter summing logic
		var sum int64
		if oldVal, _, exists := s.mem.GetMetricsCount(value.ID); exists {
			sum = oldVal + *value.Delta
		} else {
			sum = *value.Delta
		}
		// calculate new hash
		hashSumServer := s.hashServer.GenerateHash(storage.JSONMetrics{ID: value.ID, MType: value.MType, Delta: converter.Int64ToInt64Pointer(sum), Value: value.Value})
		// store new metric
		s.mem.PutMetricsCount(value.ID, sum, hashSumServer)

		rawValue.Delta = &sum
		rawValue.Hash = hashSumServer

	}

	resp, err = json.Marshal(rawValue)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if s.database != nil {
		s.database.InsertOrUpdateMetrics(s.mem)
	}
	if s.config.GetConfigStoreIntervalServer() == 0 {
		FileStore(s.mem, nil)
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

			if !s.hashServer.ValidHashServer(*metric) {
				log.Error("Хеш-сумма не соответствует расчетной")
				http.Error(w, "Хеш-сумма не соответствует расчетной", http.StatusBadRequest)
				return
			}
			s.mem.PutMetricsGauge(metric.ID, *metric.Value, s.hashServer.GenerateHash(*metric))

		}
		if metric.Delta != nil {

			// compare hashes
			if !s.hashServer.ValidHashServer(*metric) {
				log.Error("Хеш-сумма не соответствует расчетной")
				http.Error(w, "Хеш-сумма не соответствует расчетной", http.StatusBadRequest)
				return
			}
			// counter summing logic
			var sum int64
			if oldVal, _, exists := s.mem.GetMetricsCount(metric.ID); exists {
				sum = oldVal + *metric.Delta
			} else {
				sum = *metric.Delta
			}
			// calculate new hash
			hashSumServer := s.hashServer.GenerateHash(storage.JSONMetrics{ID: metric.ID, MType: metric.MType, Delta: converter.Int64ToInt64Pointer(sum), Value: metric.Value})
			// store new metric
			s.mem.PutMetricsCount(metric.ID, sum, hashSumServer)
		}
	}

	if s.database != nil {
		s.database.InsertOrUpdateMetrics(s.mem)
	}
	if s.config.GetConfigStoreIntervalServer() == 0 {
		FileStore(s.mem, nil)
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
	log.Infoln("VALUE METRIC RECV", rawValue)
	if value.MType == "gauge" {
		val, hash, exists := s.mem.GetMetricsGauge(value.ID)
		if !exists {
			log.Error("Element " + value.ID + " not exists")
			http.Error(w, "Element "+value.ID+" not exists", http.StatusNotFound)
			return
		}
		rawValue.Value = &val
		rawValue.Hash = hash
	}
	if value.MType == "counter" {
		val, hash, exists := s.mem.GetMetricsCount(value.ID)
		if !exists {
			log.Error("Element " + value.ID + " not exists")
			http.Error(w, "Element "+value.ID+" not exists", http.StatusNotFound)
			return
		}
		rawValue.Delta = &val
		rawValue.Hash = hash
	}
	log.Infoln("VALUE METRIC RESPONSE", rawValue)
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

	if s.config.GetConfigRestoreServer() {
		log.Info("Restore metrics")
		FileRestore(s.mem, s.config)
	}
}

func (s Server) StoreMetricsToFile() {

	if s.config.GetConfigStoreFileServer() != "" && s.config.GetConfigDBServer() == "" {
		ticker := time.NewTicker(s.config.GetConfigStoreIntervalServer())
		//for range time.Tick(config.GetConfigStoreIntervalServer()) {
		for range ticker.C {
			log.Info("Store metrics")
			FileStore(s.mem, s.config)
		}
	}
}

func StartServer(r *chi.Mux, config *config.ConfigServer) {

	log.Infof("Starting application %v\n", config.GetConfigAddressServer())
	if con := http.ListenAndServe(config.GetConfigAddressServer(), r); con != nil {
		log.Fatal(con)
	}
}

func (s Server) pingHandler(w http.ResponseWriter, r *http.Request) {

	if s.database != nil {
		if err := s.database.Ping(); err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Info("ping")
		w.WriteHeader(http.StatusOK)
		return
	}
	log.Error("db is nil")
	w.WriteHeader(http.StatusInternalServerError)
}

func (s Server) GetMem() *storage.MemStorage {
	return s.mem
}

func (s Server) GetConfig() *config.ConfigServer {
	return s.config
}
