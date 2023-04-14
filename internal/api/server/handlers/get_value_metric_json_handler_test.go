package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"
)

/*
func TestHandler_GetValueMetricJSONHandler(t *testing.T) {


	testTable := []struct {
		name             string
		server           *httptest.Server
		expectedResponse *Response
		expectedErr      error
	}{
		{
			name: "test get value metric json handler",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				result, err := json.Marshal(metric)
				if err != nil {
					t.Error(err.Error())
					return
				}
				w.Write([]byte(result))
			})),
			expectedResponse: metric,
			expectedErr:      nil,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.server.Close()
			resp, respBody, err := MakeHTTPWithBodyJSONCall(tc.server.URL)
			if err != nil {
				t.Error(err)
			}
			defer resp.Body.Close()

			assert.Equal(t, resp.StatusCode, http.StatusOK)
			assert.Equal(t, resp.Header.Get("Content-Type"), "application/json")
			assert.Equal(t, respBody, tc.expectedResponse)
		})
	}
}
*/

func TestHandler_GetValueMetricJSONHandler(t *testing.T) {

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
