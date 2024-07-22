package main

import (
	"github.com/orchestd/dependencybundler/bundler"
	"github.com/orchestd/dependencybundler/depBundler/middlewares/cors"
	"github.com/orchestd/nats-redirect/api/server/monolith"
	"github.com/orchestd/nats-redirect/application/defaultApp"
	"github.com/orchestd/nats-redirect/infrastucure/natsredirector"
	"github.com/orchestd/nats-redirect/infrastucure/reader"
	"go.uber.org/fx"
)

func deps() []interface{} {
	return append(internalDeps(), externalDeps()...)
}

func internalDeps() []interface{} {
	return []interface{}{
		defaultApp.NewNatsRedirectApp,
		monolith.NewNatsRedirectInterfaceInterface,
		natsredirector.NewMessagingRedirector,
		reader.NewReader,
	}
}

func externalDeps() []interface{} {
	return []interface{}{
		fx.Annotated{Group: bundler.RouterInterceptors, Target: cors.CorsMiddleware},
		//natsio.NewNatsServiceWithBasicAuth,
	}
}
