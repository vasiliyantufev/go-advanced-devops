package configserver

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestConfigServer(t *testing.T) {
	configServerInstance := NewConfigServer()
	configServer := ConfigServer{
		Address:         "localhost:8080",
		AddressPProfile: "localhost:8088",
		Restore:         true,
		StoreInterval:   300 * time.Second,
		DebugLevel:      log.DebugLevel,
		StoreFile:       "/tmp/devops-metrics-db.json",
		Key:             "",
		DSN:             "",
		RootPath:        "file://./migrations",
	}
	assert.Equal(t, configServerInstance, &configServer)
}
