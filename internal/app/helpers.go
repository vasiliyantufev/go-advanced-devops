package app

import (
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"math/rand"
	"runtime"
)

func DataFromRuntime(agent *storage.MemStorage, stats *runtime.MemStats) {

	agent.PutMetricsGauge("Alloc", float64(stats.Alloc))
	agent.PutMetricsGauge("BuckHashSys", float64(stats.BuckHashSys))
	agent.PutMetricsGauge("Frees", float64(stats.Frees))
	agent.PutMetricsGauge("GCCPUFraction", float64(stats.GCCPUFraction))
	agent.PutMetricsGauge("GCSys", float64(stats.GCSys))
	agent.PutMetricsGauge("HeapAlloc", float64(stats.HeapAlloc))
	agent.PutMetricsGauge("HeapIdle", float64(stats.HeapIdle))
	agent.PutMetricsGauge("HeapInuse", float64(stats.HeapInuse))
	agent.PutMetricsGauge("HeapObjects", float64(stats.HeapObjects))
	agent.PutMetricsGauge("HeapReleased", float64(stats.HeapReleased))
	agent.PutMetricsGauge("HeapSys", float64(stats.HeapSys))
	agent.PutMetricsGauge("LastGC", float64(stats.LastGC))
	agent.PutMetricsGauge("Lookups", float64(stats.Lookups))
	agent.PutMetricsGauge("MCacheInuse", float64(stats.MCacheInuse))
	agent.PutMetricsGauge("MCacheSys", float64(stats.MCacheSys))
	agent.PutMetricsGauge("MSpanInuse", float64(stats.MSpanInuse))
	agent.PutMetricsGauge("MSpanSys", float64(stats.MSpanSys))
	agent.PutMetricsGauge("Mallocs", float64(stats.Mallocs))
	agent.PutMetricsGauge("NextGC", float64(stats.NextGC))
	agent.PutMetricsGauge("NumForcedGC", float64(stats.NumForcedGC))
	agent.PutMetricsGauge("NumGC", float64(stats.NumGC))
	agent.PutMetricsGauge("OtherSys", float64(stats.OtherSys))
	agent.PutMetricsGauge("PauseTotalNs", float64(stats.PauseTotalNs))
	agent.PutMetricsGauge("StackInuse", float64(stats.StackInuse))
	agent.PutMetricsGauge("StackSys", float64(stats.StackSys))
	agent.PutMetricsGauge("Sys", float64(stats.Sys))
	agent.PutMetricsGauge("TotalAlloc", float64(stats.TotalAlloc))
	agent.PutMetricsGauge("RandomValue", rand.Float64())
	pollCount, _ := agent.GetMetricsCount("PollCount")
	agent.PutMetricsCount("PollCount", pollCount+1)
}
