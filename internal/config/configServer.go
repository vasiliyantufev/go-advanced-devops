package config

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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
	DebugLevel  log.Level `env:"DEBUG_LEVEL" envDefault:"debug"`
	StoreFile   string    `env:"STORE_FILE"`
	Restore     bool      `env:"RESTORE" envDefault:"true"`
	Key         string    `env:"KEY" envDefault:""`
	DatabaseDns string    `env:"DATABASE_DSN"`
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
	if cfgSrv.Key == "" {
		cfgSrv.Key = flags.GetFlagKeyServer()
	}
	if cfgSrv.DatabaseDns == "" {
		cfgSrv.DatabaseDns = flags.GetFlagDataBaseServer()
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

func GetConfigDebugLevelServer() log.Level {
	return cfgSrv.DebugLevel
}

func GetConfigKeyServer() string {
	return cfgSrv.Key
}

func GetConfigRestoreServer() bool {
	return cfgSrv.Restore
}

func GetConfigDBServer() string {
	return cfgSrv.DatabaseDns
}

func GetHashServer(mid string, mtype string, delta int64, value float64) string {

	var data string
	switch mtype {
	case "counter":
		data = fmt.Sprintf("%s:%s:%d", mid, mtype, delta)
	case "gauge":
		data = fmt.Sprintf("%s:%s:%f", mid, mtype, value)
	}

	h := hmac.New(sha256.New, []byte(cfgSrv.Key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
