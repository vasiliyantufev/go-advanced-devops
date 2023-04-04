// Module config server
package configServer

import (
	"flag"
	"time"

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

// Creates a new instance with the configuration for the server
func NewConfigServer() *ConfigServer {

	//f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	//
	//cfgSrv := ConfigServer{}
	//
	//// Установка флагов
	//f.StringVar(&cfgSrv.Address, "a", "localhost:8080", "Адрес сервера")
	//f.BoolVar(&cfgSrv.Restore, "r", true, "Булево значение, определяющее, загружать или нет начальные значения из указанного файла при старте сервера")
	//f.DurationVar(&cfgSrv.StoreInterval, "i", 300*time.Second, "Интервал времени в секундах, по истечении которого текущие показания сервера сбрасываются на диск")
	//f.StringVar(&cfgSrv.StoreFile, "f", "/tmp/devops-metrics-db.json", "Файл где хранятся значения")
	//f.StringVar(&cfgSrv.Key, "k", "", "Ключ для генерации хеша")
	//f.StringVar(&cfgSrv.DSN, "d", "", "База данных")
	//f.Parse(os.Args)

	cfgSrv := ConfigServer{}

	// Установка флагов
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
	return cfg.DSN
}
