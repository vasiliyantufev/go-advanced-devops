package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_GetMetricURLParamsHandler(t *testing.T) {

	value := "42"
	var contType = "text/plain; charset=utf-8"

	testTable := []struct {
		name             string
		server           *httptest.Server
		expectedResponse string
		expectedErr      error
	}{
		{
			name: "test get value metric url params handler",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", contType)
				w.Write([]byte(value))
			})),
			expectedResponse: "42",
			expectedErr:      nil,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.server.Close()
			resp, respBody, err := MakeHTTPWithBodyValueJSONCall(tc.server.URL)
			if err != nil {
				t.Error(err)
			}
			defer resp.Body.Close()

			assert.Equal(t, resp.StatusCode, http.StatusOK)
			assert.Equal(t, resp.Header.Get("Content-Type"), contType)
			assert.Equal(t, string(respBody), tc.expectedResponse)
		})
	}
}
