package main

import (
	"context"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {

	config.SetConfigAgent()
	mem := storage.NewMemStorage()

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()
	go PutMetrics(ctx, mem)
	go SentMetrics(ctx, mem)
	<-ctx.Done()
	log.Println("agent shutdown on signal with:", ctx.Err())
}

func PutMetrics(ctx context.Context, mem *storage.MemStorage) {
	ticker := time.NewTicker(config.GetConfigPollIntervalAgent())
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
			app.DataFromRuntime(mem, stats)
		}
	}
}

func SentMetrics(ctx context.Context, mem *storage.MemStorage) {
	// Create a Resty Client
	client := resty.New()
	urlPath := "http://" + config.GetConfigAddressAgent() + "/updates/"
	ticker := time.NewTicker(config.GetConfigReportIntervalAgent())
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Println("report ticker stopped by ctx")
			return
		case <-ticker.C:
			log.Info("Sent metrics")
			metrics := mem.GetAllMetrics()

			_, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(metrics).
				Post(urlPath)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
