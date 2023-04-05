// Package routerdevops - setting routes
package routerdevops

import (
	"net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/server/handlers"
	middlewaredevops "github.com/vasiliyantufev/go-advanced-devops/internal/api/server/middleware"
)

// Route - setting routes
func Route(s *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Compress(1, "application/json", "text/html"))
	r.Use(middlewaredevops.GzipMiddleware)

	r.Get("/", s.IndexHandler)
	r.Get("/ping", s.PingHandler)
	r.Route("/value", func(r chi.Router) {
		r.Get("/{type}/{name}", s.GetMetricURLParamsHandler)
		r.Post("/", s.GetValueMetricJSONHandler)
	})
	r.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", s.CreateMetricURLParamsHandler)
		r.Post("/", s.CreateMetricJSONHandler)
	})
	r.Route("/debug/pprof/", func(r chi.Router) {
		r.Get("/", pprof.Index)
		r.Get("/profile", pprof.Profile)
		r.Get("/cmdline", pprof.Cmdline)
		r.Get("/symbol", pprof.Symbol)
		r.Get("/trace", pprof.Trace)
		r.Get("/{cmd}", pprof.Index)
	})
	r.Post("/updates/", s.CreateMetricsJSONHandler)

	return r
}
