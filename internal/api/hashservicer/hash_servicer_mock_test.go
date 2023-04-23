package hashservicer

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_hs "github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer/mock"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/helpers"
	"github.com/vasiliyantufev/go-advanced-devops/internal/converter"
	"github.com/vasiliyantufev/go-advanced-devops/internal/models"
)

func TestIsEnabledMock(t *testing.T) {
	ctl := gomock.NewController(t)
	ctl.Finish()

	hs := mock_hs.NewMockHashServices(ctl)
	hs.EXPECT().IsEnabled().Return(true).Times(1)
	assert.Equal(t, true, hs.IsEnabled())
}

func TestGenerateHashMock(t *testing.T) {
	ctl := gomock.NewController(t)
	ctl.Finish()

	hash := helpers.RandStr(10)

	clientMetric := models.Metric{
		ID:    "Alloc",
		MType: "gauge",
		Delta: nil,
		Value: converter.Uint64ToFloat64Pointer(1),
	}
	hs := mock_hs.NewMockHashServices(ctl)
	hs.EXPECT().GenerateHash(clientMetric).Return(hash).Times(1)
	assert.True(t, len(hs.GenerateHash(clientMetric)) > 0)
}

func TestValidHashMock(t *testing.T) {
	ctl := gomock.NewController(t)
	ctl.Finish()

	clientMetric := models.Metric{
		ID:    "Alloc",
		MType: "gauge",
		Delta: nil,
		Value: converter.Uint64ToFloat64Pointer(1),
	}
	hashServer := mock_hs.NewMockHashServices(ctl)
	hashServer.EXPECT().ValidHashServer(clientMetric).Return(true).Times(1)
	assert.True(t, hashServer.ValidHashServer(clientMetric))
}
