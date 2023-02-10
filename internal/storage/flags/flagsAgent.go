package flags

import (
	"flag"
)

var fgAgt flagsAgent

// Cтруктура хранения флагов
type flagsAgent struct {
	address        string
	reportInterval string
	pollInterval   string
	key            string
}

// Реализация структуры для взаимодействия с ней
func initFlagsAgent(a, r, p, k *string) flagsAgent {
	return flagsAgent{
		address:        *a,
		reportInterval: *r,
		pollInterval:   *p,
		key:            *k,
	}
}

func SetFlagsAgent() {
	// Установка флагов
	address := flag.String("a", "localhost:8080", "Адрес сервера")
	reportInterval := flag.String("r", "10s", "Интервал времени в секундах, по истечении которого текущие показания отправляются на сервера.")
	pollInterval := flag.String("p", "2s", "Интервал времени в секундах, по истечении которого текущие показания мертрик обновляются на клиенте")
	key := flag.String("k", "key", "Ключ для генерации хеша")

	flag.Parse()
	flags := initFlagsAgent(address, reportInterval, pollInterval, key)
	fgAgt.address = flags.address
	fgAgt.reportInterval = flags.reportInterval
	fgAgt.pollInterval = flags.pollInterval
	fgAgt.key = flags.key
}

func GetFlagAddressAgent() string {
	return fgAgt.address
}

func GetFlagReportIntervalAgent() string {
	return fgAgt.reportInterval
}

func GetFlagPollIntervalAgent() string {
	return fgAgt.pollInterval
}

func GetKeyAgent() string {
	return fgAgt.key
}
