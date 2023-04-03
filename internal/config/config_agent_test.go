package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfigAgent(t *testing.T) {

	//initialArgs := os.Args
	//defer func() { os.Args = initialArgs }()
	//os.Args = []string{"", "a=localhost:8080", "r=10s", "p=2s", "k=secretKey", "l=2"}
	//
	//flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	//flagSet.Parse(os.Args[1:])

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
