package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(storage.MetricsGauge)
	log.Print(storage.MetricsCounter)
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {

	typeMetrics := chi.URLParam(r, "type")
	nameMetrics := chi.URLParam(r, "name")
	valueMetrics := chi.URLParam(r, "value")

	log.Print(typeMetrics)
	log.Print(nameMetrics)
	log.Print(valueMetrics)

	//str := r.URL.String()
	//res := strings.Split(str, "/")
	//typeMetrics := res[1]
	//nameMetrics := res[2]
	//valueMetrics := res[3]

	//if typeMetrics == "gauge" {
	//	val, err := strconv.ParseFloat(string(valueMetrics), 64)
	//	if err != nil {
	//		panic(err)
	//	}
	//	storage.MetricsGauge[nameMetrics] = val
	//}
	//if typeMetrics == "counter" {
	//	val, err := strconv.ParseInt(string(valueMetrics), 10, 64)
	//	if err != nil {
	//		panic(err)
	//	}
	//	storage.MetricsCounter[nameMetrics] = val
	//}

	log.Fatal()

	//short := mux.Vars(r)

	//valueID := chi.URLParam(r, "value")
	//
	//log.Print(valueID)

	//if short["id"] == "" {
	//	http.Error(w, "The query parameter is missing", http.StatusBadRequest)
	//	return
	//}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>MetricsGaugeHandler</h1>"))
}

func MetricsCounterHandler(w http.ResponseWriter, r *http.Request) {

	//resp, err := io.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//s := r.URL.String()
	//key := s[strings.LastIndex(s, "/")+1:]
	//
	//f, err := strconv.ParseInt(string(resp), 10, 64)
	//if err != nil {
	//	panic(err)
	//}
	//
	//storage.MetricsCounter[key] = f

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>MetricsCounterHandler</h1>"))
}
