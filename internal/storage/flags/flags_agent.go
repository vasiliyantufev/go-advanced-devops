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
}

// Реализация структуры для взаимодействия с ней
func initFlagsAgent(address, reportInterval, pollInterval, key *string) flagsAgent {
	return flagsAgent{
		Address:        *address,
		ReportInterval: *reportInterval,
		PollInterval:   *pollInterval,
		Key:            *key,
	}
}

func NewFlagsAgent() *flagsAgent {

	var fgAgt flagsAgent

	// Установка флагов
	address := flag.String("a", "localhost:8080", "Адрес сервера")
	reportInterval := flag.String("r", "10s", "Интервал времени в секундах, по истечении которого текущие показания отправляются на сервера.")
	pollInterval := flag.String("p", "2s", "Интервал времени в секундах, по истечении которого текущие показания мертрик обновляются на клиенте")
	key := flag.String("k", "key", "Ключ для генерации хеша")

	flag.Parse()
	flags := initFlagsAgent(address, reportInterval, pollInterval, key)
	fgAgt.Address = flags.Address
	fgAgt.ReportInterval = flags.ReportInterval
	fgAgt.PollInterval = flags.PollInterval
	fgAgt.Key = flags.Key

	return &fgAgt
}
