package app

import (
	"github.com/shirou/gopsutil/v3/mem"
	"math/rand"
	"runtime"

	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
)

func DataFromRuntime(memAgent *storage.MemStorage, config *config.ConfigAgent, stats *runtime.MemStats, hashServer *config.HashServer) {

	memAgent.PutMetricsGauge("Alloc", float64(stats.Alloc), hashServer.GetHashServer("Alloc", "gauge", 0, float64(stats.Alloc), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("BuckHashSys", float64(stats.BuckHashSys), hashServer.GetHashServer("BuckHashSys", "gauge", 0, float64(stats.BuckHashSys), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("Frees", float64(stats.Frees), hashServer.GetHashServer("Frees", "gauge", 0, float64(stats.Frees), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("GCCPUFraction", float64(stats.GCCPUFraction), hashServer.GetHashServer("GCCPUFraction", "gauge", 0, float64(stats.GCCPUFraction), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("GCSys", float64(stats.GCSys), hashServer.GetHashServer("GCSys", "gauge", 0, float64(stats.GCSys), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("HeapAlloc", float64(stats.HeapAlloc), hashServer.GetHashServer("HeapAlloc", "gauge", 0, float64(stats.HeapAlloc), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("HeapIdle", float64(stats.HeapIdle), hashServer.GetHashServer("HeapIdle", "gauge", 0, float64(stats.HeapIdle), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("HeapInuse", float64(stats.HeapInuse), hashServer.GetHashServer("HeapInuse", "gauge", 0, float64(stats.HeapInuse), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("HeapObjects", float64(stats.HeapObjects), hashServer.GetHashServer("HeapObjects", "gauge", 0, float64(stats.HeapObjects), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("HeapReleased", float64(stats.HeapReleased), hashServer.GetHashServer("HeapReleased", "gauge", 0, float64(stats.HeapReleased), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("HeapSys", float64(stats.HeapSys), hashServer.GetHashServer("HeapSys", "gauge", 0, float64(stats.HeapSys), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("LastGC", float64(stats.LastGC), hashServer.GetHashServer("LastGC", "gauge", 0, float64(stats.LastGC), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("Lookups", float64(stats.Lookups), hashServer.GetHashServer("Lookups", "gauge", 0, float64(stats.Lookups), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("MCacheInuse", float64(stats.MCacheInuse), hashServer.GetHashServer("MCacheInuse", "gauge", 0, float64(stats.MCacheInuse), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("MCacheSys", float64(stats.MCacheSys), hashServer.GetHashServer("MCacheSys", "gauge", 0, float64(stats.MCacheSys), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("MSpanInuse", float64(stats.MSpanInuse), hashServer.GetHashServer("MSpanInuse", "gauge", 0, float64(stats.MSpanInuse), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("MSpanSys", float64(stats.MSpanSys), hashServer.GetHashServer("MSpanSys", "gauge", 0, float64(stats.MSpanSys), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("Mallocs", float64(stats.Mallocs), hashServer.GetHashServer("Mallocs", "gauge", 0, float64(stats.Mallocs), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("NextGC", float64(stats.NextGC), hashServer.GetHashServer("NextGC", "gauge", 0, float64(stats.NextGC), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("NumForcedGC", float64(stats.NumForcedGC), hashServer.GetHashServer("NumForcedGC", "gauge", 0, float64(stats.NumForcedGC), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("NumGC", float64(stats.NumGC), hashServer.GetHashServer("NumGC", "gauge", 0, float64(stats.NumGC), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("OtherSys", float64(stats.OtherSys), hashServer.GetHashServer("OtherSys", "gauge", 0, float64(stats.OtherSys), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("PauseTotalNs", float64(stats.PauseTotalNs), hashServer.GetHashServer("PauseTotalNs", "gauge", 0, float64(stats.PauseTotalNs), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("StackInuse", float64(stats.StackInuse), hashServer.GetHashServer("StackInuse", "gauge", 0, float64(stats.StackInuse), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("StackSys", float64(stats.StackSys), hashServer.GetHashServer("StackSys", "gauge", 0, float64(stats.StackSys), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("Sys", float64(stats.Sys), hashServer.GetHashServer("Sys", "gauge", 0, float64(stats.Sys), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("TotalAlloc", float64(stats.TotalAlloc), hashServer.GetHashServer("TotalAlloc", "gauge", 0, float64(stats.TotalAlloc), config.GetConfigKeyAgent()))
	randV := rand.Float64()
	memAgent.PutMetricsGauge("RandomValue", randV, hashServer.GetHashServer("RandomValue", "gauge", 0, float64(randV), config.GetConfigKeyAgent()))

	pollCount, _, _ := memAgent.GetMetricsCount("PollCount")
	memAgent.PutMetricsCount("PollCount", pollCount+1, hashServer.GetHashServer("PollCount", "counter", pollCount+1, 0, config.GetConfigKeyAgent()))
}

func DataFromRuntimeUsePsutil(memAgent *storage.MemStorage, config *config.ConfigAgent, v *mem.VirtualMemoryStat, hashServer *config.HashServer) {

	memAgent.PutMetricsGauge("TotalMemory", float64(v.Total), hashServer.GetHashServer("TotalMemory", "gauge", 0, float64(v.Total), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("FreeMemory", float64(v.Free), hashServer.GetHashServer("FreeMemory", "gauge", 0, float64(v.Free), config.GetConfigKeyAgent()))
	memAgent.PutMetricsGauge("CPUutilization1", float64(v.UsedPercent), hashServer.GetHashServer("CPUutilization1", "gauge", 0, float64(v.UsedPercent), config.GetConfigKeyAgent()))
}
