package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"html/template"
	"net/http"
	"strconv"
)

var MemServer = storage.NewMemStorage()

type ViewData struct {
	MapG map[string]float64
	MapC map[string]int64
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	data := ViewData{
		MapG: MemServer.DataMetricsGauge,
		MapC: MemServer.DataMetricsCount,
	}

	tmpl, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		log.Errorf("Parse failed: %s", err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Errorf("Execution failed: %s", err)
	}
	w.WriteHeader(http.StatusOK)

}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {

	typeMetrics := chi.URLParam(r, "type")
	if typeMetrics == "" {
		log.Error("The query parameter type is missing")
		http.Error(w, "The query parameter type is missing", http.StatusBadRequest)
		return
	}
	if typeMetrics != "gauge" && typeMetrics != "counter" {
		log.Error("The type incorrect")
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

	if typeMetrics == "gauge" {
		val, err := strconv.ParseFloat(string(valueMetrics), 64)
		if err != nil {
			log.Error("The query parameter value " + valueMetrics + " is incorrect")
			http.Error(w, "The query parameter value "+valueMetrics+" is incorrect", http.StatusBadRequest)
			return
		}
		MemServer.PutMetricsGauge(nameMetrics, val)
	}
	if typeMetrics == "counter" {
		val, err := strconv.ParseInt(string(valueMetrics), 10, 64)
		if err != nil {
			log.Error("The query parameter value " + valueMetrics + " is incorrect")
			http.Error(w, "The query parameter value "+valueMetrics+" is incorrect", http.StatusBadRequest)
			return
		}
		var sum int64
		sum = 0
		for _, val := range MemServer.DataMetricsCount {
			sum = sum + val
		}
		MemServer.PutMetricsCount(nameMetrics, sum+val)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>MetricsHandler</h1>"))
}

func GetMetricsHandler(w http.ResponseWriter, r *http.Request) {

	typeMetrics := chi.URLParam(r, "type")
	if typeMetrics == "" {
		log.Error("The query parameter type is missing")
		http.Error(w, "The query parameter type is missing", http.StatusBadRequest)
		return
	}
	if typeMetrics != "gauge" && typeMetrics != "counter" {
		log.Error("The type incorrect")
		http.Error(w, "The type incorrect", http.StatusNotImplemented)
		return
	}
	nameMetrics := chi.URLParam(r, "name")
	if nameMetrics == "" {
		log.Error("The query parameter name is missing")
		http.Error(w, "The query parameter name is missing", http.StatusBadRequest)
		return
	}
	_, exists1 := MemServer.GetMetricsGauge(nameMetrics)
	_, exists2 := MemServer.GetMetricsCount(nameMetrics)
	if !exists1 && !exists2 {
		log.Error("The name " + nameMetrics + " incorrect")
		http.Error(w, "The name "+nameMetrics+" incorrect", http.StatusNotFound)
		return
	}

	var resp string
	if typeMetrics == "gauge" {
		val, _ := MemServer.GetMetricsGauge(nameMetrics)
		resp = fmt.Sprint(val)
	}
	if typeMetrics == "counter" {
		val, _ := MemServer.GetMetricsCount(nameMetrics)
		resp = fmt.Sprint(val)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
