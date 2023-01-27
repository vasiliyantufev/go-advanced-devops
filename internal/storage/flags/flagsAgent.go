package flags

import (
	"flag"
)

var fgAgt flagsAgent

// Cтруктура хранения флагов
type flagsAgent struct {
	a string //ADDRESS
	r string //REPORT_INTERVAL
	p string //POLL_INTERVAL
}

// Реализация структуры для взаимодействия с ней
func initFlagsAgent(a, r, p *string) flagsAgent {
	return flagsAgent{
		a: *a,
		r: *r,
		p: *p,
	}
}

func SetFlagsAgent() {
	// Установка флагов
	address := flag.String("a", "localhost:8080", "Адрес сервера")
	reportInterval := flag.String("r", "10s", "Интервал времени в секундах, по истечении которого текущие показания отправляются на сервера.")
	pollInterval := flag.String("p", "2s", "Интервал времени в секундах, по истечении которого текущие показания мертрик обновляются на клиенте")
	flag.Parse()
	flags := initFlagsAgent(address, reportInterval, pollInterval)

	fgAgt.a = flags.a
	fgAgt.r = flags.r
	fgAgt.p = flags.p
}

func GetFlagAddressAgent() string {
	return fgAgt.a
}

func GetFlagReportIntervalAgent() string {
	return fgAgt.r
}

func GetFlagPollIntervalAgent() string {
	return fgAgt.p
}
