package configserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigServer(t *testing.T) {
	configServerInstance := NewConfigServer()
	configServer := ConfigServer{
		Address:         configServerInstance.Address,
		AddressPProfile: configServerInstance.AddressPProfile,
		Restore:         configServerInstance.Restore,
		StoreInterval:   configServerInstance.StoreInterval,
		DebugLevel:      configServerInstance.DebugLevel,
		StoreFile:       configServerInstance.StoreFile,
		Key:             configServerInstance.Key,
		DSN:             configServerInstance.DSN,
		MigrationsPath:  configServerInstance.MigrationsPath,
		TemplatePath:    configServerInstance.TemplatePath,
		configFile:      configServerInstance.configFile,
	}

	assert.Equal(t, configServerInstance, &configServer)
}
