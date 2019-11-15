package servers

import (
	"fmt"
	"github.com/jw3/ppc/common"
	"os"
)

type ServerConfig struct {
	brokerHostname string
	BrokerURI      string
	ClientID       string
}

func NewServerConfiguration() *ServerConfig {
	host, ok := os.LookupEnv(common.EnvVarBrokerHostname)
	if !ok {
		host = "localhost"
	}
	uri := fmt.Sprintf("tcp://%s:1883", host)

	sc := &ServerConfig{
		brokerHostname: host,
		BrokerURI:      uri,
		ClientID:       "ppc",
	}
	return sc
}
