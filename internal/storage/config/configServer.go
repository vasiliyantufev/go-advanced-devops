package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/flags"
	"time"
)

var CfgSrv configServer

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

	err := env.Parse(&CfgSrv)
	if err != nil {
		log.Fatal(err)
	}

	log.Error(CfgSrv)

	if CfgSrv.Address == "" {
		CfgSrv.Address = flags.GetFlagAddressServer()
	}
	if CfgSrv.StoreInterval.String() == "0s" {
		CfgSrv.StoreInterval, _ = time.ParseDuration(flags.GetFlagStoreIntervalServer())
	}
	if CfgSrv.StoreFile == "" {
		CfgSrv.StoreFile = flags.GetFlagStoreFileServer()
	}
}

func GetConfigAddressServer() string {
	return CfgSrv.Address
}

func GetConfigStoreIntervalServer() time.Duration {
	return CfgSrv.StoreInterval
}

func GetConfigStoreFileServer() string {
	return CfgSrv.StoreFile
}

func GetConfigRestoreServer() bool {
	return CfgSrv.Restore
}
