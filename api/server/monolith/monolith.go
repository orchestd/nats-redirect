/* This file is auto-generated and should not be modified */

package monolith

import (
	"github.com/orchestd/dependencybundler/interfaces/validations"
	"github.com/orchestd/nats-redirect/domain/application"
)

type NatsRedirectInterfaceInterface struct {
	natsRedirectApp application.NewNatsRedirectApp
	validation      validations.Validations
}

func NewNatsRedirectInterfaceInterface(natsRedirectApp application.NewNatsRedirectApp, validation validations.Validations) NatsRedirectInterfaceInterface {
	return NatsRedirectInterfaceInterface{natsRedirectApp: natsRedirectApp, validation: validation}
}
