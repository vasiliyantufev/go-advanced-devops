package flags

import (
	"flag"
)

// Cтруктура хранения флагов
type flagsServer struct {
	Address       string
	StoreFile     string
	StoreInterval string
	Key           string
	DNS           string
	Restore       bool
}

// Реализация структуры для взаимодействия с ней
func initFlagsServer(address, storeFile, storeInterval, key, DNS *string, restore *bool) flagsServer {
	return flagsServer{
		Address:       *address,
		StoreFile:     *storeFile,
		StoreInterval: *storeInterval,
		Key:           *key,
		DNS:           *DNS,
		Restore:       *restore,
	}
}

func NewFlagsServer() flagsServer {

	var fgSrv flagsServer

	// Установка флагов
	address := flag.String("a", "localhost:8080", "Адрес сервера")
	file := flag.String("f", "/tmp/devops-metrics-db.json", "Строка, имя файла, где хранятся значения")
	interval := flag.String("i", "300s", "Интервал времени в секундах, по истечении которого текущие показания сервера сбрасываются на диск ")
	restore := flag.Bool("r", true, "Булево значение, определяющее, загружать или нет начальные значения из указанного файла при старте сервера")
	key := flag.String("k", "key", "Ключ для генерации хеша")
	db := flag.String("d", "db", "База данных")

	flag.Parse()
	flags := initFlagsServer(address, file, interval, key, db, restore)
	fgSrv.Address = flags.Address
	fgSrv.StoreFile = flags.StoreFile
	fgSrv.StoreInterval = flags.StoreInterval
	fgSrv.Key = flags.Key
	fgSrv.Restore = flags.Restore
	fgSrv.DNS = flags.DNS

	//log.Fatal(fgSrv.DNS)

	return fgSrv
}
