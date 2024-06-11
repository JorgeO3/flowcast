package configs

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppName        string `env:"APP_NAME"`
	Version        string `env:"VERSION"`
	LogLevel       string `env:"LOG_LEVEL"`
	HttpHost       string `env:"HTTP_HOST"`
	HttpPort       string `env:"HTTP_PORT"`
	GrpcBrokerHost string `env:"GRPC_BROKER_HOST"`
	GrpcBrokerPort string `env:"GRPC_BROKER_PORT"`
	GrpcAuthHost   string `env:"GRPC_AUTH_HOST"`
	GrpcAuthPort   string `env:"GRPC_AUTH_PORT"`
}

func NewProxyConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("error parsing environment variables - e: %s", err)
	}
	return &Config{}
}
