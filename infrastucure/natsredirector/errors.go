package natsredirector

import "errors"

var (
	ErrForwardTypeDoesNotMatch = errors.New("_nats: forward type does not match")
	ErrSourceServerNotFound    = errors.New("_nats: source server not found")
	ErrTargetServerNotFound    = errors.New("_nats: target server not found")
	ErrRequestTypeNotSupported = errors.New("_nats: request type not supported")
)
