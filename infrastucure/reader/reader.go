package reader

import (
	"encoding/json"
	"github.com/orchestd/nats-redirect/domain/repository/reader"
	"os"
)

type osReader struct {
}

func NewReader() reader.Reader {
	return osReader{}
}

func (o osReader) ReadFile(path string, target interface{}) error {
	if data, err := os.ReadFile(path); err != nil {
		return err
	} else if err = json.Unmarshal(data, target); err != nil {
		return err
	} else {
		return nil
	}
}
