// Package main is the entry point for the packaging service.
package main

import (
	"log"

	"github.com/JorgeO3/flowcast/configs"
	"github.com/JorgeO3/flowcast/internal/transcode/app"
	"github.com/JorgeO3/flowcast/pkg/logger"
)

func main() {
	// Configuration
	cfg, err := configs.LoadTranscodeConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Create a new instance of the logger with the log level specified in the configuration.
	logg := logger.New(cfg.LogLevel)

	// Start server
	app.Run(cfg, logg)
}
