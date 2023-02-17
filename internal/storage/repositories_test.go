package storage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dataTEST = NewMemStorage()

func TestPutMetricsGauge(t *testing.T) {

	tests := []struct {
		name  string
		typeM string
		value float64
		hash  string
	}{
		{name: "Test1", typeM: "gauge", value: 1.11, hash: "fb28d951abb877a1877ca104c1c215429b411f99a27870be00838c2a735900f3"}, //
		{name: "Test2", typeM: "gauge", value: 2.22, hash: "6e5225371040d53dd7f5c580437e9c444d2f5228bd716d20de142e7ad13bdc4f"}, //
		{name: "Test3", typeM: "gauge", value: 3.33, hash: "c8e9edfc49ac27642f60961bd6c64b3d41dc792de70d0f3097952c85d34f03fb"}, //
	}

	for _, tt := range tests {
		dataTEST.PutMetricsGauge(tt.name, tt.value, tt.hash)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			val, _, ok := dataTEST.GetMetricsGauge(tt.name)

			assert.Equal(t, ok, true,
				fmt.Sprintf("Element %v not found", val))
		})
	}
}

func TestPutMetricsCount(t *testing.T) {

	tests := []struct {
		name  string
		typeM string
		value int64
		hash  string
	}{
		{name: "Test1", typeM: "key1", value: 1, hash: "4d0426203e8c75a0146461fed132b9601503f0d23493ad2e539b068908cd46a6"}, //
		{name: "Test2", typeM: "key2", value: 2, hash: "2f5f81d87530fdba0ad57d57d314bd9980afc35ff2cb5fa09fd0143ad990636c"}, //
		{name: "Test3", typeM: "key3", value: 3, hash: "f2d29f26922610a890f5aef24e341926ec0089e8383c7afa34ae6cc3c03c0294"}, //
	}

	for _, tt := range tests {
		dataTEST.PutMetricsCount(tt.name, tt.value, tt.hash)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			val, _, ok := dataTEST.GetMetricsCount(tt.name)

			assert.Equal(t, ok, true,
				fmt.Sprintf("Element %v not found", val))
		})
	}
}
