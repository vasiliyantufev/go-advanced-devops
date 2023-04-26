// Package configagent - setting flags for the agent
package configagent

import (
	"flag"
	"time"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type ConfigAgent struct {
	Address        string        `env:"ADDRESS" json:"address"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" json:"report_interval"`
	PollInterval   time.Duration `env:"POLL_INTERVAL" json:"poll_interval"`
	Key            string        `env:"KEY" json:"key"`
	RateLimit      int           `env:"RATE_LIMIT" json:"rate_limit"`
	CryptoKey      string        `env:"CRYPTO_KEY" json:"crypto_key"`
	ConfigFile     string        `env:"CONFIG"`
}

// NewConfigAgent - creates a new instance with the configuration for the agent
func NewConfigAgent() *ConfigAgent {
	var cfgAgt ConfigAgent
	// setting flags for the agent
	flag.StringVar(&cfgAgt.Address, "a", "localhost:8080", "Server address")
	flag.DurationVar(&cfgAgt.ReportInterval, "r", 10*time.Second, "Time interval in seconds after which the current readings are sent to the server")
	flag.DurationVar(&cfgAgt.PollInterval, "p", 2*time.Second, "Time interval in seconds after which the current metric readings are updated on the client")
	flag.StringVar(&cfgAgt.Key, "k", "", "Key to generate hash")
	flag.IntVar(&cfgAgt.RateLimit, "l", 2, "Number of concurrent outgoing requests to the server")
	flag.StringVar(&cfgAgt.CryptoKey, "crypto-key", "", "Path to crypto key")
	//flag.StringVar(&cfgAgt.CryptoKey, "crypto-key", "./certificates/server.crt", "Path to crypto key")
	flag.Parse()

	err := env.Parse(&cfgAgt)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(cfgAgt)
	return &cfgAgt
}
