package flags

import (
	"flag"
)

// Cтруктура хранения флагов
type flagsAgent struct {
	Address        string
	ReportInterval string
	PollInterval   string
	Key            string
	RateLimit      int
}

// Реализация структуры для взаимодействия с ней
func initFlagsAgent(address, reportInterval, pollInterval, key *string, rateLimit *int) flagsAgent {
	return flagsAgent{
		Address:        *address,
		ReportInterval: *reportInterval,
		PollInterval:   *pollInterval,
		Key:            *key,
		RateLimit:      *rateLimit,
	}
}

func NewFlagsAgent() *flagsAgent {

	var fgAgt flagsAgent

	// Установка флагов
	address := flag.String("a", "localhost:8080", "Адрес сервера")
	reportInterval := flag.String("r", "10s", "Интервал времени в секундах, по истечении которого текущие показания отправляются на сервера.")
	pollInterval := flag.String("p", "2s", "Интервал времени в секундах, по истечении которого текущие показания мертрик обновляются на клиенте")
	key := flag.String("k", "key", "Ключ для генерации хеша")
	rateLimit := flag.Int("l", 100, "Количество одновременно исходящих запросов на сервер")

	flag.Parse()
	flags := initFlagsAgent(address, reportInterval, pollInterval, key, rateLimit)
	fgAgt.Address = flags.Address
	fgAgt.ReportInterval = flags.ReportInterval
	fgAgt.PollInterval = flags.PollInterval
	fgAgt.Key = flags.Key
	fgAgt.RateLimit = flags.RateLimit

	return &fgAgt
}
