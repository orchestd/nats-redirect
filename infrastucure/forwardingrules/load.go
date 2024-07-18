package forwardingrules

import (
	"encoding/json"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/nats-redirect/domain/repository/forwardingrules"
	"os"
)

const rulesFilePathKey = "rulesFilePath"

func New(conf configuration.Config) forwardingrules.Rules {
	var roles Rules

	if rulesFilePath, err := conf.Get(rulesFilePathKey).String(); err != nil {
		panic(err)
	} else if data, err := os.ReadFile(rulesFilePath); err != nil {
		panic(err)
	} else if err = json.Unmarshal(data, &roles); err != nil {
		panic(err)
	} else {
		return roles
	}
}
