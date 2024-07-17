package config

import (
	"encoding/json"
	"os"
)

type Client interface {
	Forward() error
}

//type Config interface {
//	Unmarshal(data []byte) (Config, error)
//	NewClient(lgr *logger.Logger) (Client, error)
//}

func LoadConfig(path string) (Config, error) {
	var cnf Config

	data, err := os.ReadFile(path)
	if err != nil {
		return cnf, err
	}

	err = json.Unmarshal(data, &cnf)

	return cnf, err
}
