//go:build !grpc

// Package app provides the entry point to the catalog service.
package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JorgeO3/flowcast/configs"
	controller "github.com/JorgeO3/flowcast/internal/catalog/controller/http"
	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
	actr "github.com/JorgeO3/flowcast/internal/catalog/repository/act"
	rar "github.com/JorgeO3/flowcast/internal/catalog/repository/rawaudio"
	uc "github.com/JorgeO3/flowcast/internal/catalog/usecase"

	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/minio"
	"github.com/JorgeO3/flowcast/pkg/mongodb"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
	"github.com/JorgeO3/flowcast/pkg/validator"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

// Run starts the rest server for the catalog service.
func Run(cfg *configs.CatalogConfig, logg logger.Interface) {
	logg.Info("Starting catalog service")

	// MongoDB connection
	mg, err := mongodb.New(cfg.DatabaseURL)
	if err != nil {
		logg.Fatal("mongo connection error: %v", err)
	}
	defer mg.Close()

	logg.Debug("Connected to MongoDB")

	db := mg.Client.Database(entity.Database)
	actRepo := actr.NewRepository(db, entity.ActCollection)

	// Raw audio bucket
	raClient, err := minio.New(
		cfg.RawAudioBucketURL,
		minio.WithSSL(false),
		minio.WithCredentials(cfg.RawAudioBucketAccessKey, cfg.RawAudioBucketSecretKey),
	)
	if err != nil {
		logg.Fatal("minio rawClient connection error: %v", err)
	}
	raRepo := rar.NewRepository(raClient.GetClient(), cfg.RawAudioBucketName)

	// Assets bucket
	assetsClient, err := minio.New(
		cfg.AssetsBucketURL,
		minio.WithSSL(false),
		minio.WithCredentials(cfg.AssetsBucketAccessKey, cfg.AssetsBucketSecretKey),
	)
	if err != nil {
		logg.Fatal("minio assetsClient connection error: %v", err)
	}
	assetsRepo := rar.NewRepository(assetsClient.GetClient(), cfg.AssetsBucketName)

	// Create a new validator for the act controller.
	val := validator.New()

	// Create the redpanda producer for the act controller.
	pr, err := redpanda.NewProducer(cfg.RedpandaBrokers)
	if err != nil {
		logg.Fatal("failed to create redpanda producer: %w", err)
	}

	// Repositories
	repos := repository.
		NewRepositoryBuilder().
		WithActRepository(actRepo).
		WithAssetsRepository(assetsRepo).
		WithRawAudioRepository(raRepo).
		Build()

	// Create the use cases for the act controller.
	createActUC := uc.NewCreateAct(
		uc.WithCreateActProducer(pr),
		uc.WithCreateActLogger(logg),
		uc.WithCreateActValidator(val),
		uc.WithCreateActRepository(actRepo),
		uc.WithCreateActRARepository(raRepo),
		uc.WithCreateActAssRepository(assetsRepo),
	)

	updateActUC := uc.NewUpdateAct(
		uc.WithUpdateActLogger(logg),
		uc.WithUpdateActProducer(pr),
		uc.WithUpdateActValidator(val),
		uc.WithUpdateRepositories(repos),
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
		uc.WithDeleteActRaRepository(raRepo),
		uc.WithDeleteActProducer(pr),
		uc.WithDeleteActAssRepository(assetsRepo),
	)

	createManyUC := uc.NewCreateActs(
		uc.WithCreateActsLogger(logg),
		uc.WithCreateActsValidator(val),
		uc.WithCreateActsRepository(actRepo),
		uc.WithCreateActsRaRepository(raRepo),
		uc.WithCreateActsProducer(pr),
		uc.WithCreateActsAssRepository(assetsRepo),
	)

	getActsUC := uc.NewGetActs(
		uc.WithGetActsLogger(logg),
		uc.WithGetActsValidator(val),
		uc.WithGetActsRepository(actRepo),
	)

	actController := controller.New(
		controller.WithConfig(cfg),
		controller.WithLogger(logg),
		controller.WithGetActsUC(getActsUC),
		controller.WithCreateActUC(createActUC),
		controller.WithUpdateActUC(updateActUC),
		controller.WithDeleteActUC(deleteActUC),
		controller.WithGetActByIDUC(getActByIDUC),
		controller.WithCreateManyUC(createManyUC),
	)

	// Create a new router using the chi library.
	r := chi.NewRouter()

	// Set up middlewares for the router.
	r.Use(middleware.RequestID)                 // Middleware to inject request ID into the context.
	r.Use(middleware.RealIP)                    // Middleware to set the RemoteAddr to the IP address of the request.
	r.Use(logger.LoggingMiddleware(logg))       // Custom middleware to log HTTP requests with zerolog.
	r.Use(middleware.Recoverer)                 // Middleware to recover from panics and send an appropriate error response.
	r.Use(middleware.Heartbeat("/"))            // Middleware to provide a healthcheck endpoint at the root path.
	r.Use(middleware.RequestSize(1024 * 1024))  // Middleware to limit the maximum request size to 1 MB.
	r.Use(middleware.Timeout(60 * time.Second)) // Middleware to set a timeout of 60 seconds for each request.
	r.Use(httprate.LimitByIP(100, time.Minute)) // Middleware to limit the number of requests per IP address.
	r.Use(logger.ErrorHandlingMiddleware(logg)) // Custom middleware to handle server errors.
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Set up act routes for the router.
	r.Route("/acts", func(r chi.Router) {
		r.Get("/", actController.GetActs)
		r.Post("/", actController.CreateAct)
		r.Post("/bulk", actController.CreateMany)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", actController.GetAct)
			r.Put("/", actController.UpdateAct)
			r.Delete("/", actController.DeleteAct)
		})
	})

	// Construct the server address using the host and port specified in the configuration.
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	logg.Info("Http server listening on " + addr)

	// Start the server and listen for incoming requests.
	if err := http.ListenAndServe(addr, r); err != nil {
		logg.Fatal("failed to start server: %w", err)
	}
}
