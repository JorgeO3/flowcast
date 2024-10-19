package configs

import (
	"github.com/caarlos0/env/v11"
)

// AudsyncConfig holds the configuration for the audsync service.
type AudsyncConfig struct {
	AppName     string `env:"AUDSYNC_APP_NAME"`
	Host        string `env:"AUDSYNC_HOST"`
	Port        string `env:"AUDSYNC_PORT"`
	DatabaseURL string `env:"AUDSYNC_DB_URL"`
	DBName      string `env:"AUDSYNC_DB_NAME"`
	// Redpanda
	Version  string `env:"VERSION"`
	LogLevel string `env:"LOG_LEVEL"`
}

// LoadAudsyncConfig loads the configuration for the proxy service.
func LoadAudsyncConfig() (*AudsyncConfig, error) {
	cfg := &AudsyncConfig{}
	if err := env.Parse(cfg); err != nil {
		return &AudsyncConfig{}, nil
	}
	return cfg, nil
}
