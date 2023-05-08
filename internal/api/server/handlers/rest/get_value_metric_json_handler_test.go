package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	"github.com/vasiliyantufev/go-advanced-devops/internal/converter"
	"github.com/vasiliyantufev/go-advanced-devops/internal/models"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"
)

func TestHandler_GetValueGaugeMetricJSONHandler(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer("")

	configServer := configserver.ConfigServer{
		Address:         "localhost:8080",
		AddressPProfile: "localhost:8088",
		Restore:         true,
		StoreInterval:   300 * time.Second,
		DebugLevel:      logrus.DebugLevel,
		StoreFile:       "/tmp/devops-metrics-db.json",
		Key:             "",
		DSN:             "",
		MigrationsPath:  "file://./migrations",
	}

	srv := NewHandler(memStorage, nil, &configServer, nil, hashServer)

	router := chi.NewRouter()
	router.Post("/value", srv.GetValueMetricJSONHandler)

	rand.Seed(time.Now().UnixNano())
	var value = rand.Float64()
	var statusExpect = http.StatusOK
	var contentTypeExpect = "application/json"
	var metricGauge = models.Metric{
		ID:    "alloc1",
		MType: "gauge",
		Value: &value,
	}
	srv.memStorage.PutMetricsGauge(metricGauge.ID, *metricGauge.Value, hashServer.GenerateHash(models.Metric{ID: metricGauge.ID, MType: metricGauge.MType, Delta: nil, Value: converter.Float64ToFloat64Pointer(*metricGauge.Value)}))

	reqBody, err := json.Marshal(metricGauge)
	if err != nil {
		logrus.Fatal(err)
	}

	router.ServeHTTP(responseRecorder, httptest.NewRequest(http.MethodPost, "/value", bytes.NewBuffer(reqBody)))
	statusGet := responseRecorder.Code
	contentTypeGet := responseRecorder.Header().Get("Content-Type")

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
	assert.Equal(t, contentTypeExpect, contentTypeGet, fmt.Sprintf("Incorrect Content-Type. Expect %s, got %s", contentTypeExpect, contentTypeGet))
}

func TestHandler_GetValueCountMetricJSONHandler(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer("")

	configServer := configserver.ConfigServer{
		Address:         "localhost:8080",
		AddressPProfile: "localhost:8088",
		Restore:         true,
		StoreInterval:   300 * time.Second,
		DebugLevel:      logrus.DebugLevel,
		StoreFile:       "/tmp/devops-metrics-db.json",
		Key:             "",
		DSN:             "",
		MigrationsPath:  "file://./migrations",
	}

	srv := NewHandler(memStorage, nil, &configServer, nil, hashServer)

	router := chi.NewRouter()
	router.Post("/value", srv.GetValueMetricJSONHandler)

	rand.Seed(time.Now().UnixNano())
	var delta = rand.Int63()
	var statusExpect = http.StatusOK
	var contentTypeExpect = "application/json"
	var metricCount = models.Metric{
		ID:    "alloc2",
		MType: "counter",
		Delta: &delta,
	}
	srv.memStorage.PutMetricsCount(metricCount.ID, *metricCount.Delta, hashServer.GenerateHash(models.Metric{ID: metricCount.ID, MType: metricCount.MType, Delta: converter.Int64ToInt64Pointer(*metricCount.Delta), Value: nil}))

	reqBody, err := json.Marshal(metricCount)
	if err != nil {
		logrus.Fatal(err)
	}

	router.ServeHTTP(responseRecorder, httptest.NewRequest(http.MethodPost, "/value", bytes.NewBuffer(reqBody)))
	statusGet := responseRecorder.Code
	contentTypeGet := responseRecorder.Header().Get("Content-Type")

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
	assert.Equal(t, contentTypeExpect, contentTypeGet, fmt.Sprintf("Incorrect Content-Type. Expect %s, got %s", contentTypeExpect, contentTypeGet))
}

func TestHandler_GetValueGaugeNoExistsMetricJSONHandler(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer("")

	configServer := configserver.ConfigServer{
		Address:         "localhost:8080",
		AddressPProfile: "localhost:8088",
		Restore:         true,
		StoreInterval:   300 * time.Second,
		DebugLevel:      logrus.DebugLevel,
		StoreFile:       "/tmp/devops-metrics-db.json",
		Key:             "",
		DSN:             "",
		MigrationsPath:  "file://./migrations",
	}

	srv := NewHandler(memStorage, nil, &configServer, nil, hashServer)

	router := chi.NewRouter()
	router.Post("/value", srv.GetValueMetricJSONHandler)

	rand.Seed(time.Now().UnixNano())
	var value = rand.Float64()
	var statusExpect = http.StatusNotFound
	var metricGauge = models.Metric{
		ID:    "alloc1",
		MType: "gauge",
		Value: &value,
	}

	reqBody, err := json.Marshal(metricGauge)
	if err != nil {
		logrus.Fatal(err)
	}

	router.ServeHTTP(responseRecorder, httptest.NewRequest(http.MethodPost, "/value", bytes.NewBuffer(reqBody)))
	statusGet := responseRecorder.Code

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
}

func TestHandler_GetValueCountNoExistsMetricJSONHandler(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer("")

	configServer := configserver.ConfigServer{
		Address:         "localhost:8080",
		AddressPProfile: "localhost:8088",
		Restore:         true,
		StoreInterval:   300 * time.Second,
		DebugLevel:      logrus.DebugLevel,
		StoreFile:       "/tmp/devops-metrics-db.json",
		Key:             "",
		DSN:             "",
		MigrationsPath:  "file://./migrations",
	}

	srv := NewHandler(memStorage, nil, &configServer, nil, hashServer)

	router := chi.NewRouter()
	router.Post("/value", srv.GetValueMetricJSONHandler)

	rand.Seed(time.Now().UnixNano())
	var delta = rand.Int63()
	var statusExpect = http.StatusNotFound
	var metricCount = models.Metric{
		ID:    "alloc2",
		MType: "counter",
		Delta: &delta,
	}

	reqBody, err := json.Marshal(metricCount)
	if err != nil {
		logrus.Fatal(err)
	}

	router.ServeHTTP(responseRecorder, httptest.NewRequest(http.MethodPost, "/value", bytes.NewBuffer(reqBody)))
	statusGet := responseRecorder.Code

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
}
