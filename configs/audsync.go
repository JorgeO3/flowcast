package configs

import (
	"github.com/caarlos0/env/v11"
)

// AudsyncConfig holds the configuration for the audsync service.
type AudsyncConfig struct {
	AppName                      string   `env:"AUDSYNC_APP_NAME"`
	Host                         string   `env:"AUDSYNC_HOST"`
	Port                         string   `env:"AUDSYNC_PORT"`
	Version                      string   `env:"AUDSYNC_VERSION"`
	LogLevel                     string   `env:"AUDSYNC_LOG_LEVEL"`
	DBName                       string   `env:"AUDSYNC_DB_NAME"`
	DBUrl                        string   `env:"AUDSYNC_DB_URL"`
	DBMigrations                 string   `env:"AUDSYNC_MIGRATIONS"`
	AssetsBucketName             string   `env:"ASSETS_BUCKET_NAME"`
	AssetsBucketURL              string   `env:"ASSETS_BUCKET_URL"`
	AssetsBucketAccessKey        string   `env:"ASSETS_BUCKET_ACCESS_KEY"`
	AssetsBucketSecretKey        string   `env:"ASSETS_BUCKET_SECRET_KEY"`
	TranscodaudioBucketName      string   `env:"TRANSCODAUDIO_BUCKET_NAME"`
	TranscodaudioBucketURL       string   `env:"TRANSCODAUDIO_BUCKET_URL"`
	TranscodaudioBucketAccessKey string   `env:"TRANSCODAUDIO_BUCKET_ACCESS_KEY"`
	TranscodaudioBucketSecretKey string   `env:"TRANSCODAUDIO_BUCKET_SECRET_KEY"`
	RedpandaBrokers              []string `env:"REPANDA_BROKERS"`
}

// LoadAudsyncConfig loads the configuration for the proxy service.
func LoadAudsyncConfig() (*AudsyncConfig, error) {
	cfg := &AudsyncConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
