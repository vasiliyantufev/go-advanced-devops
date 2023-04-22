// model - jsonmetric
package models

import (
	"fmt"
	"strconv"
)

//type Batch struct {
//	Metrics []Metric
//}

type Metric struct {
	ID    string   `json:"id"`              // metric name
	MType string   `json:"type"`            // a parameter that takes the value gauge or counter
	Delta *int64   `json:"delta,omitempty"` // metric value in case of passing counter
	Value *float64 `json:"value,omitempty"` // metric value in case of passing gauge
	Hash  string   `json:"hash"`            // hash value
}

func (J Metric) String() string {
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

//func (batches *Batch) AddItem(metric Metric) []Metric {
//	batches.Metrics = append(batches.Metrics, metric)
//	return batches.Metrics
//}
