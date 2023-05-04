// Package configserver - setting flags for the server
package configserver

import (
	"encoding/json"
	"flag"
	"io"
	"os"
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
	//DSN string `env:"DATABASE_DSN" envDefault:"host=localhost port=5432 user=postgres password=postgres dbname=praktikum sslmode=disable"`
	//DSN      string `env:"DATABASE_DSN" envDefault:"host=localhost port=5432 user=postgres password=myPassword dbname=praktikum sslmode=disable"`
	MigrationsPath string `env:"ROOT_PATH" envDefault:"file://./migrations"`
	TemplatePath   string `env:"TEMPLATE_PATH" envDefault:"./web/templates/index.html"`
	configFile     string `env:"CONFIG"`
	TrustedSubnet  string `env:"TRUSTED_SUBNET"`
}

// NewConfigServer - creates a new instance with the configuration for the server
func NewConfigServer() *ConfigServer {
	// Set default values
	configServer := ConfigServer{
		Address:         "localhost:8080",
		AddressPProfile: "localhost:8088",
		Restore:         true,
		StoreInterval:   300 * time.Second,
		StoreFile:       "/tmp/devops-metrics-db.json",
	}

	// setting flags for the server
	flag.StringVar(&configServer.Address, "a", configServer.Address, "Server address")
	flag.BoolVar(&configServer.Restore, "r", configServer.Restore, "Boolean value specifying whether or not to load the initial values from the specified file when the server starts")
	flag.DurationVar(&configServer.StoreInterval, "i", configServer.StoreInterval, "Time interval in seconds after which the current server readings are flushed to disk")
	flag.StringVar(&configServer.StoreFile, "f", configServer.StoreFile, "The file where the values are stored")
	flag.StringVar(&configServer.Key, "k", configServer.Key, "Key to generate hash")
	flag.StringVar(&configServer.CryptoKey, "crypto-key", configServer.CryptoKey, "Path to crypto key")
	flag.StringVar(&configServer.Certificate, "certificate", configServer.Certificate, "Path to certificate")
	//flag.StringVar(&cfgSrv.CryptoKey, "crypto-key", "./certificates/server.key", "Path to crypto key")
	//flag.StringVar(&cfgSrv.Certificate, "certificate", "./certificates/server.crt", "Path to certificate")
	flag.StringVar(&configServer.DSN, "d", configServer.DSN, "Database configuration")
	flag.StringVar(&configServer.configFile, "c", configServer.configFile, "Path to config file")
	flag.StringVar(&configServer.TrustedSubnet, "t", configServer.TrustedSubnet, "CIDR")
	flag.Parse()

	err := env.Parse(&configServer)
	if err != nil {
		log.Fatal(err)
	}

	if configServer.configFile != "" {
		fileConfig := ConfigServer{}
		fileConfig, err = parseFileJSON(configServer.configFile)
		if err != nil {
			log.Fatal(err)
		}
		mergeConfig(&configServer, &fileConfig)
	}

	log.SetLevel(configServer.DebugLevel)
	log.Debug(configServer)

	return &configServer
}

func parseFileJSON(path string) (ConfigServer, error) {
	fileConfig := ConfigServer{}
	jsonFile, err := os.Open(path)
	if err != nil {
		return fileConfig, err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return fileConfig, err
	}

	if err = json.Unmarshal(jsonData, &fileConfig); err != nil {
		return fileConfig, err
	}
	return fileConfig, nil
}

func mergeConfig(configServer *ConfigServer, fileConfig *ConfigServer) {
	if configServer.Address == "" && fileConfig.Address != "" {
		configServer.Address = fileConfig.Address
	}
	if configServer.Key == "" && fileConfig.Key != "" {
		configServer.Key = fileConfig.Key
	}
	if configServer.AddressPProfile == "" && fileConfig.AddressPProfile != "" {
		configServer.AddressPProfile = fileConfig.AddressPProfile
	}
	if configServer.DSN == "" && fileConfig.DSN != "" {
		configServer.DSN = fileConfig.DSN
	}
	if configServer.CryptoKey == "" && fileConfig.CryptoKey != "" {
		configServer.CryptoKey = fileConfig.CryptoKey
	}
	if configServer.Certificate == "" && fileConfig.Certificate != "" {
		configServer.Certificate = fileConfig.Certificate
	}
	if configServer.StoreFile == "" && fileConfig.StoreFile != "" {
		configServer.StoreFile = fileConfig.StoreFile
	}
	if configServer.StoreInterval == 0 && fileConfig.StoreInterval != 0 {
		configServer.StoreInterval = fileConfig.StoreInterval
	}
	if configServer.TrustedSubnet == "" && fileConfig.TrustedSubnet != "" {
		configServer.TrustedSubnet = fileConfig.TrustedSubnet
	}
}
