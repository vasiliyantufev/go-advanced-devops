package config

import (
	"flag"
	"time"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type ConfigServicerAgent interface {
	GetConfigAddressAgent() string
	GetConfigReportIntervalAgent() time.Duration
	GetConfigPollIntervalAgent() time.Duration
	GetConfigKeyAgent() string
	GetConfigRateLimitAgent() int
}

type ConfigAgent struct {
	Address        string        `env:"ADDRESS"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL"`
	PollInterval   time.Duration `env:"POLL_INTERVAL"`
	Key            string        `env:"KEY"`
	RateLimit      int           `env:"RATE_LIMIT"`
}

func NewConfigAgent() *ConfigAgent {

	var cfgAgt ConfigAgent

	flag.StringVar(&cfgAgt.Address, "a", "localhost:8080", "Адрес сервера")
	flag.DurationVar(&cfgAgt.ReportInterval, "r", 10*time.Second, "Интервал времени в секундах, по истечении которого текущие показания отправляются на сервера")
	flag.DurationVar(&cfgAgt.PollInterval, "p", 2*time.Second, "Интервал времени в секундах, по истечении которого текущие показания мертрик обновляются на клиенте")
	flag.StringVar(&cfgAgt.Key, "k", "", "Ключ для генерации хеша")
	flag.IntVar(&cfgAgt.RateLimit, "l", 2, "Количество одновременно исходящих запросов на сервер")
	flag.Parse()

	err := env.Parse(&cfgAgt)
	if err != nil {
		log.Fatal(err)
	}

	return &cfgAgt
}

func (cfg ConfigAgent) GetConfigAddressAgent() string {
	return cfg.Address
}

func (cfg ConfigAgent) GetConfigReportIntervalAgent() time.Duration {
	return cfg.ReportInterval
}

func (cfg ConfigAgent) GetConfigPollIntervalAgent() time.Duration {
	return cfg.PollInterval
}

func (cfg ConfigAgent) GetConfigKeyAgent() string {
	return cfg.Key
}

func (cfg ConfigAgent) GetConfigRateLimitAgent() int {
	return cfg.RateLimit
}
