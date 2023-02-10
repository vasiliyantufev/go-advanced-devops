package flags

import "flag"

var fgSrv flagsServer

// Cтруктура хранения флагов
type flagsServer struct {
	address       string
	storeFile     string
	storeInterval string
	key           string
	databaseDns   string
	restore       bool
}

// Реализация структуры для взаимодействия с ней
func initFlagsServer(a, f, i, k, d *string, r *bool) flagsServer {
	return flagsServer{
		address:       *a,
		storeFile:     *f,
		storeInterval: *i,
		key:           *k,
		databaseDns:   *d,
		restore:       *r,
	}
}

func SetFlagsServer() {
	// Установка флагов
	address := flag.String("a", "localhost:8080", "Адрес сервера")
	file := flag.String("f", "/tmp/devops-metrics-db.json", "Строка, имя файла, где хранятся значения")
	interval := flag.String("i", "300s", "Интервал времени в секундах, по истечении которого текущие показания сервера сбрасываются на диск ")
	restore := flag.Bool("r", true, "Булево значение, определяющее, загружать или нет начальные значения из указанного файла при старте сервера")
	key := flag.String("k", "key", "Ключ для генерации хеша")
	db := flag.String("d", "db", "База данных")

	flag.Parse()
	flags := initFlagsServer(address, file, interval, key, db, restore)
	fgSrv.address = flags.address
	fgSrv.storeFile = flags.storeFile
	fgSrv.storeInterval = flags.storeInterval
	fgSrv.key = flags.key
	fgSrv.restore = flags.restore
	fgSrv.databaseDns = flags.databaseDns
}

func GetFlagAddressServer() string {
	return fgSrv.address
}

func GetFlagStoreFileServer() string {
	return fgSrv.storeFile
}

func GetFlagStoreIntervalServer() string {
	return fgSrv.storeInterval
}

func GetFlagRestoreServer() bool {
	return fgSrv.restore
}

func GetFlagKeyServer() string {
	return fgSrv.key
}

func GetFlagDataBaseServer() string {
	return fgSrv.databaseDns
}
