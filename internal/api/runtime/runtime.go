// Package runtime
package runtime

import (
	"math/rand"
	"runtime"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/converter"
	"github.com/vasiliyantufev/go-advanced-devops/internal/models"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"
)

// Getting data using the runtime package
func DataFromRuntime(memAgent *memstorage.MemStorage, stats *runtime.MemStats, hashServer *hashservicer.HashServer) {
	memAgent.PutMetricsGauge("Alloc", float64(stats.Alloc), hashServer.GenerateHash(models.Metric{ID: "Alloc", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.Alloc)}))
	memAgent.PutMetricsGauge("BuckHashSys", float64(stats.BuckHashSys), hashServer.GenerateHash(models.Metric{ID: "BuckHashSys", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.BuckHashSys)}))
	memAgent.PutMetricsGauge("Frees", float64(stats.Frees), hashServer.GenerateHash(models.Metric{ID: "Frees", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.Frees)}))
	memAgent.PutMetricsGauge("GCCPUFraction", float64(stats.GCCPUFraction), hashServer.GenerateHash(models.Metric{ID: "GCCPUFraction", MType: "gauge", Delta: nil, Value: converter.Float64ToFloat64Pointer(stats.GCCPUFraction)}))
	memAgent.PutMetricsGauge("GCSys", float64(stats.GCSys), hashServer.GenerateHash(models.Metric{ID: "GCSys", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.GCSys)}))
	memAgent.PutMetricsGauge("HeapAlloc", float64(stats.HeapAlloc), hashServer.GenerateHash(models.Metric{ID: "HeapAlloc", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.HeapAlloc)}))
	memAgent.PutMetricsGauge("HeapIdle", float64(stats.HeapIdle), hashServer.GenerateHash(models.Metric{ID: "HeapIdle", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.HeapIdle)}))
	memAgent.PutMetricsGauge("HeapInuse", float64(stats.HeapInuse), hashServer.GenerateHash(models.Metric{ID: "HeapInuse", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.HeapInuse)}))
	memAgent.PutMetricsGauge("HeapObjects", float64(stats.HeapObjects), hashServer.GenerateHash(models.Metric{ID: "HeapObjects", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.HeapObjects)}))
	memAgent.PutMetricsGauge("HeapReleased", float64(stats.HeapReleased), hashServer.GenerateHash(models.Metric{ID: "HeapReleased", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.HeapReleased)}))
	memAgent.PutMetricsGauge("HeapSys", float64(stats.HeapSys), hashServer.GenerateHash(models.Metric{ID: "HeapSys", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.HeapSys)}))
	memAgent.PutMetricsGauge("LastGC", float64(stats.LastGC), hashServer.GenerateHash(models.Metric{ID: "LastGC", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.LastGC)}))
	memAgent.PutMetricsGauge("Lookups", float64(stats.Lookups), hashServer.GenerateHash(models.Metric{ID: "Lookups", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.Lookups)}))
	memAgent.PutMetricsGauge("MCacheInuse", float64(stats.MCacheInuse), hashServer.GenerateHash(models.Metric{ID: "MCacheInuse", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.MCacheInuse)}))
	memAgent.PutMetricsGauge("MCacheSys", float64(stats.MCacheSys), hashServer.GenerateHash(models.Metric{ID: "MCacheSys", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.MCacheSys)}))
	memAgent.PutMetricsGauge("MSpanInuse", float64(stats.MSpanInuse), hashServer.GenerateHash(models.Metric{ID: "MSpanInuse", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.MSpanInuse)}))
	memAgent.PutMetricsGauge("MSpanSys", float64(stats.MSpanSys), hashServer.GenerateHash(models.Metric{ID: "MSpanSys", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.MSpanSys)}))
	memAgent.PutMetricsGauge("Mallocs", float64(stats.Mallocs), hashServer.GenerateHash(models.Metric{ID: "Mallocs", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.Mallocs)}))
	memAgent.PutMetricsGauge("NextGC", float64(stats.NextGC), hashServer.GenerateHash(models.Metric{ID: "NextGC", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.NextGC)}))
	memAgent.PutMetricsGauge("NumForcedGC", float64(stats.NumForcedGC), hashServer.GenerateHash(models.Metric{ID: "NumForcedGC", MType: "gauge", Delta: nil, Value: converter.Uint32ToFloat64Pointer(stats.NumForcedGC)}))
	memAgent.PutMetricsGauge("NumGC", float64(stats.NumGC), hashServer.GenerateHash(models.Metric{ID: "NumGC", MType: "gauge", Delta: nil, Value: converter.Uint32ToFloat64Pointer(stats.NumGC)}))
	memAgent.PutMetricsGauge("OtherSys", float64(stats.OtherSys), hashServer.GenerateHash(models.Metric{ID: "OtherSys", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.OtherSys)}))
	memAgent.PutMetricsGauge("PauseTotalNs", float64(stats.PauseTotalNs), hashServer.GenerateHash(models.Metric{ID: "PauseTotalNs", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.PauseTotalNs)}))
	memAgent.PutMetricsGauge("StackInuse", float64(stats.StackInuse), hashServer.GenerateHash(models.Metric{ID: "StackInuse", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.StackInuse)}))
	memAgent.PutMetricsGauge("StackSys", float64(stats.StackSys), hashServer.GenerateHash(models.Metric{ID: "StackSys", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.StackSys)}))
	memAgent.PutMetricsGauge("Sys", float64(stats.Sys), hashServer.GenerateHash(models.Metric{ID: "Sys", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.Sys)}))
	memAgent.PutMetricsGauge("TotalAlloc", float64(stats.TotalAlloc), hashServer.GenerateHash(models.Metric{ID: "TotalAlloc", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(stats.TotalAlloc)}))

	randV := rand.Float64()
	memAgent.PutMetricsGauge("RandomValue", randV, hashServer.GenerateHash(models.Metric{ID: "RandomValue", MType: "gauge", Delta: nil, Value: converter.Float64ToFloat64Pointer(randV)}))

	pollCount, _, _ := memAgent.GetMetricsCount("PollCount")
	pollCount++
	memAgent.PutMetricsCount("PollCount", pollCount, hashServer.GenerateHash(models.Metric{ID: "PollCount", MType: "counter", Delta: converter.Int64ToInt64Pointer(pollCount), Value: nil}))
}

// Getting data using the psutil package
func DataFromRuntimeUsePsutil(memAgent *memstorage.MemStorage, v *mem.VirtualMemoryStat, hashServer *hashservicer.HashServer) {
	memAgent.PutMetricsGauge("TotalMemory", float64(v.Total), hashServer.GenerateHash(models.Metric{ID: "TotalMemory", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(v.Total)}))
	memAgent.PutMetricsGauge("FreeMemory", float64(v.Free), hashServer.GenerateHash(models.Metric{ID: "FreeMemory", MType: "gauge", Delta: nil, Value: converter.Uint64ToFloat64Pointer(v.Free)}))
	memAgent.PutMetricsGauge("CPUutilization1", float64(v.UsedPercent), hashServer.GenerateHash(models.Metric{ID: "CPUutilization1", MType: "gauge", Delta: nil, Value: converter.Float64ToFloat64Pointer(v.UsedPercent)}))
}
