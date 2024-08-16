// Package http provides the HTTP Controller for the catalog service.
package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/internal/auth/errors"
	"gitlab.com/JorgeO3/flowcast/internal/catalog/entity"
	"gitlab.com/JorgeO3/flowcast/internal/catalog/usecase"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
)

// Controller is the HTTP controller for the catalog service.
type Controller struct {
	CreateActUC *usecase.CreateActUC
	Logger      logger.Interface
	Cfg         *configs.CatalogConfig
}

// CreateAct creates a new act.
func (c *Controller) CreateAct(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var input entity.Act

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.handleError(w, errors.NewBadRequest("Invalid input", err))
		return
	}

	output, err := c.CreateActUC.Execute(ctx, input)
	if err != nil {
		c.handleError(w, err)
		return
	}

	c.respondJSON(w, output)
}
