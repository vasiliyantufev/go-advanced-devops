package storage

import (
	"flag"
)

var FgAgt FlagsAgent

// Cтруктура хранения флагов
type FlagsAgent struct {
	a string //ADDRESS
	r string //REPORT_INTERVAL
	p string //POLL_INTERVAL
}

// Реализация структуры для взаимодействия с ней
func InitFlagsAgent(a, r, p *string) FlagsAgent {
	return FlagsAgent{
		a: *a,
		r: *r,
		p: *p,
	}
}

func GetFlagsAgent() {
	// Установка флагов
	address := flag.String("a", "localhost:8080", "Адрес сервера")
	report_interval := flag.String("r", "10s", "Интервал времени в секундах, по истечении которого текущие показания отправляются на сервера.")
	poll_interval := flag.String("p", "2s", "Интервал времени в секундах, по истечении которого текущие показания мертрик обновляются на клиенте")
	flag.Parse()
	flags := InitFlagsAgent(address, report_interval, poll_interval)

	FgAgt.a = flags.a
	FgAgt.r = flags.r
	FgAgt.p = flags.p
}

func GetFlagAddressAgent() string {
	return FgAgt.a
}

func GetFlagReportIntervalAgent() string {
	return FgAgt.r
}

func GetFlagPollIntervalAgent() string {
	return FgAgt.p
}
