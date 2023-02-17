package main

import (
	"context"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

func main() {

	//config.SetConfigAgent()

	cfg := config.NewConfigAgent()
	mem := storage.NewMemStorage()
	hashServer := &config.HashServer{}

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()
	go PutMetrics(ctx, mem, cfg, hashServer)
	go SentMetrics(ctx, mem, cfg)
	<-ctx.Done()
	log.Println("agent shutdown on signal with:", ctx.Err())
}

func PutMetrics(ctx context.Context, mem *storage.MemStorage, cfg *config.ConfigAgent, hashServer *config.HashServer) {
	ticker := time.NewTicker(cfg.GetConfigPollIntervalAgent())
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
			app.DataFromRuntime(mem, cfg, stats, hashServer)
		}
	}
}

func SentMetrics(ctx context.Context, mem *storage.MemStorage, cfg *config.ConfigAgent) {
	// Create a Resty Client
	client := resty.New()
	urlPath := "http://" + cfg.GetConfigAddressAgent() + "/updates/"
	ticker := time.NewTicker(cfg.GetConfigReportIntervalAgent())
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
