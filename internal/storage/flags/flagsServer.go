package flags

import "flag"

var fgSrv flagsServer

// Cтруктура хранения флагов
type flagsServer struct {
	a string //ADDRESS
	f string //STORE_FILE
	i string //STORE_INTERVAL
	r bool   //RESTORE
}

// Реализация структуры для взаимодействия с ней
func initFlagsServer(a, f, i *string, r *bool) flagsServer {
	return flagsServer{
		a: *a,
		f: *f,
		r: *r,
		i: *i,
	}
}

func SetFlagsServer() {
	// Установка флагов
	address := flag.String("a", "localhost:8080", "Адрес сервера")
	file := flag.String("f", "/tmp/devops-metrics-db.json", "Строка, имя файла, где хранятся значения")
	interval := flag.String("i", "2s", "Интервал времени в секундах, по истечении которого текущие показания сервера сбрасываются на диск ")
	restore := flag.Bool("r", true, "Булево значение, определяющее, загружать или нет начальные значения из указанного файла при старте сервера")
	flag.Parse()
	flags := initFlagsServer(address, file, interval, restore)

	fgSrv.a = flags.a
	fgSrv.f = flags.f
	fgSrv.i = flags.i
	fgSrv.r = flags.r
}

func GetFlagAddressServer() string {
	return fgSrv.a
}

func GetFlagStoreFileServer() string {
	return fgSrv.f
}

func GetFlagStoreIntervalServer() string {
	return fgSrv.i
}

func GetFlagRestoreServer() bool {
	return fgSrv.r
}
