// Package app provides the entry point to the authentication service.
package app

import (
	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
)

// Run starts the auth service.
func Run(cfg *configs.AuthConfig, logg logger.Interface) {
	logg.Info("Starting auth service")
}
