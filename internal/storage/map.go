package storage

var MetricsGauge = make(map[string]float64)
var MetricsCounter = make(map[string]int64)

func InitMap() {

	MetricsCounter["poll_count"] = 0
	MetricsCounter["random_value"] = 0

	MetricsGauge["alloc"] = 0
	MetricsGauge["buck_hash_sys"] = 0
	MetricsGauge["Frees"] = 0
	MetricsGauge["GCCPUFraction"] = 0
	MetricsGauge["GCSys"] = 0
	MetricsGauge["HeapAlloc"] = 0
	MetricsGauge["HeapIdle"] = 0
	MetricsGauge["HeapInuse"] = 0
	MetricsGauge["HeapObjects"] = 0
	MetricsGauge["HeapReleased"] = 0
	MetricsGauge["HeapSys"] = 0
	MetricsGauge["LastGC"] = 0
	MetricsGauge["Lookups"] = 0
	MetricsGauge["MCacheInuse"] = 0
	MetricsGauge["MCacheSys"] = 0
	MetricsGauge["MSpanInuse"] = 0
	MetricsGauge["MSpanSys"] = 0
	MetricsGauge["Mallocs"] = 0
	MetricsGauge["NextGC"] = 0
	MetricsGauge["NumForcedGC"] = 0
	MetricsGauge["NumGC"] = 0
	MetricsGauge["OtherSys"] = 0
	MetricsGauge["PauseTotalNs"] = 0
	MetricsGauge["StackInuse"] = 0
	MetricsGauge["StackSys"] = 0
	MetricsGauge["Sys"] = 0
	MetricsGauge["TotalAlloc"] = 0
}
