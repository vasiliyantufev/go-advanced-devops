package config

import (
	"time"

	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/flags"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type ConfigServicerAgent interface {
	GetConfigAddressAgent() string
	GetConfigReportIntervalAgent() time.Duration
	GetConfigPollIntervalAgent() time.Duration
	GetConfigKeyAgent() string
}

type ConfigAgent struct {
	//Address        string        `env:"ADDRESS" envDefault:"localhost:8080"`
	Address string `env:"ADDRESS"`
	//ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL"`
	//PollInterval time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
	PollInterval time.Duration `env:"POLL_INTERVAL"`
	//PollInterval time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
	Key string `env:"KEY" envDefault:""`
}

func NewConfigAgent() *ConfigAgent {

	var cfgAgt ConfigAgent

	err := env.Parse(&cfgAgt)
	if err != nil {
		log.Fatal(err)
	}

	flags := flags.NewFlagsAgent()

	if cfgAgt.Address == "" {
		cfgAgt.Address = flags.Address
	}
	if cfgAgt.ReportInterval.String() == "0s" {
		cfgAgt.ReportInterval, _ = time.ParseDuration(flags.ReportInterval)
	}
	if cfgAgt.PollInterval.String() == "0s" {
		cfgAgt.PollInterval, _ = time.ParseDuration(flags.PollInterval)
	}
	if cfgAgt.Key == "" {
		cfgAgt.Key = flags.Key
	}

	log.Info(cfgAgt)

	return &cfgAgt
}

func (cfg ConfigAgent) GetConfigAddressAgent() string {
	return cfg.Address
}

func (cfg ConfigAgent) GetConfigReportIntervalAgent() time.Duration {
	return cfg.ReportInterval
}

func (cfg ConfigAgent) GetConfigPollIntervalAgent() time.Duration {
	return cfg.PollInterval
}

func (cfg ConfigAgent) GetConfigKeyAgent() string {
	return cfg.Key
}
