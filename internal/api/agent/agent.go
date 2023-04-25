// Package agent
package agent

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	runtime2 "github.com/vasiliyantufev/go-advanced-devops/internal/api/runtime"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configagent"
	"github.com/vasiliyantufev/go-advanced-devops/internal/models"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"

	log "github.com/sirupsen/logrus"
)

type Agent interface {
	StartWorkers(ctx context.Context, ai Agent)
	putMetricsWorker(ctx context.Context)
	putMetricsUsePsutilWorker(ctx context.Context)
	writeMetricsToChanWorker(ctx context.Context)
	sentMetricsWorker(ctx context.Context, url string, client *resty.Client)
	makePostRequest(client *resty.Client, j []models.Metric, url string)
}

type agent struct {
	jobs       chan []models.Metric
	mem        *memstorage.MemStorage
	psutil     *memstorage.MemStorage
	cfg        *configagent.ConfigAgent
	hashServer *hashservicer.HashServer
}

// Creates a new agent instance
func NewAgent(jobs chan []models.Metric, mem *memstorage.MemStorage, memPsutil *memstorage.MemStorage, cfg *configagent.ConfigAgent, hashServer *hashservicer.HashServer) *agent {
	return &agent{jobs: jobs, mem: mem, psutil: memPsutil, cfg: cfg, hashServer: hashServer}
}

func (a agent) StartWorkers(ctx context.Context, ai Agent) {
	// Create a Resty Client
	client := resty.New()
	var urlPath string

	if a.cfg.CryptoKey != "" {
		// Create a Go's http.Transport so we can set it in resty.
		caCert, err := ioutil.ReadFile(a.cfg.CryptoKey)
		if err != nil {
			log.Fatal(err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		transport := http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		}
		// Set the previous transport that we created, set the scheme
		client.SetTransport(&transport).SetScheme("https")
		urlPath = "https://" + a.cfg.Address + "/updates/"
	} else {
		// Set the scheme
		urlPath = "http://" + a.cfg.Address + "/updates/"
	}

	go ai.putMetricsWorker(ctx)
	go ai.putMetricsUsePsutilWorker(ctx)
	go ai.writeMetricsToChanWorker(ctx)

	for i := 0; i < a.cfg.RateLimit; i++ {
		go ai.sentMetricsWorker(ctx, urlPath, client)
	}
}

// Get metrics using runtime and write them to memory
func (a agent) putMetricsWorker(ctx context.Context) {
	ticker := time.NewTicker(a.cfg.PollInterval)
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
			runtime2.DataFromRuntime(a.mem, stats, a.hashServer)
		}
	}
}

// Gets metrics using psutil and write to memory
func (a agent) putMetricsUsePsutilWorker(ctx context.Context) {

	ticker := time.NewTicker(a.cfg.PollInterval)
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
			runtime2.DataFromRuntimeUsePsutil(a.psutil, v, a.hashServer)
		}
	}
}

// Writes metrics to a channel
func (a agent) writeMetricsToChanWorker(ctx context.Context) {
	ticker := time.NewTicker(a.cfg.ReportInterval)
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
func (a agent) sentMetricsWorker(ctx context.Context, url string, client *resty.Client) {
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

func (a agent) makePostRequest(client *resty.Client, j []models.Metric, url string) {
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(j).
		Post(url)
	if err != nil {
		log.Error(err)
		return
	}
	log.Println("Sent metrics success ", url)
}
