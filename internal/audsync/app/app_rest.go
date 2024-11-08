// Package app provides the entry point to the audsync service.
package app

import (
	"time"

	"github.com/JorgeO3/flowcast/configs"
	hc "github.com/JorgeO3/flowcast/internal/audsync/controller/http"
	"github.com/JorgeO3/flowcast/internal/audsync/infrastructure/kafka"
	"github.com/JorgeO3/flowcast/internal/audsync/repository"
	"github.com/JorgeO3/flowcast/internal/audsync/repository/assets"
	"github.com/JorgeO3/flowcast/internal/audsync/repository/audprocess"
	apuc "github.com/JorgeO3/flowcast/internal/audsync/usecase/audprocess"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"

	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/minio"
	"github.com/JorgeO3/flowcast/pkg/postgres"
	"github.com/JorgeO3/flowcast/pkg/txmanager"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// Run starts the rest server for the audsync service.
func Run(cfg *configs.AudsyncConfig, logg logger.Interface) {
	logg.Info("Starting audsync service")

	// Postgres
	pg, err := postgres.New(cfg.DBUrl)
	if err != nil {
		logg.Fatal("postgres connection error: %w", err)
	}
	defer pg.Close()

	// Run migrations
	pg.RunMigrations(cfg.DBMigrations, cfg.DBName)

	// Start transaction manager
	_ = txmanager.New(pg.Pool)

	// Assets bucket
	assetsClient, err := minio.New(
		cfg.AssetsBucketURL,
		minio.WithSSL(false),
		minio.WithCredentials(cfg.AssetsBucketAccessKey, cfg.AssetsBucketSecretKey),
	)
	if err != nil {
		logg.Fatal("minio assetsClient connection error: %v", err)
	}
	assetsRepo := assets.NewRepository(assetsClient.GetClient(), cfg.AssetsBucketName)

	// Transcoded Audio bucket
	transcodedAudioClient, err := minio.New(
		cfg.TranscodaudioBucketURL,
		minio.WithSSL(false),
		minio.WithCredentials(cfg.TranscodaudioBucketAccessKey, cfg.TranscodaudioBucketSecretKey),
	)
	if err != nil {
		logg.Fatal("minio transcodedAudioClient connection error: %v", err)
	}
	transcodaudioRepo := assets.NewRepository(transcodedAudioClient.GetClient(), cfg.TranscodaudioBucketName)

	// Audprocess repository
	audprocessRepo := audprocess.NewRepository(pg)

	// Validator
	validator := validator.New()

	// Repositories
	repos := repository.NewRepositories(
		repository.WithAssets(assetsRepo),
		repository.WithProcess(audprocessRepo),
		repository.WithTranscodaudio(transcodaudioRepo),
	)

	createProcessUC := apuc.NewCreateProcessUC(
		apuc.WithCreateProcessLogger(logg),
		apuc.WithCreateProcessRepos(repos),
		apuc.WithCreateProcessValidator(validator),
	)

	deleteProcessUC := apuc.NewDeleteProcessUC(
		apuc.WithDeleteProcessLogger(logg),
		apuc.WithDeleteProcessRepos(repos),
		apuc.WithDeleteProcessValidator(validator),
	)

	getManyProcessUC := apuc.NewGetManyProcessUC(
		apuc.WithGetManyProcessLogger(logg),
		apuc.WithGetManyProcessRepos(repos),
		apuc.WithGetManyProcessValidator(validator),
	)

	getProcessesUC := apuc.NewGetProcessUC(
		apuc.WithGetProcessLogger(logg),
		apuc.WithGetProcessRepos(repos),
		apuc.WithGetProcessValidator(validator),
	)

	updateProcessUC := apuc.NewUpdateProcessUC(
		apuc.WithUpdateProcessLogger(logg),
		apuc.WithUpdateProcessRepos(repos),
		apuc.WithUpdateProcessValidator(validator),
	)

	// Kafka Consumer
	consumer, err := kafka.NewConsumer(
		cfg.RedpandaBrokers,
		[]string{"audprocess"},
		kafka.WithConsumerConfig(cfg),
		kafka.WithConsumerLogger(logg),
		kafka.WithConsumerCreateProcessUC(createProcessUC),
		kafka.WithConsumerUpdateProcessUC(updateProcessUC),
		kafka.WithConsumerDeleteProcessUC(deleteProcessUC),
	)
	if err != nil {
		logg.Fatal("kafka consumer connection error: %v", err)
	}

	controller := hc.Controller{
		GetManyProcessUC: getManyProcessUC,
		GetProcessUC:     getProcessesUC,
		Logger:           logg,
		Cfg:              cfg,
	}

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

	r.Route("/process", func(r chi.Router) {
		r.Get("/", controller.GetManyProcess)
		r.Get("/{id}", controller.GetProcess)
	})

	orchestrator := NewOrchestrator(15*time.Second, logg)

	httpServer := NewHTTPServer(cfg.HTTPAddr, r, logg)
	orchestrator.AddService(httpServer)

	kafkaService := NewKafkaConsumerService(consumer, logg)
	orchestrator.AddService(kafkaService)

	// Iniciar todos los servicios y manejar el shutdown
	if err := orchestrator.Start(); err != nil {
		logg.Fatal("Service orchestration failed", err)
	}
}
