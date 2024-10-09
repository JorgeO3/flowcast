//go:build !grpc

// Package app provides the entry point to the catalog service.
package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JorgeO3/flowcast/configs"
	actc "github.com/JorgeO3/flowcast/internal/catalog/controller/http/act"
	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	actr "github.com/JorgeO3/flowcast/internal/catalog/repository/act"
	rar "github.com/JorgeO3/flowcast/internal/catalog/repository/rawaudio"
	actuc "github.com/JorgeO3/flowcast/internal/catalog/usecase/act"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/minio"
	"github.com/JorgeO3/flowcast/pkg/mongodb"
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
		logg.Fatal("minio connection error: %v", err)
	}

	raRepo := rar.NewRepository(raClient.GetClient(), cfg.RawAudioBucketName)

	val := validator.New()

	createActUC := actuc.NewCreateAct(
		actuc.WithCreateActLogger(logg),
		actuc.WithCreateActValidator(val),
		actuc.WithCreateActRepository(actRepo),
		actuc.WithCreateActRARepository(raRepo),
	)

	updateActUC := actuc.NewUpdateAct(
		actuc.WithUpdateActLogger(logg),
		actuc.WithUpdateActValidator(val),
		actuc.WithUpdateActRepository(actRepo),
		actuc.WithUpdateActRaRepository(raRepo),
	)

	getActByIDUC := actuc.NewGetActByID(
		actuc.WithGetAcByIDLogger(logg),
		actuc.WithGetAcByIDValidator(val),
		actuc.WithGetAcByIDRepository(actRepo),
	)

	deleteActUC := actuc.NewDeleteAct(
		actuc.WithDeleteActLogger(logg),
		actuc.WithDeleteActValidator(val),
		actuc.WithDeleteActRepository(actRepo),
		actuc.WithDeleteActRaRepository(raRepo),
	)

	createManyUC := actuc.NewCreateActs(
		actuc.WithCreateActsLogger(logg),
		actuc.WithCreateActsValidator(val),
		actuc.WithCreateActsRepository(actRepo),
		actuc.WithCreateActsRaRepository(raRepo),
	)

	getActsUC := actuc.NewGetActs(
		actuc.WithGetActsLogger(logg),
		actuc.WithGetActsValidator(val),
		actuc.WithGetActsRepository(actRepo),
	)

	controller := actc.New(
		actc.WithConfig(cfg),
		actc.WithLogger(logg),
		actc.WithGetActsUC(getActsUC),
		actc.WithCreateActUC(createActUC),
		actc.WithUpdateActUC(updateActUC),
		actc.WithDeleteActUC(deleteActUC),
		actc.WithGetActByIDUC(getActByIDUC),
		actc.WithCreateManyUC(createManyUC),
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

	// Set up routes for the router.
	r.Route("/acts", func(r chi.Router) {
		r.Get("/", controller.GetActs)
		r.Post("/", controller.CreateAct)
		r.Post("/bulk", controller.CreateMany)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", controller.GetAct)
			r.Put("/", controller.UpdateAct)
			r.Delete("/", controller.DeleteAct)
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
