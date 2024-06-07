package configs

type Config struct {
	appName         string
	version         string
	httpHost        string
	httpPort        string
	grpcProductHost string
	grpcProductPort string
}

func NewProxyConfig() (*Config, error) {
	cfg := &Config{}

	return &Config{}, nil
}
