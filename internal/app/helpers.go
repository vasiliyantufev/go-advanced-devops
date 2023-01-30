package app

import (
	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"math/rand"
	"runtime"
)

func DataFromRuntime(agent *storage.MemStorage, stats *runtime.MemStats) {

	agent.PutMetricsGauge("Alloc", float64(stats.Alloc), config.GetHashAgent("Alloc", "gauge", 0, float64(stats.Alloc)))
	agent.PutMetricsGauge("BuckHashSys", float64(stats.BuckHashSys), config.GetHashAgent("BuckHashSys", "gauge", 0, float64(stats.BuckHashSys)))
	agent.PutMetricsGauge("Frees", float64(stats.Frees), config.GetHashAgent("Frees", "gauge", 0, float64(stats.Frees)))
	agent.PutMetricsGauge("GCCPUFraction", float64(stats.GCCPUFraction), config.GetHashAgent("GCCPUFraction", "gauge", 0, float64(stats.GCCPUFraction)))
	agent.PutMetricsGauge("GCSys", float64(stats.GCSys), config.GetHashAgent("GCSys", "gauge", 0, float64(stats.GCSys)))
	agent.PutMetricsGauge("HeapAlloc", float64(stats.HeapAlloc), config.GetHashAgent("HeapAlloc", "gauge", 0, float64(stats.HeapAlloc)))
	agent.PutMetricsGauge("HeapIdle", float64(stats.HeapIdle), config.GetHashAgent("HeapIdle", "gauge", 0, float64(stats.HeapIdle)))
	agent.PutMetricsGauge("HeapInuse", float64(stats.HeapInuse), config.GetHashAgent("HeapInuse", "gauge", 0, float64(stats.HeapInuse)))
	agent.PutMetricsGauge("HeapObjects", float64(stats.HeapObjects), config.GetHashAgent("HeapObjects", "gauge", 0, float64(stats.HeapObjects)))
	agent.PutMetricsGauge("HeapReleased", float64(stats.HeapReleased), config.GetHashAgent("", "gauge", 0, float64(stats.HeapReleased)))
	agent.PutMetricsGauge("HeapSys", float64(stats.HeapSys), config.GetHashAgent("HeapSys", "gauge", 0, float64(stats.HeapSys)))
	agent.PutMetricsGauge("LastGC", float64(stats.LastGC), config.GetHashAgent("LastGC", "gauge", 0, float64(stats.LastGC)))
	agent.PutMetricsGauge("Lookups", float64(stats.Lookups), config.GetHashAgent("Lookups", "gauge", 0, float64(stats.Lookups)))
	agent.PutMetricsGauge("MCacheInuse", float64(stats.MCacheInuse), config.GetHashAgent("MCacheInuse", "gauge", 0, float64(stats.MCacheInuse)))
	agent.PutMetricsGauge("MCacheSys", float64(stats.MCacheSys), config.GetHashAgent("MCacheSys", "gauge", 0, float64(stats.MCacheSys)))
	agent.PutMetricsGauge("MSpanInuse", float64(stats.MSpanInuse), config.GetHashAgent("MSpanInuse", "gauge", 0, float64(stats.MSpanInuse)))
	agent.PutMetricsGauge("MSpanSys", float64(stats.MSpanSys), config.GetHashAgent("MSpanSys", "gauge", 0, float64(stats.MSpanSys)))
	agent.PutMetricsGauge("Mallocs", float64(stats.Mallocs), config.GetHashAgent("Mallocs", "gauge", 0, float64(stats.Mallocs)))
	agent.PutMetricsGauge("NextGC", float64(stats.NextGC), config.GetHashAgent("NextGC", "gauge", 0, float64(stats.NextGC)))
	agent.PutMetricsGauge("NumForcedGC", float64(stats.NumForcedGC), config.GetHashAgent("NumForcedGC", "gauge", 0, float64(stats.NumForcedGC)))
	agent.PutMetricsGauge("NumGC", float64(stats.NumGC), config.GetHashAgent("NumGC", "gauge", 0, float64(stats.NumGC)))
	agent.PutMetricsGauge("OtherSys", float64(stats.OtherSys), config.GetHashAgent("OtherSys", "gauge", 0, float64(stats.OtherSys)))
	agent.PutMetricsGauge("PauseTotalNs", float64(stats.PauseTotalNs), config.GetHashAgent("PauseTotalNs", "gauge", 0, float64(stats.PauseTotalNs)))
	agent.PutMetricsGauge("StackInuse", float64(stats.StackInuse), config.GetHashAgent("StackInuse", "gauge", 0, float64(stats.StackInuse)))
	agent.PutMetricsGauge("StackSys", float64(stats.StackSys), config.GetHashAgent("StackSys", "gauge", 0, float64(stats.StackSys)))
	agent.PutMetricsGauge("Sys", float64(stats.Sys), config.GetHashAgent("Sys", "gauge", 0, float64(stats.Sys)))
	agent.PutMetricsGauge("TotalAlloc", float64(stats.TotalAlloc), config.GetHashAgent("TotalAlloc", "gauge", 0, float64(stats.TotalAlloc)))
	randV := rand.Float64()
	agent.PutMetricsGauge("RandomValue", randV, config.GetHashAgent("RandomValue", "gauge", 0, float64(randV)))

	pollCount, _, _ := agent.GetMetricsCount("PollCount")
	agent.PutMetricsCount("PollCount", pollCount+1, config.GetHashAgent("PollCount", "counter", pollCount+1, 0))
}
