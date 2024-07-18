package defaultApp

import (
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/log"
)

type natsRedirectApp struct {
	logger log.Logger
	conf   configuration.Config
}

func NewNatsRedirectApp(logger log.Logger, conf configuration.Config) interface{} {
	return &natsRedirectApp{logger: logger, conf: conf}
}
