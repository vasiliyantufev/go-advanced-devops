package storage

import (
	"flag"
)

// Cтруктура хранения флагов
type FlagsServer struct {
	a string //ADDRESS
	f string //STORE_FILE
	i string //STORE_INTERVAL
	r bool   //RESTORE
}

// Реализация структуры для взаимодействия с ней
func InitFlagsServer(a, f, i *string, r *bool) FlagsServer {
	return FlagsServer{
		a: *a,
		f: *f,
		r: *r,
		i: *i,
	}
}

func GetFlagsServer() FlagsServer {
	// Установка флагов
	address := flag.String("a", "localhost:8080", "Адрес сервера")
	file := flag.String("f", "/tmp/devops-metrics-db.json", "Строка, имя файла, где хранятся значения")
	interval := flag.String("i", "22s", "Интервал времени в секундах, по истечении которого текущие показания сервера сбрасываются на диск ")
	restore := flag.Bool("r", true, "Булево значение, определяющее, загружать или нет начальные значения из указанного файла при старте сервера")

	flag.Parse()
	flags := InitFlagsServer(address, file, interval, restore)

	return flags
}
