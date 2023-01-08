package storage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dataTEST = NewMemStorage()

func TestPutMetricsGauge(t *testing.T) {

	tests := []struct {
		name  string
		key   string
		value float64
	}{
		{name: "Test1", key: "key1", value: 1.11},
		{name: "Test2", key: "key2", value: 2.22},
		{name: "Test3", key: "key3", value: 3.33},
	}

	for _, tt := range tests {
		dataTEST.PutMetricsGauge(tt.key, tt.value)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			val, ok := dataTEST.GetMetricsGauge(tt.key)

			assert.Equal(t, ok, true,
				fmt.Sprintf("Element %v not found", val))
		})
	}
}

func TestPutMetricsCount(t *testing.T) {

	tests := []struct {
		name  string
		key   string
		value int64
	}{
		{name: "Test1", key: "key1", value: 1},
		{name: "Test2", key: "key2", value: 2},
		{name: "Test3", key: "key3", value: 3},
	}

	for _, tt := range tests {
		dataTEST.PutMetricsCount(tt.key, tt.value)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			val, ok := dataTEST.GetMetricsCount(tt.key)

			assert.Equal(t, ok, true,
				fmt.Sprintf("Element %v not found", val))
		})
	}
}
