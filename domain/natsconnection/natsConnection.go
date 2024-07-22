package natsconnection

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/nats-io/nats.go"
	"golang.org/x/exp/maps"
	"slices"
)

type ConnectionCredentials struct {
	Url         string
	Credentials Credentials
}

type Credentials interface {
	ConnectOption() nats.Option
}

var userPassValidation = []string{"username", "password"}
var jwtValidation = []string{"jwt", "seed"}

type UserPassConnData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtConnData struct {
	Jwt  string `json:"jwt"`
	Seed string `json:"seed"`
}

func (u UserPassConnData) ConnectOption() nats.Option {
	return nats.UserInfo(u.Username, u.Password)
}

func (j JwtConnData) ConnectOption() nats.Option {
	return nats.UserJWTAndSeed(j.Jwt, j.Seed)
}

func (s *ConnectionCredentials) UnmarshalJSON(data []byte) error {
	type Alias ConnectionCredentials
	aux := &struct {
		Connection interface{} `json:"connection"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Assert the type of Connection and convert it to the desired type
	switch v := aux.Connection.(type) {
	case map[string]interface{}:

		if ok := validateConnType(v, userPassValidation); ok {
			var connection UserPassConnData
			err := mapstructure.Decode(v, &connection)
			if err != nil {
				return fmt.Errorf("failed decoding map into struct UserPassConnData: %s", err.Error())
			}
			s.Credentials = connection
		} else if ok = validateConnType(v, jwtValidation); ok {
			var connection JwtConnData
			err := mapstructure.Decode(v, &connection)
			if err != nil {
				return fmt.Errorf("failed decoding map into struct JwtConnData: %s", err.Error())
			}
			s.Credentials = connection
		} else {
			return fmt.Errorf("fields are missing for any connection type: %T", maps.Keys(v))
		}
	default:
		return fmt.Errorf("unexpected type for connection: %T", aux.Connection)
	}

	return nil
}

func validateConnType(m map[string]interface{}, fields []string) bool {
	var validationFailed bool

	for k := range m {
		if !slices.Contains(fields, k) {
			validationFailed = true
			break
		}
	}

	return !validationFailed
}
