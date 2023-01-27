package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCounterHandler(t *testing.T) {

	wg := httptest.NewRecorder()
	wp := httptest.NewRecorder()
	wg2 := httptest.NewRecorder()
	wp2 := httptest.NewRecorder()

	rtr := chi.NewRouter()
	rtr.Get("/value/{type}/{name}", GetMetricsHandler)
	rtr.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", MetricsHandler)
	})

	var val1 int64 = 22
	rtr.ServeHTTP(wp, httptest.NewRequest("POST", "/update/counter/testSetGet33/"+fmt.Sprint(val1), nil))
	rtr.ServeHTTP(wg, httptest.NewRequest("GET", "/value/counter/testSetGet33", nil))
	bodyGet := wg.Body.String()

	k, _ := strconv.ParseInt(string(bodyGet), 10, 64)
	assert.Equal(t, val1, k,
		fmt.Sprintf("Incorrect body. Expect %s, got %s", fmt.Sprint(val1), bodyGet))

	var val2 int64 = 33
	rtr.ServeHTTP(wp2, httptest.NewRequest("POST", "/update/counter/testSetGet33/"+fmt.Sprint(val2), nil))
	rtr.ServeHTTP(wg2, httptest.NewRequest("GET", "/value/counter/testSetGet33", nil))
	bodyGet = wg2.Body.String()

	k, _ = strconv.ParseInt(string(bodyGet), 10, 64)
	assert.Equal(t, val1+val2, k,
		fmt.Sprintf("Incorrect body. Expect %s, got %s", fmt.Sprint(val1+val2), bodyGet))
}

func TestGaugeHandler(t *testing.T) {

	wg := httptest.NewRecorder()
	wp := httptest.NewRecorder()

	rtr := chi.NewRouter()
	rtr.Get("/value/{type}/{name}", GetMetricsHandler)
	rtr.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", MetricsHandler)
	})

	var val int64 = 22
	rtr.ServeHTTP(wp, httptest.NewRequest("POST", "/update/gauge/testSetGet22/"+fmt.Sprint(val), nil))
	rtr.ServeHTTP(wg, httptest.NewRequest("GET", "/value/gauge/testSetGet22", nil))
	bodyGet := wg.Body.String()

	k, _ := strconv.ParseInt(string(bodyGet), 10, 64)
	assert.Equal(t, val, k,
		fmt.Sprintf("Incorrect body. Expect %s, got %s", fmt.Sprint(val), bodyGet))

}
