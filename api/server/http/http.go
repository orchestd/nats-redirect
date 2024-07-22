package http

import (
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/transport"
	"github.com/orchestd/nats-redirect/api/server/monolith"
)

func InitHandlers(router transport.IRouter, m monolith.NatsRedirectInterfaceInterface, c configuration.Config) {
}
