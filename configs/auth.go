// Package configs contains the configuration for the auth service.
package configs

import "github.com/caarlos0/env/v11"

// AuthConfig -.
type AuthConfig struct {
	AppName        string `env:"APP_NAME"`
	HTTPHost       string `env:"HTTP_HOST"`
	HTTPPort       string `env:"HTTP_PORT"`
	DatabaseURL    string `env:"PG_URL"`
	DBName         string `env:"DB_NAME"`
	Version        string `env:"VERSION"`
	LogLevel       string `env:"LOG_LEVEL"`
	MigrationsPath string `env:"MIGRATIONS_PATH"`
}

// LoadAuthConfig -.
func LoadAuthConfig() (*AuthConfig, error) {
	cfg := &AuthConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
