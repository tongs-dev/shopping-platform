package common

import (
	"errors"
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
)

// GetConsulConfig sets up a configuration center using Consul as the key-value store,
// returns the configuration object loaded from Consul.
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	if host == "" || port <= 0 {
		return nil, errors.New("invalid Consul host or port")
	}

	// Creates a Consul Configuration Source
	consulSource := consul.NewSource(
		// Builds the Consul address dynamically
		consul.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		// Retrieves only the configuration keys under the specified prefix, default is /micro/config
		consul.WithPrefix(prefix),
		// Allows retrieving keys without the prefix
		consul.StripPrefix(true),
	)

	// Initializes the Config Object
	conf, err := config.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize config: %w", err)
	}

	// Loads the Consul Configuration
	if err := conf.Load(consulSource); err != nil {
		return nil, fmt.Errorf("failed to load config from Consul: %w", err)
	}

	return conf, err
}
