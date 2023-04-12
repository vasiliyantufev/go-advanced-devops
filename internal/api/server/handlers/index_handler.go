package handlers

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
	_ "github.com/vasiliyantufev/go-advanced-devops/internal/api/server"
)

type ViewData struct {
	MapG map[string]float64
	MapC map[string]int64
}

// IndexHandler - the page that displays all the metrics with parameters
func (s Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		log.Errorf("Parse failed: %s", err)
		http.Error(w, "Error loading index page", http.StatusInternalServerError)
		return
	}

	gauges := make(map[string]float64)
	counters := make(map[string]int64)

	metrics := s.memStorage.GetAllMetrics()

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
