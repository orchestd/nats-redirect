package main

import (
	"github.com/orchestd/dependencybundler/bundler"
	"github.com/orchestd/dependencybundler/depBundler"
	"github.com/orchestd/nats-redirect/api/server/http"
	"github.com/orchestd/nats-redirect/infrastucure/natsredirector"
)

type Configuration struct {
	depBundler.DependencyBundlerConfiguration
	natsredirector.Configuration
	RulesFilePath string `json:"rulesFilePath"`
}

var appConf Configuration

func main() {
	bundler.CreateApplication(&appConf, http.InitHandlers, deps()...)
}
