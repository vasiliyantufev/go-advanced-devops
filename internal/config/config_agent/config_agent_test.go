package config_agent

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfigAgent(t *testing.T) {
	configAgentInstance := NewConfigAgent()
	configAgent := ConfigAgent{
		Address:        "localhost:8080",
		ReportInterval: 10 * time.Second,
		PollInterval:   2 * time.Second,
		Key:            "",
		RateLimit:      2,
	}
	assert.Equal(t, configAgentInstance, &configAgent)
}
