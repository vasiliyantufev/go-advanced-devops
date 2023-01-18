package jsonMetrics

type JsonMetricsFromFile struct {
	DataMetricsGauge InnerGauge   `json:"DataMetricsGauge"`
	DataMetricsCount InnerCounter `json:"DataMetricsCount"`
}

type InnerCounter struct {
	PollCount int64 `json:"PollCount"`
}
type InnerGauge struct {
	Alloc         float64 `json:"Alloc"`
	BuckHashSys   float64 `json:"BuckHashSys"`
	Frees         float64 `json:"Frees"`
	GCCPUFraction float64 `json:"GCCPUFraction"`
	GCSys         float64 `json:"GCSys"`
	HeapAlloc     float64 `json:"HeapAlloc"`
	HeapIdle      float64 `json:"HeapIdle"`
	HeapInuse     float64 `json:"HeapInuse"`
	HeapObjects   float64 `json:"HeapObjects"`
	HeapReleased  float64 `json:"HeapReleased"`
	HeapSys       float64 `json:"HeapSys"`
	LastGC        float64 `json:"LastGC"`
	Lookups       float64 `json:"Lookups"`
	MCacheInuse   float64 `json:"MCacheInuse"`
	MCacheSys     float64 `json:"MCacheSys"`
	MSpanInuse    float64 `json:"MSpanInuse"`
	MSpanSys      float64 `json:"MSpanSys"`
	Mallocs       float64 `json:"Mallocs"`
	NextGC        float64 `json:"NextGC"`
	NumForcedGC   float64 `json:"NumForcedGC"`
	NumGC         float64 `json:"NumGC"`
	OtherSys      float64 `json:"OtherSys"`
	PauseTotalNs  float64 `json:"PauseTotalNs"`
	StackInuse    float64 `json:"StackInuse"`
	StackSys      float64 `json:"StackSys"`
	Sys           float64 `json:"Sys"`
	TotalAlloc    float64 `json:"TotalAlloc"`
	RandomValue   float64 `json:"RandomValue"`
}
