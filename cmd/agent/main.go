package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/vasiliyantufev/go-advanced-devops/internal/api/agent"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/hashservicer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/api/helpers"
	"github.com/vasiliyantufev/go-advanced-devops/internal/config/configagent"
	"github.com/vasiliyantufev/go-advanced-devops/internal/models"
	"github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"

	log "github.com/sirupsen/logrus"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {

	//A := "192.168.0.0/16"
	//B := helpers.GetOutboundIP().String()
	//IPAddressAgent := net.ParseIP(B)
	//ipA, ipnetA, _ := net.ParseCIDR(A)
	//fmt.Println("Network address A: ", A)
	//fmt.Println("ipA              : ", ipA)
	//fmt.Println("ipnetA           : ", ipnetA)
	//fmt.Printf("\nDoes A (%s) contain: B (%s)?\n", ipnetA, IPAddressAgent)
	//if ipnetA.Contains(IPAddressAgent) {
	//	fmt.Println("yes")
	//} else {
	//	fmt.Println("no")
	//}
	//log.Fatal(IPAddressAgent)

	helpers.PrintInfo(buildVersion, buildDate, buildCommit)

	configAgent := configagent.NewConfigAgent()
	memAgent := memstorage.NewMemStorage()
	memAgentPsutil := memstorage.NewMemStorage()
	hashServer := hashservicer.NewHashServer(configAgent.Key)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	jobs := make(chan []models.Metric, configAgent.RateLimit)
	defer close(jobs)

	agent := agent.NewAgent(jobs, memAgent, memAgentPsutil, configAgent, hashServer)
	agent.StartWorkers(ctx, agent)

	<-ctx.Done()
	log.Println("agent shutdown on signal with:", ctx.Err())
}
