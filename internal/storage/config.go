package storage

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
	"time"
)

var cfg Config

type Config struct {
	Address       string        `env:"ADDRESS" envDefault:"localhost:8080"`
	StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"300s"`
	//StoreInterval time.Duration `env:"STORE_INTERVAL"`
	StoreFile string `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	//StoreFile      string        `env:"STORE_FILE"`
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

func GetConfigServer() Config {

	var cfg Config
	cfg = GetConfigEnv()

	if cfg.Address == "" {
		cfg.Address = FgSrv.a
	}
	if cfg.StoreFile == "" {
		cfg.StoreFile = FgSrv.f
	}
	if cfg.StoreInterval.String() == "0s" {
		cfg.StoreInterval, _ = time.ParseDuration(FgSrv.i)
	}
	//log.Fatal(cfg)
	return cfg
}

func GetConfigAgent() Config {

	var cfg Config
	cfg = GetConfigEnv()

	if cfg.Address == "" {
		cfg.Address = FgAgt.a
	}
	if cfg.ReportInterval.String() == "0s" {
		cfg.ReportInterval, _ = time.ParseDuration(FgAgt.r)
	}
	if cfg.PollInterval.String() == "0s" {
		cfg.PollInterval, _ = time.ParseDuration(FgAgt.p)
	}
	return cfg
}
