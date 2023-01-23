package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/flags"
	"time"
)

var cfgAgt configAgent

type configAgent struct {
	//Address        string        `env:"ADDRESS" envDefault:"localhost:8080"`
	Address string `env:"ADDRESS"`
	//ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL"`
	//PollInterval time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
	PollInterval time.Duration `env:"POLL_INTERVAL"`
}

func SetConfigAgent() {

	err := env.Parse(&cfgAgt)
	if err != nil {
		log.Fatal(err)
	}

	if cfgAgt.Address == "" {
		cfgAgt.Address = flags.GetFlagAddressAgent()
	}
	if cfgAgt.ReportInterval.String() == "0s" {
		cfgAgt.ReportInterval, _ = time.ParseDuration(flags.GetFlagReportIntervalAgent())
	}
	if cfgAgt.PollInterval.String() == "0s" {
		cfgAgt.PollInterval, _ = time.ParseDuration(flags.GetFlagPollIntervalAgent())
	}
}

func GetConfigAddressAgent() string {
	return cfgAgt.Address
}

func GetConfigReportIntervalAgent() time.Duration {
	return cfgAgt.ReportInterval
}

func GetConfigPollIntervalAgent() time.Duration {
	return cfgAgt.PollInterval
}
