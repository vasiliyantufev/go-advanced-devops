package storage

import "time"

//type Config struct {
//	Address        string        `env:"ADDRESS"`
//	ReportInterval time.Duration `env:"REPORT_INTERVAL"`
//	PollInterval   time.Duration `env:"POLL_INTERVAL"`
//}

type Config struct {
	Address        string        `env:"ADDRESS"     envDefault:"localhost:8080"`
	Port           string        `env:"SERVER_PORT" envDefault:":8070"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
	PollInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
}

//type ServerConfig struct {
//	Address string `env:"ADDRESS" envDefault:"localhost:8080"`
//	StoreInterval time.Duration env:"STORE_INTERVAL" envDefault:"300s"
//	StoreFile string env:"STORE_FILE"
//	Restore bool env:"RESTORE" envDefault:"true"
//}
