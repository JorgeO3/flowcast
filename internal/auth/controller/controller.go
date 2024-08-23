// Package controller provides the Controller for the auth service.
package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/JorgeO3/flowcast/configs"
	"github.com/JorgeO3/flowcast/internal/auth/errors"
	"github.com/JorgeO3/flowcast/internal/auth/usecase"
	"github.com/JorgeO3/flowcast/pkg/logger"
)

// Controller handles authentication-related requests.
type Controller struct {
	UserRegistrationUseCase *usecase.UserRegistrationUseCase
	Log                     logger.Interface
	Cfg                     *configs.AuthConfig
}

// Register is an HTTP handler that processes user registration requests.
func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var input usecase.UserRegistrationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.handleError(w, errors.NewBadRequest("Invalid input", err))
		return
	}

	output, err := c.UserRegistrationUseCase.Execute(ctx, input, c.Cfg)
	if err != nil {
		c.handleError(w, err)
		return
	}

	c.respondJSON(w, output)
}
