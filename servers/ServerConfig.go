package servers

import (
	"github.com/jw3/ppc/common"
	"os"
)

type ServerConfig struct {
	ClientID       string
	BrokerURI      string
	EventPrefix    string
	FunctionPrefix string
}

func NewServerConfiguration() *ServerConfig {
	broker := envOr(common.EnvVarBrokerUri, "localhost:1883")
	funcPrefix := envOr(common.EnvVarFunctionPrefix, "/F/")
	eventPrefix := envOr(common.EnvVarEventPrefix, "/E/")

	sc := &ServerConfig{
		ClientID:       "ppc",
		BrokerURI:      broker,
		EventPrefix:    eventPrefix,
		FunctionPrefix: funcPrefix,
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
