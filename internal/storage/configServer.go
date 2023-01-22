package storage

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
	"time"
)

var CfgSrv Config

type ConfigSrv struct {
	Address       string        `env:"ADDRESS" envDefault:"localhost:8080"`
	StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"300s"`
	StoreFile     string        `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	Restore       bool          `env:"RESTORE" envDefault:"true"`
}

func SetConfigServer() Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Error(err)
	}
	return cfg
}
