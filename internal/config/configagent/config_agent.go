// Package configagent - setting flags for the agent
package configagent

import (
	"encoding/json"
	"flag"
	"io"
	"os"
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
	configFile     string        `env:"CONFIG"`
}

// NewConfigAgent - creates a new instance with the configuration for the agent
func NewConfigAgent() *ConfigAgent {
	// Set default values
	configAgent := ConfigAgent{
		Address:        "localhost:8080",
		ReportInterval: 10 * time.Second,
		PollInterval:   2 * time.Second,
		RateLimit:      2,
	}

	// setting flags for the agent
	flag.StringVar(&configAgent.Address, "a", configAgent.Address, "Server address")
	flag.DurationVar(&configAgent.ReportInterval, "r", configAgent.ReportInterval, "Time interval in seconds after which the current readings are sent to the server")
	flag.DurationVar(&configAgent.PollInterval, "p", configAgent.PollInterval, "Time interval in seconds after which the current metric readings are updated on the client")
	flag.StringVar(&configAgent.Key, "k", configAgent.Key, "Key to generate hash")
	flag.IntVar(&configAgent.RateLimit, "l", configAgent.RateLimit, "Number of concurrent outgoing requests to the server")
	flag.StringVar(&configAgent.CryptoKey, "crypto-key", configAgent.CryptoKey, "Path to crypto key")
	//flag.StringVar(&cfgAgt.CryptoKey, "crypto-key", "./certificates/server.crt", "Path to crypto key")
	flag.Parse()

	err := env.Parse(&configAgent)
	if err != nil {
		log.Fatal(err)
	}

	if configAgent.configFile != "" {
		fileConfig := ConfigAgent{}
		fileConfig, err = parseFileJSON(configAgent.configFile)
		if err != nil {
			log.Fatal(err)
		}
		mergeConfig(&configAgent, &fileConfig)
	}

	log.Debug(configAgent)
	return &configAgent
}

func parseFileJSON(path string) (ConfigAgent, error) {
	fileConfig := ConfigAgent{}
	jsonFile, err := os.Open(path)
	if err != nil {
		return fileConfig, err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return fileConfig, err
	}

	if err = json.Unmarshal(jsonData, &fileConfig); err != nil {
		return fileConfig, err
	}
	return fileConfig, nil
}

func mergeConfig(configAgent *ConfigAgent, fileConfig *ConfigAgent) {
	if configAgent.Address == "" && fileConfig.Address != "" {
		configAgent.Address = fileConfig.Address
	}
	if configAgent.ReportInterval == 0 && fileConfig.ReportInterval != 0 {
		configAgent.ReportInterval = fileConfig.ReportInterval
	}
	if configAgent.PollInterval == 0 && fileConfig.PollInterval != 0 {
		configAgent.PollInterval = fileConfig.PollInterval
	}
	if configAgent.RateLimit == 0 && fileConfig.RateLimit != 0 {
		configAgent.RateLimit = fileConfig.RateLimit
	}
	if configAgent.Key == "" && fileConfig.Key != "" {
		configAgent.Key = fileConfig.Key
	}
	if configAgent.CryptoKey == "" && fileConfig.CryptoKey != "" {
		configAgent.CryptoKey = fileConfig.CryptoKey
	}
}
