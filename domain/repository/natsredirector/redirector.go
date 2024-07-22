package natsredirector

import (
	"context"
	"github.com/orchestd/nats-redirect/domain/forwardingrules"
	"github.com/orchestd/nats-redirect/domain/natsconnection"
)

type Redirector interface {
	ConnectServers(ctx context.Context, servers []natsconnection.ConnectionCredentials) error
	Forward(rules []forwardingrules.Rule) error
}
