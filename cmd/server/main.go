package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"net/http"
)

func main() {

	storage.InitMap()

	r := chi.NewRouter()

	r.Get("/", app.IndexHandler)

	// http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>;
	r.Route("/update", func(r chi.Router) {
		for name, _ := range storage.MetricsGauge {
			r.Post("/gauge/"+name, app.MetricsGaugeHandler)
		}
		for name, _ := range storage.MetricsCounter {
			r.Post("/counter/"+name, app.MetricsCounterHandler)
		}
	})

	http.ListenAndServe(":8080", r)
}
