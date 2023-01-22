package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/flags"
	"time"
)

var cfgSrv configServer

type configServer struct {
	Address       string        `env:"ADDRESS" envDefault:"localhost:8080"`
	StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"300s"`
	StoreFile     string        `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	Restore       bool          `env:"RESTORE" envDefault:"true"`
}

func SetConfigServer() {

	err := env.Parse(&cfgSrv)
	if err != nil {
		log.Fatal(err)
	}

	if cfgSrv.Address == "" {
		cfgSrv.Address = flags.GetFlagAddressServer()
	}
	if cfgSrv.StoreInterval.String() == "0s" {
		cfgSrv.StoreInterval, _ = time.ParseDuration(flags.GetFlagStoreIntervalServer())
	}
	if cfgSrv.StoreFile == "" {
		cfgSrv.StoreFile = flags.GetFlagStoreFileServer()
	}
}

func GetConfigAddressServer() string {
	return cfgSrv.Address
}

func GetConfigStoreIntervalServer() time.Duration {
	return cfgSrv.StoreInterval
}

func GetConfigStoreFileServer() string {
	return cfgSrv.StoreFile
}

func GetConfigRestoreServer() bool {
	return cfgSrv.Restore
}
