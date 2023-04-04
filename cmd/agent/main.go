// module agent
package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/vasiliyantufev/go-advanced-devops/internal/app"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configagent"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage"

	log "github.com/sirupsen/logrus"
)

// agent main
func main() {

	configAgent := configagent.NewConfigAgent()
	memAgent := storage.NewMemStorage()
	memAgentPsutil := storage.NewMemStorage()
	hashServer := app.NewHashServer(configAgent.GetConfigKeyAgent())

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	jobs := make(chan []storage.JSONMetrics, configAgent.RateLimit)
	defer close(jobs)

	agent := app.NewAgent(jobs, memAgent, memAgentPsutil, configAgent, hashServer)
	agent.StartWorkers(ctx, agent)

	<-ctx.Done()
	log.Println("agent shutdown on signal with:", ctx.Err())
}
