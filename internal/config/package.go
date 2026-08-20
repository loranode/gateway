// Package config loads runtime settings from the environment.
package config

import (
	"fmt"
	"strings"

	env "github.com/caarlos0/env/v11"
)

// Load reads configuration from the environment, applying defaults and
// normalising a bare NODE_ADDR host to include the default port.
func Load() (Config, error) {
	var c Config

	if err := env.Parse(&c); err != nil {
		return Config{}, fmt.Errorf("parse env: %w", err)
	}

	if !strings.Contains(c.NodeAddr, ":") {
		c.NodeAddr += ":4403"
	}

	return c, nil
}
