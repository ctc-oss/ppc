package servers

import (
	"github.com/jw3/ppc/common"
	"os"
)

type ServerConfig struct {
	ClientID          string
	BrokerURI         string
	AppPrefix         string
	EventChannelId    string
	FunctionChannelId string
}

func NewServerConfiguration() *ServerConfig {
	broker := envOr(common.EnvVarBrokerUri, "localhost:1883")
	appPrefix := envOr(common.EnvVarAppPrefix, "xr")
	funcChannel := envOr(common.EnvVarFunctionChannel, "F")
	eventChannel := envOr(common.EnvVarEventChannel, "E")

	sc := &ServerConfig{
		ClientID:          "ppc",
		BrokerURI:         broker,
		AppPrefix:         appPrefix,
		EventChannelId:    eventChannel,
		FunctionChannelId: funcChannel,
	}
	return sc
}

func envOr(key, or string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		v = or
	}
	return v
}
