package configs

import "github.com/caarlos0/env/v11"

type AuthConfig struct {
	AppName     string `env:"APP_NAME"`
	DatabaseURL string `env:"PG_URL"`
	Version     string `env:"VERSION"`
	LogLevel    string `env:"LOG_LEVEL"`
}

func LoadAuthConfig() (*AuthConfig, error) {
	cfg := &AuthConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
