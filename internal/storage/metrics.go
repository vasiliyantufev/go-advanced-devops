package storage

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func ValidURLParamMetrics(typeMetrics, nameMetrics, valueMetrics string) error {

	if typeMetrics == "" {
		return fmt.Errorf("The query parameter type is missing")
	}
	if typeMetrics != "gauge" && typeMetrics != "counter" {
		return fmt.Errorf("The type incorrect " + typeMetrics)
	}
	if nameMetrics == "" {
		return fmt.Errorf("The query parameter name is missing")
	}
	if valueMetrics == "" {
		log.Error("The query parameter value is missing")
	}
	return nil
}

func ValidURLParamGetMetrics(typeMetrics, nameMetrics string) error {

	if typeMetrics == "" {
		return fmt.Errorf("The query parameter type is missing")
	}
	if typeMetrics != "gauge" && typeMetrics != "counter" {
		return fmt.Errorf("The type incorrect " + typeMetrics)
	}
	if nameMetrics == "" {
		return fmt.Errorf("The query parameter name is missing")
	}
	return nil
}
