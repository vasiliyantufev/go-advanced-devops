package storage

var MetricsGauge = make(map[string]float64)
var MetricsCounter = make(map[string]int64)

func InitMap() {

	MetricsCounter["poll_count"] = 0
	MetricsCounter["random_value"] = 0

	MetricsGauge["alloc"] = 0
	MetricsGauge["buck_hash_sys"] = 0
	MetricsGauge["frees"] = 0
	MetricsGauge["gc_cpu_fraction"] = 0
	MetricsGauge["gc_sys"] = 0
	MetricsGauge["heap_alloc"] = 0
	MetricsGauge["heap_idle"] = 0
	MetricsGauge["heap_inuse"] = 0
	MetricsGauge["heap_objects"] = 0
	MetricsGauge["heap_released"] = 0
	MetricsGauge["heap_sys"] = 0
	MetricsGauge["last_gc"] = 0
	MetricsGauge["lookups"] = 0
	MetricsGauge["mcache_inuse"] = 0
	MetricsGauge["mcache_sys"] = 0
	MetricsGauge["mspan_inuse"] = 0
	MetricsGauge["mspan_sys"] = 0
	MetricsGauge["mallocs"] = 0
	MetricsGauge["next_gc"] = 0
	MetricsGauge["num_forced_gc"] = 0
	MetricsGauge["num_gc"] = 0
	MetricsGauge["other_sys"] = 0
	MetricsGauge["pause_total_ns"] = 0
	MetricsGauge["stack_inuse"] = 0
	MetricsGauge["stack_sys"] = 0
	MetricsGauge["sys"] = 0
	MetricsGauge["total_alloc"] = 0
}
