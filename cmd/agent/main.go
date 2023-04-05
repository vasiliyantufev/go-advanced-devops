package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/vasiliyantufev/go-advanced-devops/internal/api/agent"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configagent"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"

	log "github.com/sirupsen/logrus"
)

func main() {
	configAgent := configagent.NewConfigAgent()
	memAgent := storage.NewMemStorage()
	memAgentPsutil := storage.NewMemStorage()
	hashServer := hashservicer.NewHashServer(configAgent.GetConfigKeyAgent())

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	jobs := make(chan []storage.JSONMetrics, configAgent.RateLimit)
	defer close(jobs)

	agent := agent.NewAgent(jobs, memAgent, memAgentPsutil, configAgent, hashServer)
	agent.StartWorkers(ctx, agent)

	<-ctx.Done()
	log.Println("agent shutdown on signal with:", ctx.Err())
}
