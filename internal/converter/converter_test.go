package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverter(t *testing.T) {
	var val64ui uint64 = 1
	var val32ui uint32 = 1
	var convertVal = float64(val64ui)
	assert.Equal(t, &convertVal, Uint64ToFloat64Pointer(val64ui))
	assert.Equal(t, &convertVal, Uint32ToFloat64Pointer(val32ui))

	var val64uf = 1.1
	var val32i int64 = 1
	assert.Equal(t, &val64uf, Float64ToFloat64Pointer(val64uf))
	assert.Equal(t, &val32i, Int64ToInt64Pointer(val32i))
}
