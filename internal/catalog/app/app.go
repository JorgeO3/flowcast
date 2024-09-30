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
		logg.Fatal("mongo connection error: %v", err)
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

	getActByIDUC := uc.NewGetActByID(
		uc.WithGetAcByIDLogger(logg),
		uc.WithGetAcByIDValidator(val),
		uc.WithGetAcByIDRepository(actRepo),
	)

	deleteActUC := uc.NewDeleteAct(
		uc.WithDeleteActLogger(logg),
		uc.WithDeleteActValidator(val),
		uc.WithDeleteActRepository(actRepo),
	)

	createManyUC := uc.NewCreateMany(
		uc.WithCreateManyLogger(logg),
		uc.WithCreateManyValidator(val),
		uc.WithCreateManyRepository(actRepo),
	)

	getActsUC := uc.NewGetActs(
		uc.WithGetActsLogger(logg),
		uc.WithGetActsValidator(val),
		uc.WithGetActsRepository(actRepo),
	)

	controller := c.New(
		c.WithConfig(cfg),
		c.WithLogger(logg),
		c.WithGetActsUC(getActsUC),
		c.WithCreateActUC(createActUC),
		c.WithUpdateActUC(updateActUC),
		c.WithDeleteActUC(deleteActUC),
		c.WithGetActByIDUC(getActByIDUC),
		c.WithCreateManyUC(createManyUC),
	)

	// Create a new router using the chi library.
	r := chi.NewRouter()

	// Set up middlewares for the router.
	r.Use(middleware.RequestID)                 // Middleware to inject request ID into the context.
	r.Use(middleware.RealIP)                    // Middleware to set the RemoteAddr to the IP address of the request.
	r.Use(logger.ZerologMiddleware(logg))       // Custom middleware to log HTTP requests with zerolog.
	r.Use(middleware.Recoverer)                 // Middleware to recover from panics and send an appropriate error response.
	r.Use(middleware.Heartbeat("/"))            // Middleware to provide a healthcheck endpoint at the root path.
	r.Use(middleware.RequestSize(1024 * 1024))  // Middleware to limit the maximum request size to 1 MB.
	r.Use(middleware.Timeout(60 * time.Second)) // Middleware to set a timeout of 60 seconds for each request.
	r.Use(logger.ErrorHandlingMiddleware(logg)) // Custom middleware to handle server errors.

	// Set up routes for the router.
	r.Get("/acts", controller.GetActs)
	r.Post("/acts", controller.CreateAct)
	r.Get("/acts/{id}", controller.GetAct)
	r.Put("/acts/{id}", controller.UpdateAct)
	r.Delete("/acts/{id}", controller.DeleteAct)
	r.Post("/acts/bulk", controller.CreateMany)

	// Construct the server address using the host and port specified in the configuration.
	addr := fmt.Sprintf("%s:%s", cfg.HTTPHost, cfg.HTTPPort)
	logg.Info("Starting server on " + addr)

	// Start the server and listen for incoming requests.
	if err := http.ListenAndServe(addr, r); err != nil {
		logg.Fatal("failed to start server: %w", err)
	}
}
