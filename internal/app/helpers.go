package app

import (
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"math/rand"
	"runtime"
)

func DataFromRuntime(agent *storage.MemStorage, metric runtime.MemStats) {

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
