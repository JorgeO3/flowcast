package configs

import "github.com/caarlos0/env/v11"

// CatalogConfig holds the configuration for the catalog service.
type CatalogConfig struct {
	AppName                 string   `env:"CATALOG_APP_NAME"`
	Host                    string   `env:"CATALOG_HOST"`
	Port                    string   `env:"CATALOG_PORT"`
	DatabaseURL             string   `env:"CATALOG_DB_URL"`
	DBName                  string   `env:"CATALOG_DB_NAME"`
	RawAudioBucketName      string   `env:"RAW_AUDIO_BUCKET_NAME"`
	RawAudioBucketURL       string   `env:"RAW_AUDIO_BUCKET_URL"`
	RawAudioBucketAccessKey string   `env:"RAW_AUDIO_BUCKET_ACCESS_KEY"`
	RawAudioBucketSecretKey string   `env:"RAW_AUDIO_BUCKET_SECRET_KEY"`
	AssetsBucketName        string   `env:"ASSETS_BUCKET_NAME"`
	AssetsBucketURL         string   `env:"ASSETS_BUCKET_URL"`
	AssetsBucketAccessKey   string   `env:"ASSETS_BUCKET_ACCESS_KEY"`
	AssetsBucketSecretKey   string   `env:"ASSETS_BUCKET_SECRET_KEY"`
	RedpandaBrokers         []string `env:"REPANDA_BROKERS"`

	Version  string `env:"CATALOG_VERSION"`
	LogLevel string `env:"CATALOG_LOG_LEVEL"`
}

// LoadCatalogConfig loads the configuration for the catalog service.
func LoadCatalogConfig() (*CatalogConfig, error) {
	cfg := &CatalogConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
