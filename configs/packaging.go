package configs

import (
	"github.com/caarlos0/env/v11"
)

// PackagingConfig holds the configuration for the audsync service.
type PackagingConfig struct {
	AppName     string `env:"AUDSYNC_APP_NAME"`
	Host        string `env:"AUDSYNC_HOST"`
	Port        string `env:"AUDSYNC_PORT"`
	DatabaseURL string `env:"AUDSYNC_DB_URL"`
	DBName      string `env:"AUDSYNC_DB_NAME"`
	// Redpanda
	Version  string `env:"VERSION"`
	LogLevel string `env:"LOG_LEVEL"`
}

// LoadPackagingConfig loads the configuration for the proxy service.
func LoadPackagingConfig() (*PackagingConfig, error) {
	cfg := &PackagingConfig{}
	if err := env.Parse(cfg); err != nil {
		return &PackagingConfig{}, nil
	}
	return cfg, nil
}
