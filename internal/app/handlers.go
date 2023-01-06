package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"log"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(storage.MetricsGauge)
	log.Print(storage.MetricsCounter)
}

// http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>;
func MetricsHandler(w http.ResponseWriter, r *http.Request) {

	typeMetrics := chi.URLParam(r, "type")
	if typeMetrics == "" {
		http.Error(w, "The query parameter type is missing", http.StatusBadRequest)
		return
	}
	if typeMetrics != "gauge" && typeMetrics != "counter" {
		http.Error(w, "The type incorrect", http.StatusNotImplemented)
		return
	}

	nameMetrics := chi.URLParam(r, "name")
	if nameMetrics == "" {
		http.Error(w, "The query parameter name is missing", http.StatusBadRequest)
		return
	}
	_, exists1 := storage.MetricsGauge[nameMetrics]
	_, exists2 := storage.MetricsCounter[nameMetrics]
	if !exists1 && !exists2 {
		http.Error(w, "The name "+nameMetrics+" incorrect", http.StatusBadRequest)
		return
	}

	valueMetrics := chi.URLParam(r, "value")
	if valueMetrics == "" {
		http.Error(w, "The query parameter value is missing", http.StatusBadRequest)
		return
	}

	//log.Print(typeMetrics)
	//log.Print(nameMetrics)
	//log.Print(valueMetrics)

	if typeMetrics == "gauge" {
		val, err := strconv.ParseFloat(string(valueMetrics), 64)
		if err != nil {
			panic(err)
		}
		storage.MetricsGauge[nameMetrics] = val
	}
	if typeMetrics == "counter" {
		val, err := strconv.ParseInt(string(valueMetrics), 10, 64)
		if err != nil {
			panic(err)
		}
		storage.MetricsCounter[nameMetrics] = val
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>MetricsHandler</h1>"))
}
