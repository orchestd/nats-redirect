package redirector

import (
	"github.com/orchestd/nats-redirect/internal/client"
	"github.com/orchestd/nats-redirect/internal/rules"
	"github.com/orchestd/nats-redirect/logger"
	"github.com/orchestd/nats-redirect/utils"
)

type Redirector struct {
	client *client.Client
	config rules.Config
	logger *logger.Logger
}

func New(logger *logger.Logger, cfg rules.Config, client *client.Client) (*Redirector, error) {
	return &Redirector{
		client: client,
		config: cfg,
		logger: logger,
	}, nil
}

func (r *Redirector) Start() {
	err := r.client.Forward()
	if err != nil {
		r.logger.Error("client failed forwarding err: %s", err.Error())
	}
	utils.GoForever()
}
