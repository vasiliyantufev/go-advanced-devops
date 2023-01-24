package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/flags"
	"time"
)

var cfgSrv configServer

type configServer struct {
	//Address       string        `env:"ADDRESS" envDefault:"localhost:8080"`
	Address string `env:"ADDRESS"`
	//StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"300s"`
	StoreInterval time.Duration `env:"STORE_INTERVAL"`
	//StoreFile string `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	StoreFile string `env:"STORE_FILE"`
	Restore   bool   `env:"RESTORE" envDefault:"true"`
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

	log.Info(cfgSrv)
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