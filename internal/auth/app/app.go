// Package app is the entry point for the authentication service.
package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/internal/auth/controller"
	"gitlab.com/JorgeO3/flowcast/internal/auth/repository"
	"gitlab.com/JorgeO3/flowcast/internal/auth/usecase"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
	"gitlab.com/JorgeO3/flowcast/pkg/postgres"
)

// Run is the main function that sets up and starts the authentication service.
func Run(cfg *configs.AuthConfig) {
	// Create a new instance of the logger with the log level specified in the configuration.
	logg := logger.New(cfg.LogLevel)

	// Establish a connection to the PostgreSQL database using the database URL provided in the configuration.
	pg, err := postgres.New(cfg.DatabaseURL)
	if err != nil {
		logg.Fatal(fmt.Errorf("postgres connection error: %s", err))
	}
	defer pg.Close()

	// Run database migrations using the migrations path and database name specified in the configuration.
	pg.RunMigrations(cfg.MigrationsPath, cfg.DBName)

	// Initialize the user repository using the PostgreSQL database connection.
	userRepository := repository.NewPostgresUserRepo(pg)

	// Initialize the use cases related to user authentication.
	userRegistrationUseCase := usecase.NewUserRegistrationUseCase(userRepository)
	userAuthenticationUseCase := usecase.NewUserAuthenticationUseCase(userRepository)
	confirmRegistrationUseCase := usecase.NewConfirmRegistrationUseCase(userRepository)

	// Initialize the authentication controller with the use cases and logger.
	authController := &controller.AuthController{
		UserRegistrationUseCase:    userRegistrationUseCase,
		UserAuthenticationUseCase:  userAuthenticationUseCase,
		ConfirmRegistrationUseCase: confirmRegistrationUseCase,
		Logger:                     logg,
	}

	// Create a new router using the chi library.
	r := chi.NewRouter()

	// Set up middlewares for the router.
	r.Use(middleware.Logger)            // Middleware to log HTTP requests.
	r.Use(middleware.Recoverer)         // Middleware to recover from panics and send an appropriate error response.
	r.Use(middleware.Heartbeat("/"))    // Middleware to provide a healthcheck endpoint at the root path.
	r.Use(middleware.RequestSize(1024)) // Middleware to limit the maximum request size to 1 KB.

	// Define routes and corresponding handlers for the authentication service.
	r.Post("/register", authController.Register)
	r.Post("/authentication", authController.Authenticate)
	r.Post("/confirmation", authController.ConfirmRegistration)

	// Construct the server address using the host and port specified in the configuration.
	addr := fmt.Sprintf("%s:%s", cfg.HTTPHost, cfg.HTTPPort)
	fmt.Printf("Starting server on %s", addr)

	// Start the server and listen for incoming requests.
	if err := http.ListenAndServe(addr, r); err != nil {
		logg.Fatal("Failed to start server: %v", err)
	}
}
