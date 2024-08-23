// Package app provides the entry point to the authentication service.
package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JorgeO3/flowcast/configs"
	"github.com/JorgeO3/flowcast/internal/auth/controller"
	"github.com/JorgeO3/flowcast/internal/auth/repository"
	"github.com/JorgeO3/flowcast/internal/auth/service"
	"github.com/JorgeO3/flowcast/internal/auth/usecase"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/postgres"
	"github.com/JorgeO3/flowcast/pkg/smtp"
	"github.com/JorgeO3/flowcast/pkg/txmanager"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

// Run starts the auth service.
func Run(cfg *configs.AuthConfig, logg logger.Interface) {
	// Establish a connection to the PostgreSQL database using the database URL provided in the configuration.

	fmt.Printf("Database URL: %v", cfg.DatabaseURL)

	pg, err := postgres.New(cfg.DatabaseURL)
	if err != nil {
		logg.Fatal(fmt.Errorf("postgres connection error: %w", err))
	}
	defer pg.Close()

	// Run database migrations using the migrations path and database name specified in the configuration.
	pg.RunMigrations(cfg.MigrationsPath, cfg.DBName)

	// Initialize the transaction manager using the PostgreSQL database connection.
	txManager := txmanager.New(pg.Pool)

	// Initialize the user repository using the PostgreSQL database connection.
	userRepo := repository.NewPostgresUserRepo(pg)

	// Initialize the user preference repository using the PostgreSQL database connection.
	userPrefRepo := repository.NewPostgresUserPrefRepo(pg)

	// Initialize the social link repository using the PostgreSQL database connection.
	socialRepo := repository.NewPostgresSocialLinkRepo(pg)

	// Initialize the email verification token repository using the PostgreSQL database connection.
	emailRepo := repository.NewPostgresEmailVerificationTokenRepo(pg)

	// Create a new SMTP client using the SMTP configuration provided in the configuration.
	smtpCfg := smtp.NewConfig(cfg.SMTPHost, cfg.SMTPPort, cfg.AccEmail, cfg.AccPassword)
	smtpClient := smtp.NewSMTPClient(smtpCfg)

	mailer := service.NewMailerService(smtpClient)

	// Initialize the use cases related to user authentication.
	userRegUC := usecase.NewUserRegistrationUseCase(
		usecase.WithLogger(logg),
		usecase.WithMailer(mailer),
		usecase.WithUserRepo(userRepo),
		usecase.WithEmailRepo(emailRepo),
		usecase.WithTxManager(txManager),
		usecase.WithSocialRepo(socialRepo),
		usecase.WithUserPrefRepo(userPrefRepo),
	)

	// Initialize the authentication controller with the use cases and logger.
	authController := &controller.Controller{
		Cfg:                     cfg,
		Log:                     logg,
		UserRegistrationUseCase: userRegUC,
	}

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

	// Define routes and corresponding handlers for the authentication service.
	r.Post("/register", authController.Register)
	// r.Post("/authentication", authController.Authenticate)
	// r.Post("/confirmation", authController.ConfirmRegistration)

	// Construct the server address using the host and port specified in the configuration.
	addr := fmt.Sprintf("%s:%s", cfg.HTTPHost, cfg.HTTPPort)
	logg.Info("Starting server on " + addr)

	// Start the server and listen for incoming requests.
	if err := http.ListenAndServe(addr, r); err != nil {
		logg.Fatal(fmt.Errorf("failed to start server: %w", err))
	}
}
