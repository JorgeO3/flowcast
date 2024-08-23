package main

import (
	"log"

	"github.com/JorgeO3/flowcast/configs"
	"github.com/JorgeO3/flowcast/internal/proxy"
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
