package usercase

//func CreateGaugeMetric(nameMetrics string, valueMetrics string, s r.Handler) (string, error) {
//	val, err := strconv.ParseFloat(string(valueMetrics), 64)
//	if err != nil {
//		return "", err
//	}
//	hashServer := s.HashServer.GenerateHash(models.Metric{ID: nameMetrics, MType: "gauge", Delta: nil, Value: converter.Float64ToFloat64Pointer(val)})
//	s.MemStorage.PutMetricsGauge(nameMetrics, val, hashServer)
//	resp := "Request completed successfully " + nameMetrics + "=" + fmt.Sprint(val)
//	return resp, nil
//}
