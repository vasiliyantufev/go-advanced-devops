package app

import (
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(storage.MetricsGauge)
	log.Print(storage.MetricsCounter)
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
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

	storage.MetricsGauge[key] = f

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

	storage.MetricsCounter[key] = f

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>MetricsCounterHandler</h1>"))
}
