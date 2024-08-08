// Package configs contains the configuration for the application.
package configs

import "github.com/caarlos0/env/v11"

// AuthConfig -.
type AuthConfig struct {
	AppName        string `env:"APP_NAME"`
	HTTPHost       string `env:"HTTP_HOST"`
	HTTPPort       string `env:"HTTP_PORT"`
	AccEmail       string `env:"ACC_EMAIL"`
	AccPassword    string `env:"ACC_PASSWORD"`
	SMTPHost       string `env:"SMTP_HOST"`
	SMTPPort       string `env:"SMTP_PORT"`
	EmailTemplate  string `env:"EMAIL_TEMPLATE"`
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
