package natsredirector

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/log"
	"github.com/orchestd/nats-redirect/domain/forwardingrules"
	"github.com/orchestd/nats-redirect/domain/natsconnection"
	"github.com/orchestd/nats-redirect/domain/repository/natsredirector"
	"time"
)

type natsRedirector struct {
	connections connections
	conf        configuration.Config
	logger      log.Logger
}

const (
	PUB  = "publish"
	REQ  = "request"
	TEST = "test"
)

type connections []*nats.Conn

func NewMessagingRedirector(logger log.Logger, conf configuration.Config) natsredirector.Redirector {
	return &natsRedirector{connections: make(connections, 0), logger: logger, conf: conf}
}

func (m *natsRedirector) ConnectServers(ctx context.Context, serverConnections []natsconnection.ConnectionCredentials) error {
	var conns []*nats.Conn
	for _, server := range serverConnections {
		if conn, err := nats.Connect(server.Url, server.Credentials.ConnectOption()); err != nil {
			m.logger.Warn(ctx, "failed connecting to server %s", server.Url)
			return err
		} else {
			m.logger.Info(ctx, "connected to server %s", server.Url)
			conns = append(conns, conn)
		}
	}

	m.connections = conns
	return nil
}

func (m *natsRedirector) ListenAndForward(rules []forwardingrules.Rule) error {
	for _, forwardRule := range rules {
		ctx := context.Background()
		if err := m.forward(ctx, forwardRule); err != nil {
			m.logger.Error(ctx, "unable to forward %+v, err: %s", forwardRule, err.Error())
		}
	}

	return nil
}

func (m *natsRedirector) forward(ctx context.Context, forwardRule forwardingrules.Rule) error {
	if source, ok := m.connections.findIndex(forwardRule.Source); !ok {
		return ErrSourceServerNotFound
	} else if target, ok := m.connections.findIndex(forwardRule.Target); !ok {
		return ErrTargetServerNotFound
	} else {
		for _, subject := range forwardRule.Subjects {
			if err := m.listenAndForward(ctx, source, target, forwardRule.Type, subject); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *natsRedirector) listenAndForward(ctx context.Context, source, target int, reqType, subject string) error {
	var handler nats.MsgHandler

	sourceUrl := m.connections[source].ConnectedUrl()
	targetUrl := m.connections[target].ConnectedUrl()

	logPrefix := func(msgSubject string) string {
		return fmt.Sprintf("source=%s target=%s subject=%s type=%s ",
			sourceUrl, targetUrl, msgSubject, reqType)
	}

	switch reqType {
	case PUB:
		handler = func(msg *nats.Msg) {
			m.logger.Debug(ctx, logPrefix(msg.Subject)+"msg=%s", msg.Data)
			if err := m.connections[target].Publish(msg.Subject, msg.Data); err != nil {
				m.logger.Error(ctx, logPrefix(msg.Subject)+"publish err: %s", err.Error())
				// todo need to return err somehow
			}
		}
	case REQ:
		handler = func(msg *nats.Msg) {
			m.logger.Debug(ctx, logPrefix(msg.Subject)+"msg=%s", msg.Data)
			if resp, err := m.connections[target].Request(msg.Subject, msg.Data, 5*time.Second); err != nil {
				m.logger.Error(ctx, logPrefix(msg.Subject)+"request err: %s", err.Error())
				// todo need to return err somehow
			} else if err = msg.RespondMsg(resp); err != nil {
				m.logger.Error(ctx, logPrefix(msg.Subject)+"reply err: %s", err.Error())
			}
		}
	case TEST:
		handler = func(msg *nats.Msg) {
			m.logger.Debug(ctx, logPrefix(msg.Subject)+"msg=%s", msg.Data)
		}
	default:
		return ErrRequestTypeNotSupported
	}

	_, err := m.connections[source].Subscribe(subject, handler)
	m.logger.Debug(ctx, logPrefix(subject))
	return err
}

func (connections connections) findIndex(url string) (int, bool) {
	for i, conn := range connections {
		if conn.ConnectedUrl() == url {
			return i, true
		}
	}

	return 0, false
}
