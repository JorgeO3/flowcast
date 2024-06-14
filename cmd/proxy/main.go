package main

import (
	"log"

	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/internal/proxy"
)

func main() {
	// Configuration
	cfg, err := configs.LoadProxyConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Start server
	proxy.Run(cfg)
}
