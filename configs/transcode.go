package configs

import (
	"github.com/caarlos0/env/v11"
)

// TranscodeConfig holds the configuration for the audsync service.
type TranscodeConfig struct {
	AppName  string `env:"TRANSCODE_APP_NAME"`
	Host     string `env:"TRANSCODE_HOST"`
	Port     string `env:"TRANSCODE_PORT"`
	Version  string `env:"TRANSCODE_VERSION"`
	LogLevel string `env:"TRANSCODE_LOG_LEVEL"`

	RawAudioBucketName      string `env:"RAW_AUDIO_BUCKET_NAME"`
	RawAudioBucketURL       string `env:"RAW_AUDIO_BUCKET_URL"`
	RawAudioBucketAccessKey string `env:"RAW_AUDIO_BUCKET_ACCESS_KEY"`
	RawAudioBucketSecretKey string `env:"RAW_AUDIO_BUCKET_SECRET_KEY"`

	EncodedOpusBucketName      string `env:"ENCODED_OPUS_BUCKET_NAME"`
	EncodedOpusBucketURL       string `env:"ENCODED_OPUS_BUCKET_URL"`
	EncodedOpusBucketAccessKey string `env:"ENCODED_OPUS_BUCKET_ACCESS_KEY"`
	EncodedOpusBucketSecretKey string `env:"ENCODED_OPUS_BUCKET_SECRET_KEY"`
}

// LoadTranscodeConfig loads the configuration for the proxy service.
func LoadTranscodeConfig() (*TranscodeConfig, error) {
	cfg := &TranscodeConfig{}
	if err := env.Parse(cfg); err != nil {
		return &TranscodeConfig{}, nil
	}
	return cfg, nil
}
