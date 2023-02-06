package flags

import (
	"flag"
)

var fgAgt flagsAgent

// Cтруктура хранения флагов
type flagsAgent struct {
	a        string //ADDRESS
	r        string //REPORT_INTERVAL
	p        string //POLL_INTERVAL
	k        string //KEY
	buildvcs bool
}

// Реализация структуры для взаимодействия с ней
func initFlagsAgent(a, r, p, k *string, buildvcs *bool) flagsAgent {
	return flagsAgent{
		a:        *a,
		r:        *r,
		p:        *p,
		k:        *k,
		buildvcs: *buildvcs,
	}
}

func SetFlagsAgent() {
	// Установка флагов
	address := flag.String("a", "localhost:8080", "Адрес сервера")
	reportInterval := flag.String("r", "10s", "Интервал времени в секундах, по истечении которого текущие показания отправляются на сервера.")
	pollInterval := flag.String("p", "2s", "Интервал времени в секундах, по истечении которого текущие показания мертрик обновляются на клиенте")
	key := flag.String("k", "key", "Ключ для генерации хеша")
	buildvcs := flag.Bool("buildvcs", false, "")
	flag.Parse()
	flags := initFlagsAgent(address, reportInterval, pollInterval, key, buildvcs)

	fgAgt.a = flags.a
	fgAgt.r = flags.r
	fgAgt.p = flags.p
	fgAgt.k = flags.k
	fgAgt.buildvcs = flags.buildvcs
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

func GetKeyAgent() string {
	return fgAgt.k
}
