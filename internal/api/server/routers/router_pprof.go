package routers

import (
	"net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/server/handlers"
	middlewaredevops "github.com/vasiliyantufev/go-advanced-devops/internal/api/server/middlewares"
)

// RoutePProf pprof - setting pprof routes
func RoutePProf(s *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewaredevops.GzipMiddleware)
	r.Route("/debug/pprof/", func(r chi.Router) {
		r.Get("/", pprof.Index)
		r.Get("/profile", pprof.Profile)
		r.Get("/cmdline", pprof.Cmdline)
		r.Get("/symbol", pprof.Symbol)
		r.Get("/trace", pprof.Trace)
		r.Get("/{cmd}", pprof.Index)
	})

	return r
}
