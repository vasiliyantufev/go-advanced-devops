package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"
)

func TestHandler_IndexHandler(t *testing.T) {

	responseRecorder := httptest.NewRecorder()
	responseRecorderPost := httptest.NewRecorder()

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer("")

	dirname, err := os.Getwd()
	if err != nil {
		log.Error(err)
	}
	dir, err := os.Open(path.Join(dirname, "../../../../"))
	if err != nil {
		log.Error(err)
	}
	tmplFile := dir.Name() + "/web/templates/index.html"

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
		TemplatePath:    tmplFile,
	}

	srv := NewHandler(memStorage, nil, &configServer, nil, hashServer)

	router := chi.NewRouter()
	router.Post("/update/{type}/{name}/{value}", srv.CreateMetricURLParamsHandler)
	router.Get("/", srv.IndexHandler)

	var statusExpect = http.StatusOK
	var contentTypeExpect = "text/html"

	var nameMetricGauge = "testMetricGauge"
	var nameMetricCounter = "testMetricCounter"

	rand.Seed(time.Now().UnixNano())
	var valueExpect = fmt.Sprint(rand.Int())
	router.ServeHTTP(responseRecorderPost, httptest.NewRequest("POST", "/update/gauge/"+nameMetricGauge+"/"+fmt.Sprint(valueExpect), nil))
	router.ServeHTTP(responseRecorderPost, httptest.NewRequest("POST", "/update/counter/"+nameMetricCounter+"/"+fmt.Sprint(valueExpect), nil))

	router.ServeHTTP(responseRecorder, httptest.NewRequest("GET", "/", nil))
	statusGet := responseRecorder.Code
	contentTypeGet := responseRecorder.Header().Get("Content-Type")
	body := responseRecorder.Body.String()

	assert.Equal(t, statusExpect, statusGet, fmt.Sprintf("Incorrect status code. Expect %d, got %d", statusExpect, statusGet))
	assert.Equal(t, contentTypeExpect, contentTypeGet, fmt.Sprintf("Incorrect Content-Type. Expect %s, got %s", contentTypeExpect, contentTypeGet))
	assert.True(t, strings.Contains(body, nameMetricGauge))
	assert.True(t, strings.Contains(body, nameMetricCounter))
}
