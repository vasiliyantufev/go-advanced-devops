package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	database "github.com/vasiliyantufev/go-advanced-devops/internal/db"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
)

/**/
func TestHandler_CreateMetricURLParamsHandler(t *testing.T) {

	//w := httptest.NewRecorder()
	//r := httptest.NewRequest("POST", "/update/{type}/{name}/{value}", nil)

	//h := Handler.CreateMetricURLParamsHandler
	//
	//t.Run("WithUUID", func(t *testing.T) {
	//
	//	//r := httptest.NewRequest(http.MethodGet, "/products/1", nil) // note that this URL is useless
	//	//r = mux.SetURLVars(r, map[string]string{"uuid": "1"})
	//	//w := httptest.NewRecorder()
	//
	//	//rctx := chi.NewRouteContext()
	//	//rctx.URLParams.Add("key", "value")
	//	//GetProduct(w, r)
	//
	//	rctx := chi.NewRouteContext()
	//	rctx.URLParams.Add("key", "value")
	//
	//
	//	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	//})

	testTable := []struct {
		name        string
		server      *httptest.Server
		expectedErr error
	}{
		{
			name: "test create metric url params handler",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})),
			//server:      httptest.NewServer(Handler.CreateMetricURLParamsHandler),
			expectedErr: nil,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.server.Close()
			resp, err := MakeHTTPCall(tc.server.URL)
			if err != tc.expectedErr {
				t.Error(err)
			}
			defer resp.Body.Close()

			assert.Equal(t, resp.StatusCode, http.StatusOK)
		})
	}

}

func TestHandler_CreateMetricURLParamsHandler1(t *testing.T) {

	type fields struct {
		mem        *storage.MemStorage
		config     *configserver.ConfigServer
		database   *database.DB
		hashServer *hashservicer.HashServer
	}

	//s := Handler{
	//	mem:        fields.mem,
	//	config:     tt.fields.config,
	//	database:   tt.fields.database,
	//	hashServer: tt.fields.hashServer,
	//}

	//type args struct {
	//	w http.ResponseWriter
	//	r *http.Request
	//}
	//tests := []struct {
	//	name   string
	//	fields fields
	//	args   args
	//}{
	//	// TODO: Add test cases.
	//
	//}
	//for _, tt := range tests { /**/
	//	t.Run(tt.name, func(t *testing.T) {
	//		s := Handler{
	//			mem:        tt.fields.mem,
	//			config:     tt.fields.config,
	//			database:   tt.fields.database,
	//			hashServer: tt.fields.hashServer,
	//		}
	//
	//		r := chi.NewRouter()
	//		//r.Get("/value/{type}/{name}", s.GetMetricURLParamsHandler)
	//		r.Post("/update/{type}/{name}/{value}", s.CreateMetricURLParamsHandler)
	//
	//		wp := httptest.NewRecorder()
	//
	//		var val1 int64 = 22
	//		r.ServeHTTP(wp, httptest.NewRequest("POST", "/update/counter/testSetGet33/"+fmt.Sprint(val1), nil))
	//
	//		logrus.Info(wp)
	//		logrus.Info(wp)
	//
	//		//ts := httptest.NewServer(r)
	//		//defer ts.Close()
	//	})
	//}
}

func TestHandler_CreateMetricURLParamsHandler2(t *testing.T) {
	type fields struct {
		mem        *storage.MemStorage
		config     *configserver.ConfigServer
		database   *database.DB
		hashServer *hashservicer.HashServer
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		//fields fields
		//args   args
		h Handler
	}{
		// TODO: Add test cases
		{
			name: "jjj",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//s := Handler{
			//	mem:        tt.fields.mem,
			//	config:     tt.fields.config,
			//	database:   tt.fields.database,
			//	hashServer: tt.fields.hashServer,
			//}
			//s.CreateMetricURLParamsHandler(tt.args.w, tt.args.r)
		})
	}
}
