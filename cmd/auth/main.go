// Main file for the auth service
package main

import (
	"log"

	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/internal/auth/app"
)

func main() {
	// Configuration
	cfg, err := configs.LoadAuthConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Start server
	app.Run(cfg)
}
