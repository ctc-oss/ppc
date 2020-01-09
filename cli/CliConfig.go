package cli

import (
	"fmt"
	"os"
)

const (
	EnvVarApiUri = "CLOUD_API"
)

type Config struct {
	ApiUri string
}

func NewConfiguration() *Config {
	h, ok := os.LookupEnv(EnvVarApiUri)
	if !ok {
		h = "localhost:9000"
	}

	sc := &Config{
		ApiUri: fmt.Sprintf("%s/v1", h),
	}
	return sc
}
