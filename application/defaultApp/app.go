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

const rulesFilePathKey = "rulesFilePath" // todo remove from here
const serversInSecretKey = "serversInSecret"

func NewNatsRedirectApp(logger log.Logger, conf configuration.Config, redirector natsredirector.Redirector, credentials credentials.CredentialsGetter, reader reader.Reader) application.NewNatsRedirectApp {
	app := &natsRedirectApp{logger: logger, conf: conf, redirector: redirector}

	ctx := context.Background()
	var serverConns []natsconnection.ConnectionCredentials
	var rules []forwardingrules.Rule
	//natsUser := credentials.GetCredentials().NatsUser

	if err := conf.Get(serversInSecretKey).Unmarshal(&serverConns); err != nil {
		panic(err)
	} else if err = redirector.ConnectServers(ctx, serverConns); err != nil {
		panic(err)
	} else if rulesFilePath, err := conf.Get(rulesFilePathKey).String(); err != nil {
		panic(err)
	} else if err = reader.ReadFile(rulesFilePath, &rules); err != nil {
		panic(err)
	} else if err = redirector.Forward(rules); err != nil {
		panic(err)
	} else {
		logger.Info(ctx, "hooray")
	}

	return app
}
