package client

//
//import (
//	"encoding/json"
//	"github.com/nats-io/nats.go"
//	"github.com/orchestd/nats-redirect/internal/config"
//	"github.com/orchestd/nats-redirect/logger"
//)
//
//type Config struct {
//	Servers  Servers  `json:"servers"`
//	Forwards Forwards `json:"forwards"`
//}
//
//func NewConfig() config.Config {
//	return Config{}
//}
//
//func (cnf Config) Unmarshal(data []byte) (config.Config, error) {
//	err := json.Unmarshal(data, &cnf)
//	return cnf, err
//}
//
//func (cnf Config) NewClient(log *logger.Logger) (config.Client, error) {
//	var connections []*nats.Conn
//	for _, server := range cnf.Servers {
//		conn, err := nats.Connect(server.GetUrl(), server.GetConnection().(nats.Option))
//		if err != nil {
//			return nil, err
//		}
//		log.Info("now listening to server %s", server.Url)
//		connections = append(connections, conn)
//	}
//
//	return &Client{connections: connections, cnf: cnf, logger: log}, nil
//}
