package main

import (
	"context"
	"github.com/go-resty/resty/v2"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"

	"github.com/shirou/gopsutil/v3/mem"
	log "github.com/sirupsen/logrus"
)

func main() {

	configAgent := config.NewConfigAgent()
	memAgent := storage.NewMemStorage()
	memAgentPsutil := storage.NewMemStorage()
	//hashServer := &app.HashServer{}
	hashServer := app.NewHashServer(configAgent.GetConfigKeyAgent())
	urlPath := "http://" + configAgent.GetConfigAddressAgent() + "/updates/"

	//log.Fatal(hashServer.Key)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	jobs := make(chan []storage.JSONMetrics, configAgent.RateLimit)
	defer close(jobs)

	go PutMetrics(ctx, memAgent, configAgent, hashServer)
	go PutMetricsUsePsutil(ctx, memAgentPsutil, configAgent, hashServer)
	go WrireMetricsToChan(ctx, jobs, memAgent, memAgentPsutil, configAgent)

	for i := 0; i < configAgent.RateLimit; i++ {
		go SentMetrics(ctx, jobs, urlPath)
	}

	<-ctx.Done()
	log.Println("agent shutdown on signal with:", ctx.Err())
}

func PutMetrics(ctx context.Context, memAgent *storage.MemStorage, cfg *config.ConfigAgent, hashServer *app.HashServer) {

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
			app.DataFromRuntime(memAgent, stats, hashServer)
		}
	}
}

func PutMetricsUsePsutil(ctx context.Context, memAgent *storage.MemStorage, cfg *config.ConfigAgent, hashServer *app.HashServer) {

	ticker := time.NewTicker(cfg.GetConfigPollIntervalAgent())
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Println("poll psutil ticker stopped by ctx")
			return
		case <-ticker.C:
			log.Info("Put metrics Psutil")
			v, err := mem.VirtualMemory()
			if err != nil {
				log.Error(err)
			}
			app.DataFromRuntimeUsePsutil(memAgent, v, hashServer)
		}
	}
}

func WrireMetricsToChan(ctx context.Context, jobs chan []storage.JSONMetrics, agent *storage.MemStorage, psutil *storage.MemStorage, cfg *config.ConfigAgent) {

	ticker := time.NewTicker(cfg.GetConfigReportIntervalAgent())
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Println("Read ticker stopped by ctx")
			return
		case <-ticker.C: // Запись в канал
			log.Info("Write metrics to chan")
			jobs <- agent.GetAllMetrics()
			jobs <- psutil.GetAllMetrics()
		}
	}
}

func SentMetrics(ctx context.Context, jobs chan []storage.JSONMetrics, url string) {

	client := resty.New()
	for {
		select {
		case <-ctx.Done():
			log.Println("Sent stopped by ctx")
			return
		case j := <-jobs:
			makePostRequest(client, j, url)
		}
	}
}

func makePostRequest(client *resty.Client, j []storage.JSONMetrics, url string) {

	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(j).
		Post(url)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Sent metrics success ", url)
}
