package middlewares

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/errors"
)

// ValidURLParamsWithValueMetric - valid url params (type, name, value)
func ValidURLParamsWithValueMetric(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		typeMetrics := chi.URLParam(r, "type")
		nameMetrics := chi.URLParam(r, "name")
		valueMetrics := chi.URLParam(r, "value")

		if typeMetrics == "" {
			log.Error(errors.ErrTypeIsMissing)
			http.Error(w, errors.ErrTypeIsMissing.Error(), http.StatusBadRequest)
			return
		}
		if typeMetrics != "gauge" && typeMetrics != "counter" {
			log.Error(errors.ErrTypeIncorrect)
			http.Error(w, errors.ErrTypeIncorrect.Error(), http.StatusNotImplemented)
			return
		}
		if nameMetrics == "" {
			log.Error(errors.ErrNameIsMissing)
			http.Error(w, errors.ErrNameIsMissing.Error(), http.StatusBadRequest)
			return
		}
		if valueMetrics == "" {
			log.Error(errors.ErrValueIsMissing)
			http.Error(w, errors.ErrValueIsMissing.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
