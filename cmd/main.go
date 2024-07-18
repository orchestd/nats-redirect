package main

import (
	"github.com/orchestd/nats-redirect/internal/client"
	"github.com/orchestd/nats-redirect/internal/config"
	"github.com/orchestd/nats-redirect/internal/redirector"
	"github.com/orchestd/nats-redirect/logger"
)

func main() {
	lgr := logger.New()

	if cnf, err := config.LoadConfig("test.json"); err != nil {
		lgr.Error("failed loading conf file %s", err.Error())
	} else if clnt, err := client.New(lgr, cnf); err != nil {
		lgr.Error("failed setting new client %s", err.Error())
	} else if gw, err := redirector.New(lgr, cnf, clnt); err != nil {
		lgr.Error("failed creating new redirector %s", err.Error())
	} else {
		gw.Start()
	}
}
