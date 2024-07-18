package main

import (
	"github.com/orchestd/dependencybundler/bundler"
)

type Configuration struct {
	//credentialsConfiguration.CredentialsConfiguration
	//envConfiguration.EnvConfiguration
	//monitoringConfiguration.MonitoringConfiguration
	//logConfiguration.LogConfiguration
}

var appConf Configuration

func main() {

	bundler.CreateApplication(&appConf, http.InitHandlers, deps()...)

	//app := fx.New(
	//	bundler.CredentialsFxOption(),
	//	bundler.ConfigFxOption(confStruct),
	//	bundler.LoggerFxOption(),
	//	bundler.TransportFxOption(monolithConstructor...),
	//	bundler.CacheTraceMiddlewareOption(),
	//	bundler.TracerFxOption(),
	//	bundler.DebugFxOption(),
	//	bundler.MonitoringFxOption(),
	//)
	//
	//app.Run()

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
