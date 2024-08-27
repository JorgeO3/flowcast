// Package http provides the HTTP Controller for the catalog service.
package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/JorgeO3/flowcast/configs"
	"github.com/JorgeO3/flowcast/internal/catalog/usecase"
	"github.com/JorgeO3/flowcast/pkg/logger"
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

	var input usecase.CreateActInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.Logger.Error("Error decoding request body for CreateAct")
		c.handleError(w, err)
		return
	}

	output, err := c.CreateActUC.Execute(ctx, input)
	if err != nil {
		c.Logger.Error("Error in CreateAct use case execution")
		c.handleError(w, err)
	}

	c.respondJSON(w, output)
}
