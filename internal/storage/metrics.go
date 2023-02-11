package storage

import (
	"fmt"
	"net/http"
)

func ValidURLParamMetrics(typeMetrics, nameMetrics, valueMetrics string) (int, error) {

	if typeMetrics == "" {
		return http.StatusBadRequest, fmt.Errorf("the query parameter type is missing")
	}
	if typeMetrics != "gauge" && typeMetrics != "counter" {
		return http.StatusNotImplemented, fmt.Errorf("the type incorrect " + typeMetrics)
	}
	if nameMetrics == "" {
		return http.StatusBadRequest, fmt.Errorf("the query parameter name is missing")
	}
	if valueMetrics == "" {
		return http.StatusBadRequest, fmt.Errorf("the query parameter name is missing")
	}
	return 0, nil
}

func ValidURLParamGetMetrics(typeMetrics, nameMetrics string) (int, error) {

	if typeMetrics == "" {
		return http.StatusBadRequest, fmt.Errorf("the query parameter type is missing")
	}
	if typeMetrics != "gauge" && typeMetrics != "counter" {
		return http.StatusNotImplemented, fmt.Errorf("the type incorrect " + typeMetrics)
	}
	if nameMetrics == "" {
		return http.StatusBadRequest, fmt.Errorf("the query parameter name is missing")
	}
	return 0, nil
}
