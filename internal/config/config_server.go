package config

import (
	"time"

	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/flags"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type ConfigServicerServer interface {
	GetConfigAddressServer() string
	GetConfigStoreIntervalServer() time.Duration
	GetConfigStoreFileServer() string
	GetConfigDebugLevelServer() string
	GetConfigKeyServer() string
	GetConfigRestoreServer() bool
	GetConfigDBServer() string
}

type ConfigServer struct {
	//Address       string        `env:"ADDRESS" envDefault:"localhost:8080"`
	Address string `env:"ADDRESS"`
	//StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"300s"`
	StoreInterval time.Duration `env:"STORE_INTERVAL"`
	//StoreFile string `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	DebugLevel log.Level `env:"DEBUG_LEVEL" envDefault:"debug"`
	StoreFile  string    `env:"STORE_FILE"`
	Restore    bool      `env:"RESTORE" envDefault:"true"`
	Key        string    `env:"KEY" envDefault:""`
	//DNS        string    `env:"DATABASE_DSN"`
	//DNS string `env:"DATABASE_DSN" envDefault:"host=localhost port=5432 user=postgres password=myPassword dbname=praktikum sslmode=disable"`
	DNS string `env:"DATABASE_DSN" envDefault:"host=localhost port=5432 user=postgres password=postgres dbname=praktikum sslmode=disable"`
}

func NewConfigServer() *ConfigServer {

	var cfgSrv ConfigServer

	err := env.Parse(&cfgSrv)
	if err != nil {
		log.Fatal(err)
	}

	flagsSrv := flags.NewFlagsServer()

	if cfgSrv.Address == "" {
		cfgSrv.Address = flagsSrv.Address
	}
	if cfgSrv.StoreInterval.String() == "0s" {
		cfgSrv.StoreInterval, _ = time.ParseDuration(flagsSrv.StoreInterval)
	}
	if cfgSrv.StoreFile == "" {
		cfgSrv.StoreFile = flagsSrv.StoreFile
	}
	if cfgSrv.Key == "" {
		cfgSrv.Key = flagsSrv.Key
	}
	if cfgSrv.DNS == "" {
		cfgSrv.DNS = flagsSrv.DNS
		if cfgSrv.DNS == "" {
			cfgSrv.DNS = "host=localhost port=5432 user=postgres password=postgres dbname=praktikum sslmode=disable"
		}
	}

	log.Info(cfgSrv)

	return &cfgSrv
}

func (cfg ConfigServer) GetConfigAddressServer() string {
	return cfg.Address
}

func (cfg ConfigServer) GetConfigStoreIntervalServer() time.Duration {
	return cfg.StoreInterval
}

func (cfg ConfigServer) GetConfigStoreFileServer() string {
	return cfg.StoreFile
}

func (cfg ConfigServer) GetConfigDebugLevelServer() log.Level {
	return cfg.DebugLevel
}

func (cfg ConfigServer) GetConfigKeyServer() string {
	return cfg.Key
}

func (cfg ConfigServer) GetConfigRestoreServer() bool {
	return cfg.Restore
}

func (cfg ConfigServer) GetConfigDBServer() string {
	return cfg.DNS
}
