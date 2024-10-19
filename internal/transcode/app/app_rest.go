// Package app is the entry point for the transcode service.
package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JorgeO3/flowcast/configs"
	controller "github.com/JorgeO3/flowcast/internal/transcode/controller/http"
	"github.com/JorgeO3/flowcast/internal/transcode/repository"
	"github.com/JorgeO3/flowcast/internal/transcode/usecase"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/minio"
	"github.com/JorgeO3/flowcast/pkg/validator"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

// Run starts the HTTP server for the transcode service.
func Run(cfg *configs.TranscodeConfig, logg logger.Interface) {
	logg.Info("Starting transcoding service")

	// Struct validator
	val := validator.New()

	// Raw audio repository
	raClient, err := minio.New(
		cfg.RawAudioBucketURL,
		minio.WithSSL(false),
		minio.WithCredentials(cfg.RawAudioBucketAccessKey, cfg.RawAudioBucketSecretKey),
	)
	if err != nil {
		logg.Fatal("minio connection error: %v", err)
	}
	raRepo := repository.NewRawAudioRepository(raClient.GetClient(), cfg.RawAudioBucketName)

	// Encoded Opus repository
	eoClient, err := minio.New(
		cfg.EncodedOpusBucketURL,
		minio.WithSSL(false),
		minio.WithCredentials(cfg.EncodedOpusBucketAccessKey, cfg.EncodedOpusBucketSecretKey),
	)
	if err != nil {
		logg.Fatal("minio connection error: %v", err)
	}
	eoRepo := repository.NewEncodedOpusRepository(eoClient.GetClient(), cfg.EncodedOpusBucketName)

	// Transcode song use case
	transcodeSongUC := usecase.NewTranscodeSongUC(
		usecase.WithTranscodeSongRepository(raRepo),
		usecase.WithTranscodeSongEncodedOpusRepository(eoRepo),
		usecase.WithTranscodeSongLogger(logg),
		usecase.WithTranscodeSongValidator(val),
	)

	// HTTP controller
	controller := controller.New(
		controller.WithTranscodeSongUC(transcodeSongUC),
		controller.WithLogger(logg),
		controller.WithConfig(cfg),
	)

	// Create a new router using the chi library.
	r := chi.NewRouter()

	// Set up middlewares for the router.
	r.Use(middleware.RequestID)                 // Middleware to inject request ID into the context.
	r.Use(middleware.RealIP)                    // Middleware to set the RemoteAddr to the IP address of the request.
	r.Use(logger.LoggingMiddleware(logg))       // Custom middleware to log HTTP requests with zerolog.
	r.Use(middleware.Recoverer)                 // Middleware to recover from panics and send an appropriate error response.
	r.Use(logger.ErrorHandlingMiddleware(logg)) // Custom middleware to handle server errors.
	r.Use(middleware.Heartbeat("/"))            // Middleware to provide a healthcheck endpoint at the root path.
	r.Use(middleware.RequestSize(1024 * 1024))  // Middleware to limit the maximum request size to 1 MB.
	r.Use(middleware.Timeout(60 * time.Second)) // Middleware to set a timeout of 60 seconds for each request.
	r.Use(httprate.LimitByIP(100, time.Minute)) // Middleware to limit the number of requests per IP address.
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Set up routes for the router.
	r.Post("/", controller.TranscodeSong)

	// Construct the server address using the host and port specified in the configuration.
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	logg.Info("Http server listening on " + addr)

	// Start the server and listen for incoming requests.
	if err := http.ListenAndServe(addr, r); err != nil {
		logg.Fatal("failed to start server: %w", err)
	}
}
