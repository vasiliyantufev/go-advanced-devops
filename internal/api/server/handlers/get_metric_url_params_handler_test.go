package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"
)

func TestHandler_GetMetricURLParamsCounterHandler(t *testing.T) {

	responseRecorderPost := httptest.NewRecorder()
	responseRecorderGet := httptest.NewRecorder()

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer("secretKey")

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
	router.Post("/update/{type}/{name}/{value}", srv.CreateMetricURLParamsHandler)
	router.Get("/value/{type}/{name}", srv.GetMetricURLParamsHandler)

	rand.Seed(time.Now().UnixNano())
	var valueExpect = strconv.FormatInt(rand.Int63(), 10)
	var statusExpect = http.StatusOK

	router.ServeHTTP(responseRecorderPost, httptest.NewRequest("POST", "/update/counter/testMetric/"+fmt.Sprint(valueExpect), nil))
	router.ServeHTTP(responseRecorderGet, httptest.NewRequest("GET", "/value/counter/testMetric", nil))
	valueGet := responseRecorderGet.Body.String()
	statusGet := responseRecorderGet.Code

	assert.Equal(t, valueExpect, valueGet, fmt.Sprintf("Incorrect body. Expect %s, got %s", valueExpect, valueGet))
	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
}

func TestHandler_GetMetricURLParamsGaugeHandler(t *testing.T) {

	responseRecorderPost := httptest.NewRecorder()
	responseRecorderGet := httptest.NewRecorder()

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer("secretKey")

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
	router.Post("/update/{type}/{name}/{value}", srv.CreateMetricURLParamsHandler)
	router.Get("/value/{type}/{name}", srv.GetMetricURLParamsHandler)

	rand.Seed(time.Now().UnixNano())
	var valueExpect = strconv.FormatFloat(rand.Float64(), 'f', -1, 64)
	var statusExpect = http.StatusOK

	router.ServeHTTP(responseRecorderPost, httptest.NewRequest("POST", "/update/gauge/testMetric/"+fmt.Sprint(valueExpect), nil))
	router.ServeHTTP(responseRecorderGet, httptest.NewRequest("GET", "/value/gauge/testMetric", nil))
	valueGet := responseRecorderGet.Body.String()
	statusGet := responseRecorderGet.Code

	assert.Equal(t, valueExpect, valueGet, fmt.Sprintf("Incorrect body. Expect %s, got %s", valueExpect, valueGet))
	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
}

func TestHandler_GetMetricURLParamsNotExistCounterHandler(t *testing.T) {

	responseRecorderGet := httptest.NewRecorder()

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer("secretKey")

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
	router.Get("/value/{type}/{name}", srv.GetMetricURLParamsHandler)

	var statusExpect = http.StatusNotFound
	var contentTypeExpect = "text/plain; charset=utf-8"

	router.ServeHTTP(responseRecorderGet, httptest.NewRequest("GET", "/value/counter/testMetric", nil))
	statusGet := responseRecorderGet.Code
	contentTypeGet := responseRecorderGet.Header().Get("Content-Type")

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
	assert.Equal(t, contentTypeExpect, contentTypeGet, fmt.Sprintf("Incorrect Content-Type. Expect %s, got %s", contentTypeExpect, contentTypeGet))
}

func TestHandler_GetMetricURLParamsNotExistGaugeHandler(t *testing.T) {

	responseRecorderGet := httptest.NewRecorder()

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer("secretKey")

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
	router.Post("/update/{type}/{name}/{value}", srv.CreateMetricURLParamsHandler)
	router.Get("/value/{type}/{name}", srv.GetMetricURLParamsHandler)

	var statusExpect = http.StatusNotFound
	var contentTypeExpect = "text/plain; charset=utf-8"

	router.ServeHTTP(responseRecorderGet, httptest.NewRequest("GET", "/value/gauge/testMetric", nil))
	statusGet := responseRecorderGet.Code
	contentTypeGet := responseRecorderGet.Header().Get("Content-Type")

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
	assert.Equal(t, contentTypeExpect, contentTypeGet, fmt.Sprintf("Incorrect Content-Type. Expect %s, got %s", contentTypeExpect, contentTypeGet))
}
