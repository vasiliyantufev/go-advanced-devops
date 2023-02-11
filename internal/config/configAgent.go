package config

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/flags"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

var cfgAgt configAgent

type configAgent struct {
	//Address        string        `env:"ADDRESS" envDefault:"localhost:8080"`
	Address string `env:"ADDRESS"`
	//ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL"`
	//PollInterval time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
	PollInterval time.Duration `env:"POLL_INTERVAL"`
	//PollInterval time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
	Key string `env:"KEY" envDefault:""`
}

func SetConfigAgent() {

	err := env.Parse(&cfgAgt)
	if err != nil {
		log.Fatal(err)
	}

	flags.SetFlagsAgent()

	if cfgAgt.Address == "" {
		cfgAgt.Address = flags.GetFlagAddressAgent()
	}
	if cfgAgt.ReportInterval.String() == "0s" {
		cfgAgt.ReportInterval, _ = time.ParseDuration(flags.GetFlagReportIntervalAgent())
	}
	if cfgAgt.PollInterval.String() == "0s" {
		cfgAgt.PollInterval, _ = time.ParseDuration(flags.GetFlagPollIntervalAgent())
	}
	if cfgAgt.Key == "" {
		cfgAgt.Key = flags.GetKeyAgent()
	}

	log.Info(cfgAgt)
}

func GetConfigAddressAgent() string {
	return cfgAgt.Address
}

func GetConfigReportIntervalAgent() time.Duration {
	return cfgAgt.ReportInterval
}

func GetConfigPollIntervalAgent() time.Duration {
	return cfgAgt.PollInterval
}

func GetConfigKeyAgent() string {
	return cfgAgt.Key
}

func GetHashAgent(mid string, mtype string, delta int64, value float64) string {

	var data string
	switch mtype {
	case "counter":
		data = fmt.Sprintf("%s:%s:%d", mid, mtype, delta)
	case "gauge":
		data = fmt.Sprintf("%s:%s:%f", mid, mtype, value)
	}

	h := hmac.New(sha256.New, []byte(cfgAgt.Key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
