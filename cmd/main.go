package main

import (
	"github.com/orchestd/nats-redirect/internal/client"
	"github.com/orchestd/nats-redirect/internal/config"
	"github.com/orchestd/nats-redirect/internal/redirector"
	"github.com/orchestd/nats-redirect/logger"
)

func main() {
	// logger
	lgr := logger.New()

	// load file
	if cnf, err := config.LoadConfig("test.json"); err != nil {
		lgr.Error("niii %s", err.Error())
	} else if clnt, err := client.New(lgr, cnf); err != nil {
		lgr.Error("nii %s", err.Error())
	} else if gw, err := redirector.New(lgr, cnf, clnt); err != nil {
		lgr.Error("ni %s", err.Error())
	} else {
		gw.Start()
	}
}
