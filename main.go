//go:generate go run pkg/codegen/cleanup/main.go
//go:generate go run pkg/codegen/main.go

package main

import (
	"flag"

	"github.com/rancher/aks-operator/controller"
	aksv1 "github.com/rancher/aks-operator/pkg/generated/controllers/aks.cattle.io"
	"github.com/rancher/aks-operator/pkg/version"
	core3 "github.com/rancher/wrangler/v3/pkg/generated/controllers/core"
	"github.com/rancher/wrangler/v3/pkg/kubeconfig"
	"github.com/rancher/wrangler/v3/pkg/signals"
	"github.com/rancher/wrangler/v3/pkg/start"
	"github.com/sirupsen/logrus"
)

var (
	masterURL      string
	kubeconfigFile string
	debug          bool
)

func init() {
	flag.StringVar(&kubeconfigFile, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.BoolVar(&debug, "debug", false, "Variable to set log level to debug; default is false")
	flag.Parse()
}

func main() {
	// set up signals so we handle the first shutdown signal gracefully
	ctx := signals.SetupSignalContext()

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debugf("Loglevel set to [%v]", logrus.DebugLevel)
	}
	logrus.Infof("Starting aks-operator (version: %s, commit: %s)", version.Version, version.GitCommit)

	// This will load the kubeconfig file in a style the same as kubectl
	cfg, err := kubeconfig.GetNonInteractiveClientConfig(kubeconfigFile).ClientConfig()
	if err != nil {
		logrus.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	// core
	core, err := core3.NewFactoryFromConfig(cfg)
	if err != nil {
		logrus.Fatalf("Error building core factory: %s", err.Error())
	}

	// Generated sample controller
	aks, err := aksv1.NewFactoryFromConfig(cfg)
	if err != nil {
		logrus.Fatalf("Error building aks factory: %s", err.Error())
	}

	// The typical pattern is to build all your controller/clients then just pass to each handler
	// the bare minimum of what they need.  This will eventually help with writing tests.  So
	// don't pass in something like kubeClient, apps, or sample
	controller.Register(ctx,
		core.Core().V1().Secret(),
		aks.Aks().V1().AKSClusterConfig())

	// Start all the controllers
	if err := start.All(ctx, 2, aks, core); err != nil {
		logrus.Fatalf("Error starting: %s", err.Error())
	}

	<-ctx.Done()
}
