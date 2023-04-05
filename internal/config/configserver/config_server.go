// Package configserver - setting flags for the server
package configserver

import (
	"flag"
	"time"

	"github.com/caarlos0/env/v6"

	log "github.com/sirupsen/logrus"
)

type ConfigServer struct {
	Address       string        `env:"ADDRESS"`
	StoreInterval time.Duration `env:"STORE_INTERVAL"`
	DebugLevel    log.Level     `env:"DEBUG_LEVEL" envDefault:"debug"`
	StoreFile     string        `env:"STORE_FILE"`
	Restore       bool          `env:"RESTORE" envDefault:"true"`
	Key           string        `env:"KEY"`
	DSN           string        `env:"DATABASE_DSN"`
	//DSN string `env:"DATABASE_DSN" envDefault:"host=localhost port=5432 user=postgres password=postgres dbname=praktikum sslmode=disable"`
	//DSN      string `env:"DATABASE_DSN" envDefault:"host=localhost port=5432 user=postgres password=myPassword dbname=praktikum sslmode=disable"`
	RootPath string `env:"ROOT_PATH" envDefault:"file://./migrations"`
}

// NewConfigServer - creates a new instance with the configuration for the server
func NewConfigServer() *ConfigServer {
	cfgSrv := ConfigServer{}
	// setting flags for the server
	flag.StringVar(&cfgSrv.Address, "a", "localhost:8080", "Адрес сервера")
	flag.BoolVar(&cfgSrv.Restore, "r", true, "Булево значение, определяющее, загружать или нет начальные значения из указанного файла при старте сервера")
	flag.DurationVar(&cfgSrv.StoreInterval, "i", 300*time.Second, "Интервал времени в секундах, по истечении которого текущие показания сервера сбрасываются на диск")
	flag.StringVar(&cfgSrv.StoreFile, "f", "/tmp/devops-metrics-db.json", "Файл где хранятся значения")
	flag.StringVar(&cfgSrv.Key, "k", "", "Ключ для генерации хеша")
	flag.StringVar(&cfgSrv.DSN, "d", "", "База данных")
	flag.Parse()

	err := env.Parse(&cfgSrv)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(cfgSrv)
	return &cfgSrv
}
