package main

import (
	"context"
	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {

	configAgent := config.NewConfigAgent()
	memAgent := storage.NewMemStorage()
	memAgentPsutil := storage.NewMemStorage()
	hashServer := app.NewHashServer(configAgent.GetConfigKeyAgent())

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	jobs := make(chan []storage.JSONMetrics, configAgent.RateLimit)
	defer close(jobs)

	agent := app.NewAgent(jobs, memAgent, memAgentPsutil, configAgent, hashServer)
	agent.StartWorkers(ctx)

	<-ctx.Done()
	log.Println("agent shutdown on signal with:", ctx.Err())
}
