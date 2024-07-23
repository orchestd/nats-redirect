package natsredirector

import (
	"context"
	"github.com/orchestd/nats-redirect/domain/forwardingrules"
	"github.com/orchestd/nats-redirect/domain/natsconnection"
)

type Redirector interface {
	ConnectServers(ctx context.Context, serverConnections []natsconnection.ConnectionCredentials) error
	ListenAndForward(rules []forwardingrules.Rule) error
}
