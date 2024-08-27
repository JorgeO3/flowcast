// Package app provides the entry point to the catalog service.
package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JorgeO3/flowcast/configs"
	c "github.com/JorgeO3/flowcast/internal/catalog/controller/http"
	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
	uc "github.com/JorgeO3/flowcast/internal/catalog/usecase"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/mongodb"
	"github.com/JorgeO3/flowcast/pkg/validator"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Run starts the catalog service.
func Run(cfg *configs.CatalogConfig, logg logger.Interface) {
	logg.Info("Starting catalog service")

	mg, err := mongodb.New(cfg.DatabaseURL)
	if err != nil {
		logg.Fatal("mongo connection error: %w", err)
	}
	defer mg.Close()

	logg.Debug("Connected to MongoDB")

	db := mg.Client.Database(entity.Database)
	actRepo := repository.NewMongoActRepository(db, entity.ActCollection)

	val := validator.New()

	createActUC := uc.NewCreateAct(
		uc.WithCreateActLogger(logg),
		uc.WithCreateActValidator(val),
		uc.WithCreateActRepository(actRepo),
	)

	updateActUC := uc.NewUpdateAct(
		uc.WithUpdateActLogger(logg),
		uc.WithUpdateActValidator(val),
		uc.WithUpdateActRepository(actRepo),
	)

	controller := c.New(
		c.WithConfig(cfg),
		c.WithLogger(logg),
		c.WithCreateActUC(createActUC),
		c.WithUpdateActUC(updateActUC),
	)

	// Create a new router using the chi library.
	r := chi.NewRouter()

	// Set up middlewares for the router.
	r.Use(middleware.RequestID)                 // Middleware to inject request ID into the context.
	r.Use(middleware.RealIP)                    // Middleware to set the RemoteAddr to the IP address of the request.
	r.Use(logger.ZerologMiddleware(logg))       // Custom middleware to log HTTP requests with zerolog.
	r.Use(middleware.Recoverer)                 // Middleware to recover from panics and send an appropriate error response.
	r.Use(middleware.Heartbeat("/"))            // Middleware to provide a healthcheck endpoint at the root path.
	r.Use(middleware.RequestSize(1024 * 100))   // Middleware to limit the maximum request size to 100 KB.
	r.Use(middleware.Timeout(60 * time.Second)) // Middleware to set a timeout of 60 seconds for each request.
	r.Use(logger.ErrorHandlingMiddleware(logg)) // Custom middleware to handle server errors.

	// Set up routes for the router.
	r.Post("/act", controller.CreateAct)
	r.Put("/act", controller.UpdateAct)

	// Construct the server address using the host and port specified in the configuration.
	addr := fmt.Sprintf("%s:%s", cfg.HTTPHost, cfg.HTTPPort)
	logg.Info("Starting server on " + addr)

	// Start the server and listen for incoming requests.
	if err := http.ListenAndServe(addr, r); err != nil {
		logg.Fatal("failed to start server: %w", err)
	}
}
