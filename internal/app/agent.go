// Module agent
package app

import (
	"context"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configagent"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"

	log "github.com/sirupsen/logrus"
)

type Agent interface {
	StartWorkers(ctx context.Context, ai Agent)
	putMetricsWorker(ctx context.Context)
	putMetricsUsePsutilWorker(ctx context.Context)
	writeMetricsToChanWorker(ctx context.Context)
	sentMetricsWorker(ctx context.Context, url string)
	makePostRequest(client *resty.Client, j []storage.JSONMetrics, url string)
}

type agent struct {
	jobs       chan []storage.JSONMetrics
	mem        *storage.MemStorage
	psutil     *storage.MemStorage
	cfg        *configagent.ConfigAgent
	hashServer *HashServer
}

// Creates a new agent instance
func NewAgent(jobs chan []storage.JSONMetrics, mem *storage.MemStorage, memPsutil *storage.MemStorage, cfg *configagent.ConfigAgent, hashServer *HashServer) *agent {
	return &agent{jobs: jobs, mem: mem, psutil: memPsutil, cfg: cfg, hashServer: hashServer}
}

func (a agent) StartWorkers(ctx context.Context, ai Agent) {

	urlPath := "http://" + a.cfg.GetConfigAddressAgent() + "/updates/"

	go ai.putMetricsWorker(ctx)
	go ai.putMetricsUsePsutilWorker(ctx)
	go ai.writeMetricsToChanWorker(ctx)

	for i := 0; i < a.cfg.RateLimit; i++ {
		go ai.sentMetricsWorker(ctx, urlPath)
	}
}

// Get metrics using runtime and write them to memory
func (a agent) putMetricsWorker(ctx context.Context) {

	ticker := time.NewTicker(a.cfg.GetConfigPollIntervalAgent())
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
			DataFromRuntime(a.mem, stats, a.hashServer)
		}
	}
}

// Gets metrics using psutil and write to memory
func (a agent) putMetricsUsePsutilWorker(ctx context.Context) {

	ticker := time.NewTicker(a.cfg.GetConfigPollIntervalAgent())
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
			DataFromRuntimeUsePsutil(a.psutil, v, a.hashServer)
		}
	}
}

// Writes metrics to a channel
func (a agent) writeMetricsToChanWorker(ctx context.Context) {

	ticker := time.NewTicker(a.cfg.GetConfigReportIntervalAgent())
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Println("Read ticker stopped by ctx")
			return
		case <-ticker.C: // Writes
			log.Info("Write metrics to chan")
			a.jobs <- a.mem.GetAllMetrics()
			a.jobs <- a.psutil.GetAllMetrics()
		}
	}
}

// Listens to the channel, if the metrics have arrived, forms a request and sends it to the server
func (a agent) sentMetricsWorker(ctx context.Context, url string) {

	client := resty.New()
	for {
		select {
		case <-ctx.Done():
			log.Println("Sent stopped by ctx")
			return
		case j := <-a.jobs:
			a.makePostRequest(client, j, url)
		}
	}
}

func (a agent) makePostRequest(client *resty.Client, j []storage.JSONMetrics, url string) {

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
