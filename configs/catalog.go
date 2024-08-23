package configs

import "github.com/caarlos0/env/v11"

// CatalogConfig holds the configuration for the catalog service.
type CatalogConfig struct {
	AppName     string `env:"CATALOG_APP_NAME"`
	HTTPHost    string `env:"CATALOG_HTTP_HOST"`
	HTTPPort    string `env:"CATALOG_HTTP_PORT"`
	DatabaseURL string `env:"CATALOG_DB_URL"`
	DBName      string `env:"CATALOG_DB_NAME"`
	Version     string `env:"CATALOG_VERSION"`
	LogLevel    string `env:"CATALOG_LOG_LEVEL"`
}

// LoadCatalogConfig loads the configuration for the catalog service.
func LoadCatalogConfig() (*CatalogConfig, error) {
	cfg := &CatalogConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
