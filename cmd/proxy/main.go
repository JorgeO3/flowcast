package main

import (
	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/internal/proxy"
)

func main() {
	// Configuration
	cfg := configs.NewProxyConfig()

	// Start server
	proxy.Run(cfg)
}
