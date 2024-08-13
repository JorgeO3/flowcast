// Package app provides the entry point to the catalog service.
package app

import (
	"fmt"

	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/internal/catalog/entity"
	"gitlab.com/JorgeO3/flowcast/internal/catalog/repository"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
	"gitlab.com/JorgeO3/flowcast/pkg/mongo"
)

// Run starts the catalog service.
func Run(cfg *configs.CatalogConfig, logg logger.Interface) {
	logg.Info("Starting catalog service")

	mg, err := mongo.New(cfg.DatabaseURL)
	if err != nil {
		logg.Fatal(fmt.Errorf("mongo connection error: %w", err))
	}
	defer mg.Close()

	logg.Debug("Connected to MongoDB")

	db := mg.Client.Database(entity.Database)

	actRepo := repository.NewMongoActRepository(db, entity.ActCollection)

}
