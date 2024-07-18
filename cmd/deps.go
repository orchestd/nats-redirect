package main

import (
	"github.com/orchestd/dependencybundler/bundler"
	"github.com/orchestd/dependencybundler/depBundler/middlewares/cors"
	"github.com/orchestd/nats-redirect/application/defaultApp"
	"github.com/orchestd/nats-redirect/infrastucure/forwardingrules"
	"go.uber.org/fx"
)

func deps() []interface{} {
	return append(internalDeps(), externalDeps()...)
}

func internalDeps() []interface{} {
	return []interface{}{
		defaultApp.NewNatsRedirectApp,
		forwardingrules.New,
		//messaging.NewMessagingRepo,
	}
}

func externalDeps() []interface{} {
	return []interface{}{
		fx.Annotated{Group: bundler.RouterInterceptors, Target: cors.CorsMiddleware},
		//natsio.NewNatsServiceWithBasicAuth,
	}
}
