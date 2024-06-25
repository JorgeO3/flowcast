// Package controller implements the HTTP handlers for the authentication use cases.
package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/internal/auth/usecase"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
)

// AuthController is a controller that handles authentication-related requests.
type AuthController struct {
	UserRegUC    *usecase.UserRegUC
	UserAuthUC   *usecase.UserAuthUC
	ConfirmRegUC *usecase.ConfirmRegUC
	Logger       logger.Interface
	Cfg          *configs.AuthConfig
}

// Register is an HTTP handler that processes user registration requests.
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var input usecase.UserRegInput

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		c.Logger.Info("Failed to decode user input for registration", "error", err)
		return
	}

	output, err := c.UserRegUC.Execute(ctx, input, c.Cfg)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		c.Logger.Error("Failed to register user", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		c.Logger.Error("Failed to encode response", "error", err)
		return
	}
}

// Authenticate is an HTTP handler that processes user authentication requests.
func (c *AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {
	var input usecase.UserAuthInput

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		c.Logger.Info("Failed to decode user input for authentication", "error", err)
		return
	}

	output, err := c.UserAuthUC.Execute(ctx, input, c.Cfg)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		c.Logger.Error("Failed to authenticate user", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		c.Logger.Error("Failed to encode response", "error", err)
		return
	}
}

// ConfirmRegistration is an HTTP handler that processes user registration confirmation requests.
func (c *AuthController) ConfirmRegistration(w http.ResponseWriter, r *http.Request) {
	var input usecase.ConfirmRegInput

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		c.Logger.Info("Failed to decode user input for confirmation", "error", err)
		return
	}

	output, err := c.ConfirmRegUC.Execute(ctx, input, c.Cfg)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		c.Logger.Error("Failed to confirm registration", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		c.Logger.Error("Failed to encode response", "error", err)
		return
	}
}
