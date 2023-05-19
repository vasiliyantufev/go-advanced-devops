package middlewares

import (
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/errors"
)

// TrustedSubnetMiddleware - checks the IP address in a trusted network
func TrustedSubnetMiddleware(subnet *net.IPNet) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			IPAddressAgent := net.ParseIP(r.Header.Get("X-Real-IP"))
			if !subnet.Contains(IPAddressAgent) {
				log.Error(errors.ErrAddressNotTrustedSubnet)
				http.Error(w, errors.ErrAddressNotTrustedSubnet.Error(), http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
