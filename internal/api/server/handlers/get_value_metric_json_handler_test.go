package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vasiliyantufev/go-advanced-devops/internal/converter"
)

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
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.server.Close()
			resp, respBody, err := MakeHTTPWithBodyCall(tc.server.URL)
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
