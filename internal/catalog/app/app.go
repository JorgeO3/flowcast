// Package app provides the entry point to the catalog service.
package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gitlab.com/JorgeO3/flowcast/configs"
	c "gitlab.com/JorgeO3/flowcast/internal/catalog/controller/http"
	"gitlab.com/JorgeO3/flowcast/internal/catalog/entity"
	"gitlab.com/JorgeO3/flowcast/internal/catalog/repository"
	"gitlab.com/JorgeO3/flowcast/internal/catalog/usecase"
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

	createActUC := usecase.NewCreateAct(actRepo, logg)

	controller := c.New(
		c.WithConfig(cfg),
		c.WithLogger(logg),
		c.WithCreateActUC(createActUC),
	)

	// Create a new router using the chi library.
	r := chi.NewRouter()

	// Set up middlewares for the router.
	r.Use(middleware.RequestID)                 // Middleware to inject request ID into the context.
	r.Use(middleware.RealIP)                    // Middleware to set the RemoteAddr to the IP address of the request.
	r.Use(logger.ZerologMiddleware(logg))       // Custom middleware to log HTTP requests with zerolog.
	r.Use(middleware.Recoverer)                 // Middleware to recover from panics and send an appropriate error response.
	r.Use(middleware.Heartbeat("/"))            // Middleware to provide a healthcheck endpoint at the root path.
	r.Use(middleware.RequestSize(1024))         // Middleware to limit the maximum request size to 1 KB.
	r.Use(middleware.Timeout(60 * time.Second)) // Middleware to set a timeout of 60 seconds for each request.
	r.Use(logger.ErrorHandlingMiddleware(logg)) // Custom middleware to handle server errors.

	// Set up routes for the router.
	r.Post("/act", controller.CreateAct)

	// Construct the server address using the host and port specified in the configuration.
	addr := fmt.Sprintf("%s:%s", cfg.HTTPHost, cfg.HTTPPort)
	logg.Info("Starting server on " + addr)

	// Start the server and listen for incoming requests.
	if err := http.ListenAndServe(addr, r); err != nil {
		logg.Fatal(fmt.Errorf("failed to start server: %w", err))
	}
}
