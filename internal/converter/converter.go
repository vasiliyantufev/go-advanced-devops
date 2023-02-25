package converter

func Uint64ToFloat64Pointer(met uint64) *float64 {
	val := float64(met)
	return &val
}

func Uint32ToFloat64Pointer(met uint32) *float64 {
	val := float64(met)
	return &val
}

func Float64ToFloat64Pointer(val float64) *float64 {
	return &val
}

func Int64ToInt64Pointer(val int64) *int64 {
	return &val
}
