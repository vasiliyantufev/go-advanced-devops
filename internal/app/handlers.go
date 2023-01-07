package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type ViewData struct {
	MapG map[string]float64
	MapC map[string]int64
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	data := ViewData{
		MapG: storage.MetricsGauge,
		MapC: storage.MetricsCounter,
	}

	tmpl, _ := template.ParseFiles("./web/templates/index.html")
	//err := tmpl.Execute(w, storage.MetricsGauge)
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
	w.WriteHeader(http.StatusOK)

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
			http.Error(w, "The query parameter value "+valueMetrics+" is incorrect", http.StatusBadRequest)
			return
		}
		storage.MetricsGauge[nameMetrics] = val
	}
	if typeMetrics == "counter" {
		val, err := strconv.ParseInt(string(valueMetrics), 10, 64)
		if err != nil {
			http.Error(w, "The query parameter value "+valueMetrics+" is incorrect", http.StatusBadRequest)
			return
		}

		var sum int64
		sum = 0
		for _, val := range storage.MetricsCounter {
			sum = sum + val
		}
		storage.MetricsCounter[nameMetrics] = sum + val
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>MetricsHandler</h1>"))
}

func GetMetricsHandler(w http.ResponseWriter, r *http.Request) {

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
		http.Error(w, "The name "+nameMetrics+" incorrect", http.StatusNotFound)
		return
	}

	var resp string
	if typeMetrics == "gauge" {
		resp = fmt.Sprint(storage.MetricsGauge[nameMetrics])
	}
	if typeMetrics == "counter" {
		resp = fmt.Sprint(storage.MetricsCounter[nameMetrics])
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
