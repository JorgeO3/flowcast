package app

import (
	"fmt"
	"net/http"

	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/internal/auth/controller"
	"gitlab.com/JorgeO3/flowcast/internal/auth/repository"
	"gitlab.com/JorgeO3/flowcast/internal/auth/usecase"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
	"gitlab.com/JorgeO3/flowcast/pkg/postgres"
)

func Run(cfg *configs.AuthConfig) {
	l := logger.New(cfg.LogLevel)

	pg, err := postgres.New(cfg.DatabaseURL)
	if err != nil {
		l.Fatal(fmt.Errorf("postgres connection error: %s", err))
	}
	defer pg.Close()

	userRepository := repository.NewPostgresUserRepo(pg)

	userRegistrationUseCase := usecase.NewUserRegistrationUseCase(userRepository)
	userAuthenticationUseCase := usecase.NewUserAuthenticationUseCase(userRepository)
	confirmRegistrationUseCase := usecase.NewConfirmRegistrationUseCase(userRepository)

	authController := &controller.AuthController{
		UserRegistrationUseCase:    userRegistrationUseCase,
		UserAuthenticationUseCase:  userAuthenticationUseCase,
		ConfirmRegistrationUseCase: confirmRegistrationUseCase,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/register", authController.Register)
	mux.HandleFunc("/authentication", authController.Authenticate)
	mux.HandleFunc("/confirmation", authController.ConfirmRegistration)

	http.ListenAndServe(":8080", mux)
}
