package main

import (
	"github.com/orchestd/dependencybundler/bundler"
	"github.com/orchestd/dependencybundler/depBundler"
	"github.com/orchestd/nats-redirect/api/server/http"
)

type Configuration struct {
	depBundler.DependencyBundlerConfiguration
	ServersInSecret interface{} `json:"serversInSecret"`
	RulesFilePath   string      `json:"rulesFilePath"`
	//ServersInSecret interface{} `json:"serversInSecret"`
}

var appConf Configuration

func main() {
	bundler.CreateApplication(&appConf, http.InitHandlers, deps()...)

	//lgr := logger.New()

	//if cnf, err := rules.Load("test.json"); err != nil {
	//	lgr.Error("failed loading conf file %s", err.Error())
	//} else if clnt, err := client.New(lgr, cnf); err != nil {
	//	lgr.Error("failed setting new client %s", err.Error())
	//} else if gw, err := redirector.New(lgr, cnf, clnt); err != nil {
	//	lgr.Error("failed creating new redirector %s", err.Error())
	//} else {
	//	gw.Start()
	//}
}
