package handlers

import (
	"fmt"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	//"github.com/vasiliyantufev/go-advanced-devops/internal/api/server"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"
)

/*
func TestHandler_GetValueMetricJSONHandler(t *testing.T) {

		metric := &Response{
			ID:    "alloc",
			MType: "gauge",
			Value: converter.Uint64ToFloat64Pointer(48),
			Hash:  "9ef052d6a06f3fd3f67e264174578411e823df10d7f904b2d723e73dbe0ea632",
		}

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

	wg := httptest.NewRecorder()
	wp := httptest.NewRecorder()

	memStorage := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer("secret")

	configServer := configserver.ConfigServer{
		Address:         "localhost:8080",
		AddressPProfile: "localhost:8088",
		Restore:         true,
		StoreInterval:   300 * time.Second,
		DebugLevel:      logrus.DebugLevel,
		StoreFile:       "/tmp/devops-metrics-db.json",
		Key:             "",
		DSN:             "",
		RootPath:        "file://./migrations",
	}

	srv := NewHandler(memStorage, nil, &configServer, nil, hashServer)

	rtr := chi.NewRouter()
	rtr.Get("/value/{type}/{name}", srv.GetMetricURLParamsHandler)
	rtr.Post("/update/{type}/{name}/{value}", srv.CreateMetricURLParamsHandler)

	var val1 int64 = 22
	rtr.ServeHTTP(wp, httptest.NewRequest("POST", "/update/counter/testSetGet33/"+fmt.Sprint(val1), nil))
	rtr.ServeHTTP(wg, httptest.NewRequest("GET", "/value/counter/testSetGet33", nil))
	bodyGet := wg.Body.String()

	fmt.Print(bodyGet)

	k, _ := strconv.ParseInt(string(bodyGet), 10, 64)
	assert.Equal(t, val1, k,
		fmt.Sprintf("Incorrect body. Expect %s, got %s", fmt.Sprint(val1), bodyGet))

}

//
//func TestHandler_GetValueMetricJSONHandler3(t *testing.T) {
//	type fields struct {
//		memStorage  *memstorage.MemStorage
//		fileStorage *filestorage.FileStorage
//		config      *configserver.ConfigServer
//		database    *database.DB
//		hashServer  *hashservicer.HashServer
//	}
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := Handler{
//				memStorage:  tt.fields.memStorage,
//				fileStorage: tt.fields.fileStorage,
//				config:      tt.fields.config,
//				database:    tt.fields.database,
//				hashServer:  tt.fields.hashServer,
//			}
//			s.CreateMetricURLParamsHandler(tt.args.w, tt.args.r)
//		})
//	}
//}
