package storage

import "time"

type Config struct {
	Address         string        `env:"ADDRESS"`
	Report_interval time.Duration `env:"REPORT_INTERVAL"`
	Poll_interval   time.Duration `env:"POLL_INTERVAL"`
}
