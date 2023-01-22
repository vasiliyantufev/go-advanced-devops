package main

import (
	"context"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var MemAgent = storage.NewMemStorage()

func main() {

	storage.SetConfigAgent()
	cfg := storage.GetConfigAgent()
	//cfg := storage.GetConfigEnv()

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()
	go PutMetrics(ctx, cfg)
	go SentMetrics(ctx, cfg)
	<-ctx.Done()
	log.Println("agent shutdown on signal with:", ctx.Err())
}

func PutMetrics(ctx context.Context, config storage.Config) {
	ticker := time.NewTicker(config.PollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Println("poll ticker stopped by ctx")
			return
		case <-ticker.C:
			log.Info("Put metrics")
			stats := new(runtime.MemStats)
			runtime.ReadMemStats(stats)
			app.DataFromRuntime(MemAgent, stats)
		}
	}
}

func SentMetrics(ctx context.Context, config storage.Config) {

	// Create a Resty Client
	client := resty.New()
	urlPath := "http://" + config.Address + "/update/"
	ticker := time.NewTicker(config.ReportInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Println("report ticker stopped by ctx")
			return
		case <-ticker.C:
			log.Info("Sent metrics")
			metrics := MemAgent.GetAllMetrics()
			for _, metric := range metrics {
				_, err := client.R().
					SetHeader("Content-Type", "application/json").
					SetBody(metric).
					Post(urlPath)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}
}
