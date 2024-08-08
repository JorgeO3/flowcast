package configs

import "github.com/caarlos0/env/v11"

// CatalogConfig holds the configuration for the catalog service.
type CatalogConfig struct {
	AppName        string `env:"APP_NAME"`
	HTTPHost       string `env:"HTTP_HOST"`
	HTTPPort       string `env:"HTTP_PORT"`
	DatabaseURL    string `env:"PG_URL"`
	DBName         string `env:"DB_NAME"`
	Version        string `env:"VERSION"`
	LogLevel       string `env:"LOG_LEVEL"`
	MigrationsPath string `env:"MIGRATIONS_PATH"`
}

// LoadCatalogConfig loads the configuration for the catalog service.
func LoadCatalogConfig() (*CatalogConfig, error) {
	cfg := &CatalogConfig{}
	if err := env.Parse(&cfg); err != nil {
		return &CatalogConfig{}, nil
	}
	return cfg, nil
}
