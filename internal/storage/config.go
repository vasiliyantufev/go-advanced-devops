package storage

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Address        string        `env:"ADDRESS" envDefault:"localhost:8080"`
	StoreInterval  time.Duration `env:"STORE_INTERVAL" envDefault:"4s"`
	StoreFile      string        `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	Restore        bool          `env:"RESTORE" envDefault:"true"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
	PollInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
}

func GetConfig() Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Error(err)
	}
	return cfg
}
