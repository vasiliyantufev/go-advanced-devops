// Package routers - setting routes
package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/server/handlers"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/server/middlewares"
)

// Route - setting service routes
func Route(s *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()

	//r.Use(middleware.Compress(1, "application/json", "text/html"))
	r.Use(middlewares.GzipMiddleware)

	r.Get("/", s.IndexHandler)
	r.Get("/ping", s.PingHandler)
	r.Route("/value", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middlewares.ValidURLParamsMetricMiddleware)
			r.Get("/{type}/{name}", s.GetMetricURLParamsHandler)
		})
		r.Post("/", s.GetValueMetricJSONHandler)
	})
	r.Route("/update", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middlewares.ValidURLParamsWithValueMetric)
			r.Post("/{type}/{name}/{value}", s.CreateMetricURLParamsHandler)
		})
		r.Post("/", s.CreateMetricJSONHandler)
	})
	r.Post("/updates/", s.CreateMetricsJSONHandler)

	return r
}
