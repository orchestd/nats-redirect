package defaultApp

import (
	"context"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/credentials"
	"github.com/orchestd/dependencybundler/interfaces/log"
	"github.com/orchestd/nats-redirect/domain/application"
	"github.com/orchestd/nats-redirect/domain/forwardingrules"
	"github.com/orchestd/nats-redirect/domain/natsconnection"
	"github.com/orchestd/nats-redirect/domain/repository/natsredirector"
	"github.com/orchestd/nats-redirect/domain/repository/reader"
)

type natsRedirectApp struct {
	logger     log.Logger
	conf       configuration.Config
	redirector natsredirector.Redirector
}

func NewNatsRedirectApp(logger log.Logger, conf configuration.Config, redirector natsredirector.Redirector, credentials credentials.CredentialsGetter, reader reader.Reader) application.NewNatsRedirectApp {
	ctx := context.Background()
	var serverConnections []natsconnection.ConnectionCredentials
	var rules []forwardingrules.Rule

	if err := credentials.GetCredentials().GetNatsServerConnections(&serverConnections); err != nil {
		panic(err)
	} else if rulesFilePath, err := conf.Get("rulesFilePath").String(); err != nil {
		panic(err)
	} else if err = reader.ReadFile(rulesFilePath, &rules); err != nil {
		panic(err)
	} else if err = redirector.ConnectServers(ctx, serverConnections); err != nil {
		panic(err)
	} else if err = redirector.ListenAndForward(rules); err != nil {
		panic(err)
	}

	return &natsRedirectApp{logger: logger, conf: conf, redirector: redirector}
}
