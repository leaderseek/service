package config

import (
	"fmt"

	"github.com/friendsofgo/errors"
)

type HTTPConfig struct {
	Port uint64 `env:"PORT,required"`
}

func (httpConfig *HTTPConfig) Validate() error {
	port := httpConfig.Port

	if port != 80 && port != 443 && port < 1025 {
		return errors.New("config error: invalid http port")
	}

	return nil
}

func (httpConfig *HTTPConfig) Address() string {
	return fmt.Sprintf(":%d", httpConfig.Port)
}
