package main

import (
	"fmt"
	"github.com/tianshimoyi/my-sheduler/pkg"
	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	command := app.NewSchedulerCommand(app.WithPlugin(pkg.Name, pkg.New))
	logs.InitLogs()
	defer logs.FlushLogs()
	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
