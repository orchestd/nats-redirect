package config

import (
	"encoding/json"
	"os"
)

func LoadConfig(path string) (Config, error) {
	var cnf Config

	data, err := os.ReadFile(path)
	if err != nil {
		return cnf, err
	}

	err = json.Unmarshal(data, &cnf)

	return cnf, err
}
