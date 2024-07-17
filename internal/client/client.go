package client

import (
	"github.com/nats-io/nats.go"
	"github.com/orchestd/nats-redirect/internal/config"
	"github.com/orchestd/nats-redirect/logger"
	"time"
)

type Client struct {
	connections nConnections
	cnf         config.Config
	logger      *logger.Logger
}

type nConnections []*nats.Conn

func (c *Client) Close() {
	for _, conn := range c.connections {
		conn.Close()
	}
}

func (connections nConnections) findIndex(url string) (int, bool) {
	for i, conn := range connections {
		if conn.ConnectedUrl() == url {
			return i, true
		}
	}

	return 0, false
}

func (c *Client) Forward() error {
	for _, channel := range c.cnf.Forwards {
		if err := c.forward(channel); err != nil {
			c.logger.Error("shit %s", err.Error())
		}
	}
	return nil
}

func (c *Client) forward(forward config.Forward) error {
	if source, ok := c.connections.findIndex(forward.Source); !ok {
		return ErrSourceServerNotFound
	} else if target, ok := c.connections.findIndex(forward.Target); !ok {
		return ErrTargetServerNotFound
	} else {
		for _, method := range forward.GetMethods() {
			if err := c.listenAndForward(source, target, forward.RequestType, method); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Client) listenAndForward(source, target int, reqType config.RequestType, channel string) error {
	var handler nats.MsgHandler
	sourceUrl := c.connections[source].ConnectedUrl()
	targetUrl := c.connections[target].ConnectedUrl()

	//logPreFix := func(sourceUrl, targetUrl, channel string, reqType requestType) string {
	//
	//}

	switch reqType {
	case config.PUB:
		handler = func(msg *nats.Msg) {
			c.logger.Debug("source=%s target=%s channel=%s type=%s msg=%s", sourceUrl, targetUrl, channel, config.PUB, msg.Data)
			if err := c.connections[target].Publish(channel, msg.Data); err != nil {
				c.logger.Error("source-url=%s target-url=%s publish err %s", c.connections[source].ConnectedUrl(), c.connections[target].ConnectedUrl(), err.Error())
			}
		}
	case config.REQ:
		handler = func(msg *nats.Msg) {
			c.logger.Debug("source=%s target=%s channel=%s type=%s msg=%s", sourceUrl, targetUrl, channel, config.REQ, msg.Data)
			if resp, err := c.connections[target].Request(channel, msg.Data, 5*time.Second); err != nil {
				c.logger.Error("source-url=%s target-url=%s request err %s", c.connections[source].ConnectedUrl(), c.connections[target].ConnectedUrl(), err.Error())
			} else if err = msg.RespondMsg(resp); err != nil {
				c.logger.Error("source-url=%s target-url=%s response err %s", c.connections[source].ConnectedUrl(), c.connections[target].ConnectedUrl(), err.Error())
			}
		}
	default:
		return ErrRequestTypeNotSupported
	}

	_, err := c.connections[source].Subscribe(channel, handler)
	c.logger.Debug("source=%s target=%s channel=%s type=%s", sourceUrl, targetUrl, channel, string(reqType))
	return err
}

func New(log *logger.Logger, cnf config.Config) (config.Client, error) {
	var connections []*nats.Conn
	for _, server := range cnf.Servers {
		conn, err := nats.Connect(server.GetUrl(), server.GetConnection().(nats.Option))
		if err != nil {
			return nil, err
		}
		log.Info("now listening to server %s", server.Url)
		connections = append(connections, conn)
	}

	return &Client{connections: connections, cnf: cnf, logger: log}, nil
}
