// Package app provides the entry point to the authentication service.
package app

import (
	"fmt"

	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
	"gitlab.com/JorgeO3/flowcast/pkg/mongo"
)

// Run starts the auth service.
func Run(cfg *configs.AuthConfig, logg logger.Interface) {
	logg.Info("Starting auth service")

	db, err := mongo.New(cfg.DatabaseURL)
	if err != nil {
		logg.Fatal(fmt.Errorf("mongo connection error: %w", err))
	}
	defer db.Close()
}
