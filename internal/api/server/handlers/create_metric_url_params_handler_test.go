package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"
)

//func TestHandler_CreateMetricURLParamsHandler(t *testing.T) {
//
//	testTable := []struct {
//		name        string
//		server      *httptest.Server
//		expectedErr error
//	}{
//		{
//			name: "test create metric url params handler",
//			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//				w.WriteHeader(http.StatusOK)
//			})),
//			//server:      httptest.NewServer(Handler.CreateMetricURLParamsHandler),
//			expectedErr: nil,
//		},
//	}
//	for _, tc := range testTable {
//		t.Run(tc.name, func(t *testing.T) {
//			defer tc.server.Close()
//			resp, err := MakeHTTPCall(tc.server.URL)
//			if err != tc.expectedErr {
//				t.Error(err)
//			}
//			defer resp.Body.Close()
//
//			//assert.Equal(t, resp.StatusCode, http.StatusOK)
//		})
//	}
//}

func TestHandler_CreateMetricURLParamsGaugeHandler(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

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

	rand.Seed(time.Now().UnixNano())
	var valueExpect = fmt.Sprint(rand.Int())
	var statusExpect = http.StatusOK

	router.ServeHTTP(responseRecorder, httptest.NewRequest("POST", "/update/gauge/testMetric/"+fmt.Sprint(valueExpect), nil))
	statusGet := responseRecorder.Code

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
}

func TestHandler_CreateMetricURLParamsCountHandler(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

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

	rand.Seed(time.Now().UnixNano())
	var valueExpect = fmt.Sprintf("%.3f", rand.Float64())
	var statusExpect = http.StatusOK

	router.ServeHTTP(responseRecorder, httptest.NewRequest("POST", "/update/count/testMetric/"+fmt.Sprint(valueExpect), nil))
	statusGet := responseRecorder.Code

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
}

func TestHandler_CreateMetricURLParamsGaugeValueIncorrectHandler(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

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

	var valueIncorrect = "bugaga"
	var statusExpect = http.StatusBadRequest

	router.ServeHTTP(responseRecorder, httptest.NewRequest("POST", "/update/gauge/testMetric/"+valueIncorrect, nil))
	statusGet := responseRecorder.Code

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
}

func TestHandler_CreateMetricURLParamsCountValueIncorrectHandler(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

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

	var valueIncorrect = "bugaga"
	var statusExpect = http.StatusBadRequest

	router.ServeHTTP(responseRecorder, httptest.NewRequest("POST", "/update/counter/testMetric/"+valueIncorrect, nil))
	statusGet := responseRecorder.Code

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
}

func TestHandler_CreateMetricURLParamsCountSumHandler(t *testing.T) {
	responseRecorderPostFirst := httptest.NewRecorder()
	responseRecorderPostSecond := httptest.NewRecorder()

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

	rand.Seed(time.Now().UnixNano())
	var value = 1234567890
	var statusExpect = http.StatusOK

	router.ServeHTTP(responseRecorderPostFirst, httptest.NewRequest("POST", "/update/counter/testMetric/"+fmt.Sprint(value), nil))
	router.ServeHTTP(responseRecorderPostSecond, httptest.NewRequest("POST", "/update/counter/testMetric/"+fmt.Sprint(value), nil))
	var valueGet = strings.Split(responseRecorderPostSecond.Body.String(), "=")
	statusGet := responseRecorderPostSecond.Code

	var valueExpect = fmt.Sprint(value + value)

	assert.Equal(t, valueExpect, valueGet[1], fmt.Sprintf("Incorrect body. Expect %s, got %s", valueExpect, valueGet))
	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
}
