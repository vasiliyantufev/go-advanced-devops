package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func TestIndexHandler(t *testing.T) {
//
//	r := chi.NewRouter()
//	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
//		defer r.Body.Close()
//		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
//		w.WriteHeader(http.StatusOK)
//	})
//
//	ts := httptest.NewServer(r)
//	defer ts.Close()
//
//	res, _ := TestRequest(t, ts, "GET", "/", nil)
//	defer res.Body.Close()
//
//	assert.Equal(t, res.StatusCode, http.StatusOK)
//}

func TestHandler_IndexHandler(t *testing.T) {
	testTable := []struct {
		name   string
		server *httptest.Server
	}{
		{
			name: "test index handler",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				w.WriteHeader(http.StatusOK)
			})),
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.server.Close()
			resp, err := MakeHTTPCall(tc.server.URL)

			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, resp.StatusCode, http.StatusOK)
			assert.Equal(t, resp.Header.Get("Content-Type"), "text/html")
		})
	}
}
