// storage - jsonmetrics
package storage

import (
	"fmt"
	"strconv"
)

type JSONMetrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
	Hash  string   `json:"hash"`            // значение хеш-функции
}

func (J JSONMetrics) String() string {
	switch J.MType {
	case "gauge":
		var value string
		if J.Value != nil {
			value = strconv.FormatFloat(*J.Value, 'f', 3, 64)
		} else {
			value = "empty"
		}
		return fmt.Sprintf("Metric {ID: %s Type: %s Value: %s Hash: %s}", J.ID, J.MType, value, J.Hash)
	case "counter":
		var delta string
		if J.Delta != nil {
			delta = strconv.FormatInt(*J.Delta, 10)
		} else {
			delta = "empty"
		}
		return fmt.Sprintf("Metric {ID: %s Type: %s Delta: %s Hash: %s}", J.ID, J.MType, delta, J.Hash)
	default:
		return fmt.Sprintf("Metric {ID: %s Type: unknown}", J.ID)
	}
}
