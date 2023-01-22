package storage

import (
	"time"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Address        string        `env:"ADDRESS" envDefault:"localhost:8080"`
	StoreInterval  time.Duration `env:"STORE_INTERVAL" envDefault:"300s"`
	StoreFile      string        `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	Restore        bool          `env:"RESTORE" envDefault:"true"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
	PollInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
}

func GetConfigEnv() Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Error(err)
	}
	return cfg
}

func GetConfigServer(flags FlagsServer) Config {

	var cfg Config
	cfg = GetConfigEnv()

	if cfg.Address == "" {
		cfg.Address = flags.a
	}
	if cfg.StoreFile == "" {
		cfg.StoreFile = flags.f
	}
	if cfg.StoreInterval.String() == "" {
		cfg.StoreInterval, _ = time.ParseDuration(flags.i)
	}
	return cfg
}

func GetConfigAgent(flags FlagsAgent) Config {

	var cfg Config
	cfg = GetConfigEnv()

	if cfg.Address == "" {
		cfg.Address = flags.a
	}
	if cfg.ReportInterval.String() == "" {
		cfg.ReportInterval, _ = time.ParseDuration(flags.r)
	}
	if cfg.PollInterval.String() == "" {
		cfg.PollInterval, _ = time.ParseDuration(flags.p)
	}
	return cfg
}
