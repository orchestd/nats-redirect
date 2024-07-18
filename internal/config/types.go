package config

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/nats-io/nats.go"
	"golang.org/x/exp/maps"
	"slices"
	"strings"
)

type Config struct {
	Servers  Servers  `json:"servers"`
	Forwards Forwards `json:"forwards"`
}

type Forwards []Forward

type Forward struct {
	RequestType RequestType `json:"requestType"`
	Method      string      `json:"method"`
	Source      string      `json:"source"`
	Target      string      `json:"target"`
	Identifiers []string    `json:"identifiers"`
}

type RequestType string

const (
	PUB RequestType = "publish"
	REQ             = "request"
)

type Servers []ServerConnectionDetails

type ServerConnectionDetails struct {
	Url        string      `json:"url"`
	Connection ConnectionI `json:"connection"`
}

type ConnectionI interface {
	GetOption() interface{}
}

func (s *ServerConnectionDetails) GetUrl() string {
	return s.Url
}

func (s *ServerConnectionDetails) GetConnection() interface{} {
	return s.Connection.GetOption()
}

func (s *ServerConnectionDetails) UnmarshalJSON(data []byte) error {
	type Alias ServerConnectionDetails
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
		var userPassValidation = []string{"username", "password"}
		var jwtValidation = []string{"jwt", "seed"}

		if ok := validateConnType(v, userPassValidation); ok {
			var connection UserPassConnData
			err := mapstructure.Decode(v, &connection)
			if err != nil {
				return fmt.Errorf("failed decoding map into struct UserPassConnData: %s", err.Error())
			}
			s.Connection = connection
		} else if ok = validateConnType(v, jwtValidation); ok {
			var connection JwtConnData
			err := mapstructure.Decode(v, &connection)
			if err != nil {
				return fmt.Errorf("failed decoding map into struct JwtConnData: %s", err.Error())
			}
			s.Connection = connection
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

type UserPassConnData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtConnData struct {
	Jwt  string `json:"jwt"`
	Seed string `json:"seed"`
}

func (u UserPassConnData) GetOption() interface{} {
	return nats.UserInfo(u.Username, u.Password)
}

func (j JwtConnData) GetOption() interface{} {
	return nats.UserJWTAndSeed(j.Jwt, j.Seed)
}

func (m Forward) GetMethods() []string {
	methods := make([]string, len(m.Identifiers))
	for i := range methods {
		methods[i] = strings.ReplaceAll(m.Method, "{{id}}", m.Identifiers[i])
	}
	return methods
}
