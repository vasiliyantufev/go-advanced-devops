package app

import (
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/jsonmetrics"
	"math/rand"
	"runtime"
)

func DateFromFile(agent *storage.MemStorage, metric jsonmetrics.JSONMetricsFromFile) {

	agent.PutMetricsGauge("Alloc", float64(metric.DataMetricsGauge.Alloc))
	agent.PutMetricsGauge("BuckHashSys", float64(metric.DataMetricsGauge.BuckHashSys))
	agent.PutMetricsGauge("Frees", float64(metric.DataMetricsGauge.Frees))
	agent.PutMetricsGauge("GCCPUFraction", float64(metric.DataMetricsGauge.GCCPUFraction))
	agent.PutMetricsGauge("GCSys", float64(metric.DataMetricsGauge.GCSys))
	agent.PutMetricsGauge("HeapAlloc", float64(metric.DataMetricsGauge.HeapAlloc))
	agent.PutMetricsGauge("HeapIdle", float64(metric.DataMetricsGauge.HeapIdle))
	agent.PutMetricsGauge("HeapInuse", float64(metric.DataMetricsGauge.HeapInuse))
	agent.PutMetricsGauge("HeapObjects", float64(metric.DataMetricsGauge.HeapObjects))
	agent.PutMetricsGauge("HeapReleased", float64(metric.DataMetricsGauge.HeapReleased))
	agent.PutMetricsGauge("HeapSys", float64(metric.DataMetricsGauge.HeapSys))
	agent.PutMetricsGauge("LastGC", float64(metric.DataMetricsGauge.LastGC))
	agent.PutMetricsGauge("Lookups", float64(metric.DataMetricsGauge.Lookups))
	agent.PutMetricsGauge("MCacheInuse", float64(metric.DataMetricsGauge.MCacheInuse))
	agent.PutMetricsGauge("MCacheSys", float64(metric.DataMetricsGauge.MCacheSys))
	agent.PutMetricsGauge("MSpanInuse", float64(metric.DataMetricsGauge.MSpanInuse))
	agent.PutMetricsGauge("MSpanSys", float64(metric.DataMetricsGauge.MSpanSys))
	agent.PutMetricsGauge("Mallocs", float64(metric.DataMetricsGauge.Mallocs))
	agent.PutMetricsGauge("NextGC", float64(metric.DataMetricsGauge.NextGC))
	agent.PutMetricsGauge("NumForcedGC", float64(metric.DataMetricsGauge.NumForcedGC))
	agent.PutMetricsGauge("NumGC", float64(metric.DataMetricsGauge.NumGC))
	agent.PutMetricsGauge("OtherSys", float64(metric.DataMetricsGauge.OtherSys))
	agent.PutMetricsGauge("PauseTotalNs", float64(metric.DataMetricsGauge.PauseTotalNs))
	agent.PutMetricsGauge("StackInuse", float64(metric.DataMetricsGauge.StackInuse))
	agent.PutMetricsGauge("StackSys", float64(metric.DataMetricsGauge.StackSys))
	agent.PutMetricsGauge("Sys", float64(metric.DataMetricsGauge.Sys))
	agent.PutMetricsGauge("TotalAlloc", float64(metric.DataMetricsGauge.TotalAlloc))
	agent.PutMetricsGauge("RandomValue", float64(metric.DataMetricsGauge.RandomValue))
	agent.PutMetricsCount("PollCount", int64(metric.DataMetricsCount.PollCount))

	log.Print(agent)
}

func DateFromRuntime(agent *storage.MemStorage, metric runtime.MemStats) {

	agent.PutMetricsGauge("Alloc", float64(metric.Alloc))
	agent.PutMetricsGauge("BuckHashSys", float64(metric.BuckHashSys))
	agent.PutMetricsGauge("Frees", float64(metric.Frees))
	agent.PutMetricsGauge("GCCPUFraction", float64(metric.GCCPUFraction))
	agent.PutMetricsGauge("GCSys", float64(metric.GCSys))
	agent.PutMetricsGauge("HeapAlloc", float64(metric.HeapAlloc))
	agent.PutMetricsGauge("HeapIdle", float64(metric.HeapIdle))
	agent.PutMetricsGauge("HeapInuse", float64(metric.HeapInuse))
	agent.PutMetricsGauge("HeapObjects", float64(metric.HeapObjects))
	agent.PutMetricsGauge("HeapReleased", float64(metric.HeapReleased))
	agent.PutMetricsGauge("HeapSys", float64(metric.HeapSys))
	agent.PutMetricsGauge("LastGC", float64(metric.LastGC))
	agent.PutMetricsGauge("Lookups", float64(metric.Lookups))
	agent.PutMetricsGauge("MCacheInuse", float64(metric.MCacheInuse))
	agent.PutMetricsGauge("MCacheSys", float64(metric.MCacheSys))
	agent.PutMetricsGauge("MSpanInuse", float64(metric.MSpanInuse))
	agent.PutMetricsGauge("MSpanSys", float64(metric.MSpanSys))
	agent.PutMetricsGauge("Mallocs", float64(metric.Mallocs))
	agent.PutMetricsGauge("NextGC", float64(metric.NextGC))
	agent.PutMetricsGauge("NumForcedGC", float64(metric.NumForcedGC))
	agent.PutMetricsGauge("NumGC", float64(metric.NumGC))
	agent.PutMetricsGauge("OtherSys", float64(metric.OtherSys))
	agent.PutMetricsGauge("PauseTotalNs", float64(metric.PauseTotalNs))
	agent.PutMetricsGauge("StackInuse", float64(metric.StackInuse))
	agent.PutMetricsGauge("StackSys", float64(metric.StackSys))
	agent.PutMetricsGauge("Sys", float64(metric.Sys))
	agent.PutMetricsGauge("TotalAlloc", float64(metric.TotalAlloc))
	agent.PutMetricsGauge("RandomValue", rand.Float64())
	pollCount, _ := agent.GetMetricsCount("PollCount")
	agent.PutMetricsCount("PollCount", pollCount+1)
}
