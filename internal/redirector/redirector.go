package redirector

import (
	"github.com/orchestd/nats-redirect/internal/config"
	"github.com/orchestd/nats-redirect/logger"
	"github.com/orchestd/nats-redirect/utils"
)

type Redirector struct {
	client config.Client
	config config.Config
	logger *logger.Logger
}

func New(logger *logger.Logger, cfg config.Config, client config.Client) (*Redirector, error) {
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
