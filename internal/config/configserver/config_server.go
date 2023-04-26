// Package configserver - setting flags for the server
package configserver

import (
	"flag"
	"time"

	"github.com/caarlos0/env/v6"

	log "github.com/sirupsen/logrus"
)

type ConfigServer struct {
	Address         string        `env:"ADDRESS" json:"address"`
	AddressPProfile string        `env:"ADDRESS_PPROFILE" json:"address_pprofile"`
	StoreInterval   time.Duration `env:"STORE_INTERVAL" json:"store_interval"`
	DebugLevel      log.Level     `env:"DEBUG_LEVEL" envDefault:"debug" json:"debug_level"`
	StoreFile       string        `env:"STORE_FILE" json:"store_file"`
	Restore         bool          `env:"RESTORE" json:"restore"`
	Key             string        `env:"KEY" json:"key"`
	DSN             string        `env:"DATABASE_DSN" json:"dsn"`
	CryptoKey       string        `env:"CRYPTO_KEY" json:"crypto_key"`
	Certificate     string        `env:"CERTIFICATE" json:"certificate"`
	ConfigFile      string        `env:"CONFIG"`
	//DSN string `env:"DATABASE_DSN" envDefault:"host=localhost port=5432 user=postgres password=postgres dbname=praktikum sslmode=disable"`
	//DSN      string `env:"DATABASE_DSN" envDefault:"host=localhost port=5432 user=postgres password=myPassword dbname=praktikum sslmode=disable"`
	MigrationsPath string `env:"ROOT_PATH" envDefault:"file://./migrations"`
	TemplatePath   string `env:"TEMPLATE_PATH" envDefault:"./web/templates/index.html"`
}

// NewConfigServer - creates a new instance with the configuration for the server
func NewConfigServer() *ConfigServer {
	cfgSrv := ConfigServer{}
	// setting flags for the server
	flag.StringVar(&cfgSrv.Address, "a", "localhost:8080", "Server address")
	flag.StringVar(&cfgSrv.AddressPProfile, "ap", "localhost:8088", "Profile address")
	flag.BoolVar(&cfgSrv.Restore, "r", true, "Boolean value specifying whether or not to load the initial values from the specified file when the server starts")
	flag.DurationVar(&cfgSrv.StoreInterval, "i", 300*time.Second, "Time interval in seconds after which the current server readings are flushed to disk")
	flag.StringVar(&cfgSrv.StoreFile, "f", "/tmp/devops-metrics-db.json", "The file where the values are stored")
	flag.StringVar(&cfgSrv.Key, "k", "", "Key to generate hash")
	flag.StringVar(&cfgSrv.CryptoKey, "crypto-key", "", "Path to crypto key")
	flag.StringVar(&cfgSrv.Certificate, "certificate", "", "Path to certificate")
	//flag.StringVar(&cfgSrv.CryptoKey, "crypto-key", "./certificates/server.key", "Path to crypto key")
	//flag.StringVar(&cfgSrv.Certificate, "certificate", "./certificates/server.crt", "Path to certificate")
	flag.StringVar(&cfgSrv.DSN, "d", "", "Database configuration")
	flag.StringVar(&cfgSrv.ConfigFile, "f", "", "Path to config file")
	flag.Parse()

	err := env.Parse(&cfgSrv)
	if err != nil {
		log.Fatal(err)
	}

	log.Debug(cfgSrv)
	return &cfgSrv
}

func parseFileJSON(path string) {

}
