package config

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type HashServicer interface {
	GetHashServer(mid string, mtype string, delta int64, value float64, key string) string
}
type HashServer struct {
}

func (hs HashServer) GetHashServer(mid string, mtype string, delta int64, value float64, key string) string {

	var data string
	switch mtype {
	case "counter":
		data = fmt.Sprintf("%s:%s:%d", mid, mtype, delta)
	case "gauge":
		data = fmt.Sprintf("%s:%s:%f", mid, mtype, value)
	}

	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
