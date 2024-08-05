package configs

import "github.com/caarlos0/env/v11"

// CatalogConfig holds the configuration for the catalog service.
type CatalogConfig struct {
}

// LoadCatalogConfig loads the configuration for the catalog service.
func LoadCatalogConfig() (*CatalogConfig, error) {
	cfg := &CatalogConfig{}
	if err := env.Parse(&cfg); err != nil {
		return &CatalogConfig{}, nil
	}
	return cfg, nil
}
