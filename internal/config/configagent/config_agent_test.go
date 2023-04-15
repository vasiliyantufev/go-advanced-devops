package configagent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigAgent(t *testing.T) {
	configAgentInstance := NewConfigAgent()
	configAgent := ConfigAgent{
		Address:        configAgentInstance.Address,
		ReportInterval: configAgentInstance.ReportInterval,
		PollInterval:   configAgentInstance.PollInterval,
		Key:            configAgentInstance.Key,
		RateLimit:      configAgentInstance.RateLimit,
	}
	assert.Equal(t, configAgentInstance, &configAgent)
}
