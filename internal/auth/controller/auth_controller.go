package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"gitlab.com/JorgeO3/flowcast/internal/auth/usecase"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
)

type AuthController struct {
	UserRegistrationUseCase    *usecase.UserRegistrationUseCase
	UserAuthenticationUseCase  *usecase.UserAuthenticationUseCase
	ConfirmRegistrationUseCase *usecase.ConfirmRegistrationUseCase
	Logger                     logger.Interface
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var input usecase.UserRegistrationInput

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		c.Logger.Info("Failed to decode user input for registration - error: %s", err)
		return
	}

	output, err := c.UserRegistrationUseCase.Execute(ctx, input, c.Logger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		c.Logger.Error("Failed to register user", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		c.Logger.Error("Failed to encode response", "error", err)
		return
	}
}

func (c *AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {

}

func (c *AuthController) ConfirmRegistration(w http.ResponseWriter, r *http.Request) {

}
