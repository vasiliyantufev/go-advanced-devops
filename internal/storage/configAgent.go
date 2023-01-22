package storage

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
	"time"
)

var CfgAgt Config

type ConfigAgt struct {
	Address        string        `env:"ADDRESS" envDefault:"localhost:8080"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
	PollInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
}

func SetConfigAgent() Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Error(err)
	}
	return cfg
}
