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
	"github.com/vasiliyantufev/go-advanced-devops/internal/model"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"
)

func TestHandler_CreateMetricJSONGaugeHandler(t *testing.T) {
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
	router.Post("/update", srv.CreateMetricJSONHandler)

	rand.Seed(time.Now().UnixNano())
	var value = rand.Float64()

	var statusExpect = http.StatusOK
	var metricExt = model.Metric{
		ID:    "alloc",
		MType: "gauge",
		Value: &value,
	}
	reqBody, err := json.Marshal(metricExt)
	if err != nil {
		logrus.Error(err)
	}

	router.ServeHTTP(responseRecorder, httptest.NewRequest(http.MethodPost, "/update", bytes.NewBuffer(reqBody)))
	statusGet := responseRecorder.Code
	metricGet := model.Metric{}
	err = json.Unmarshal([]byte(responseRecorder.Body.Bytes()), &metricGet)
	if err != nil {
		logrus.Error(err)
	}

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
	assert.Equal(t, metricExt.ID, metricGet.ID, fmt.Sprintf("Incorrect ID metric. Expect %s, got %s", metricExt.ID, metricGet.ID))
	assert.Equal(t, metricExt.MType, metricGet.MType, fmt.Sprintf("Incorrect type metric. Expect %s, got %s", metricExt.MType, metricGet.MType))
	assert.Equal(t, metricExt.Value, metricGet.Value, fmt.Sprintf("Incorrect value metric. Expect %d, got %d", metricExt.Value, metricGet.Value))
	assert.True(t, len(metricGet.Hash) > 0, "Empty hash metric")
}

func TestHandler_CreateMetricJSONCountHandler(t *testing.T) {
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
	router.Post("/update", srv.CreateMetricJSONHandler)

	rand.Seed(time.Now().UnixNano())
	var value = rand.Int63()
	var statusExpect = http.StatusOK
	var metricExt = model.Metric{
		ID:    "alloc",
		MType: "count",
		Delta: &value,
	}
	reqBody, err := json.Marshal(metricExt)
	if err != nil {
		logrus.Error(err)
	}

	router.ServeHTTP(responseRecorder, httptest.NewRequest(http.MethodPost, "/update", bytes.NewBuffer(reqBody)))
	statusGet := responseRecorder.Code
	metricGet := model.Metric{}
	err = json.Unmarshal([]byte(responseRecorder.Body.Bytes()), &metricGet)
	if err != nil {
		logrus.Error(err)
	}

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
	assert.Equal(t, metricExt.ID, metricGet.ID, fmt.Sprintf("Incorrect ID metric. Expect %s, got %s", metricExt.ID, metricGet.ID))
	assert.Equal(t, metricExt.MType, metricGet.MType, fmt.Sprintf("Incorrect type metric. Expect %s, got %s", metricExt.MType, metricGet.MType))
	assert.Equal(t, metricExt.Value, metricGet.Value, fmt.Sprintf("Incorrect value metric. Expect %d, got %d", metricExt.Value, metricGet.Value))
	assert.True(t, len(metricGet.Hash) > 0, "Empty hash metric")
}

func TestHandler_CreateMetricJSONGaugeKeyIncorrectHandler(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer("bugagaKey")

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
	router.Post("/update", srv.CreateMetricJSONHandler)

	rand.Seed(time.Now().UnixNano())
	var value = rand.Float64()

	var statusExpect = http.StatusBadRequest
	var metricExt = model.Metric{
		ID:    "alloc",
		MType: "gauge",
		Value: &value,
	}
	reqBody, err := json.Marshal(metricExt)
	if err != nil {
		logrus.Fatal(err)
	}

	router.ServeHTTP(responseRecorder, httptest.NewRequest(http.MethodPost, "/update", bytes.NewBuffer(reqBody)))
	statusGet := responseRecorder.Code

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))

}

func TestHandler_CreateMetricJSONCountKeyIncorrectHandler(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer("bugagaKey")

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
	router.Post("/update", srv.CreateMetricJSONHandler)

	rand.Seed(time.Now().UnixNano())
	var value = rand.Int63()
	var statusExpect = http.StatusBadRequest
	var metricExt = model.Metric{
		ID:    "alloc",
		MType: "count",
		Delta: &value,
	}
	reqBody, err := json.Marshal(metricExt)
	if err != nil {
		logrus.Fatal(err)
	}

	router.ServeHTTP(responseRecorder, httptest.NewRequest(http.MethodPost, "/update", bytes.NewBuffer(reqBody)))
	statusGet := responseRecorder.Code

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
}

func TestHandler_CreateMetricJSONCountSumHandler(t *testing.T) {
	responseRecorderPostFirst := httptest.NewRecorder()
	responseRecorderPostSecond := httptest.NewRecorder()

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
	router.Post("/update", srv.CreateMetricJSONHandler)

	rand.Seed(time.Now().UnixNano())
	var value int64 = 1234567890
	var statusExpect = http.StatusOK
	var metricExt = model.Metric{
		ID:    "alloc",
		MType: "count",
		Delta: &value,
	}
	reqBody, err := json.Marshal(metricExt)
	if err != nil {
		logrus.Error(err)
	}

	router.ServeHTTP(responseRecorderPostFirst, httptest.NewRequest(http.MethodPost, "/update", bytes.NewBuffer(reqBody)))
	router.ServeHTTP(responseRecorderPostSecond, httptest.NewRequest(http.MethodPost, "/update", bytes.NewBuffer(reqBody)))
	statusGet := responseRecorderPostSecond.Code
	metricGet := model.Metric{}
	err = json.Unmarshal([]byte(responseRecorderPostSecond.Body.Bytes()), &metricGet)
	if err != nil {
		logrus.Error(err)
	}

	var valueExpect = value + value

	assert.Equal(t, valueExpect, *metricGet.Delta, fmt.Sprintf("Incorrect Delta. Expect %d, got %d", valueExpect, *metricGet.Delta))
	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
}
