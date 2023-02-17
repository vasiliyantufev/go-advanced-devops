package app

import (
	"math/rand"
	"runtime"

	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
)

func DataFromRuntime(mem *storage.MemStorage, config *config.ConfigAgent, stats *runtime.MemStats, hashServer *config.HashServer) {

	mem.PutMetricsGauge("Alloc", float64(stats.Alloc), hashServer.GetHashServer("Alloc", "gauge", 0, float64(stats.Alloc), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("BuckHashSys", float64(stats.BuckHashSys), hashServer.GetHashServer("BuckHashSys", "gauge", 0, float64(stats.BuckHashSys), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("Frees", float64(stats.Frees), hashServer.GetHashServer("Frees", "gauge", 0, float64(stats.Frees), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("GCCPUFraction", float64(stats.GCCPUFraction), hashServer.GetHashServer("GCCPUFraction", "gauge", 0, float64(stats.GCCPUFraction), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("GCSys", float64(stats.GCSys), hashServer.GetHashServer("GCSys", "gauge", 0, float64(stats.GCSys), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("HeapAlloc", float64(stats.HeapAlloc), hashServer.GetHashServer("HeapAlloc", "gauge", 0, float64(stats.HeapAlloc), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("HeapIdle", float64(stats.HeapIdle), hashServer.GetHashServer("HeapIdle", "gauge", 0, float64(stats.HeapIdle), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("HeapInuse", float64(stats.HeapInuse), hashServer.GetHashServer("HeapInuse", "gauge", 0, float64(stats.HeapInuse), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("HeapObjects", float64(stats.HeapObjects), hashServer.GetHashServer("HeapObjects", "gauge", 0, float64(stats.HeapObjects), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("HeapReleased", float64(stats.HeapReleased), hashServer.GetHashServer("HeapReleased", "gauge", 0, float64(stats.HeapReleased), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("HeapSys", float64(stats.HeapSys), hashServer.GetHashServer("HeapSys", "gauge", 0, float64(stats.HeapSys), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("LastGC", float64(stats.LastGC), hashServer.GetHashServer("LastGC", "gauge", 0, float64(stats.LastGC), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("Lookups", float64(stats.Lookups), hashServer.GetHashServer("Lookups", "gauge", 0, float64(stats.Lookups), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("MCacheInuse", float64(stats.MCacheInuse), hashServer.GetHashServer("MCacheInuse", "gauge", 0, float64(stats.MCacheInuse), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("MCacheSys", float64(stats.MCacheSys), hashServer.GetHashServer("MCacheSys", "gauge", 0, float64(stats.MCacheSys), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("MSpanInuse", float64(stats.MSpanInuse), hashServer.GetHashServer("MSpanInuse", "gauge", 0, float64(stats.MSpanInuse), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("MSpanSys", float64(stats.MSpanSys), hashServer.GetHashServer("MSpanSys", "gauge", 0, float64(stats.MSpanSys), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("Mallocs", float64(stats.Mallocs), hashServer.GetHashServer("Mallocs", "gauge", 0, float64(stats.Mallocs), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("NextGC", float64(stats.NextGC), hashServer.GetHashServer("NextGC", "gauge", 0, float64(stats.NextGC), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("NumForcedGC", float64(stats.NumForcedGC), hashServer.GetHashServer("NumForcedGC", "gauge", 0, float64(stats.NumForcedGC), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("NumGC", float64(stats.NumGC), hashServer.GetHashServer("NumGC", "gauge", 0, float64(stats.NumGC), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("OtherSys", float64(stats.OtherSys), hashServer.GetHashServer("OtherSys", "gauge", 0, float64(stats.OtherSys), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("PauseTotalNs", float64(stats.PauseTotalNs), hashServer.GetHashServer("PauseTotalNs", "gauge", 0, float64(stats.PauseTotalNs), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("StackInuse", float64(stats.StackInuse), hashServer.GetHashServer("StackInuse", "gauge", 0, float64(stats.StackInuse), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("StackSys", float64(stats.StackSys), hashServer.GetHashServer("StackSys", "gauge", 0, float64(stats.StackSys), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("Sys", float64(stats.Sys), hashServer.GetHashServer("Sys", "gauge", 0, float64(stats.Sys), config.GetConfigKeyAgent()))
	mem.PutMetricsGauge("TotalAlloc", float64(stats.TotalAlloc), hashServer.GetHashServer("TotalAlloc", "gauge", 0, float64(stats.TotalAlloc), config.GetConfigKeyAgent()))
	randV := rand.Float64()
	mem.PutMetricsGauge("RandomValue", randV, hashServer.GetHashServer("RandomValue", "gauge", 0, float64(randV), config.GetConfigKeyAgent()))

	pollCount, _, _ := mem.GetMetricsCount("PollCount")
	mem.PutMetricsCount("PollCount", pollCount+1, hashServer.GetHashServer("PollCount", "counter", pollCount+1, 0, config.GetConfigKeyAgent()))
}
