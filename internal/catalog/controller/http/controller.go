// Package http provides the HTTP Controller for the catalog service.
package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/JorgeO3/flowcast/configs"
	uc "github.com/JorgeO3/flowcast/internal/catalog/usecase"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Controller is the HTTP controller for the catalog service.
type Controller struct {
	CreateActUC  *uc.CreateActUC
	UpdateActUC  *uc.UpdateActUC
	GetActByIDUC *uc.GetActByIDUC

	Logger logger.Interface
	Cfg    *configs.CatalogConfig
}

// CreateAct creates a new act.
func (c *Controller) CreateAct(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var input uc.CreateActInput

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

// UpdateAct updates an act.
func (c *Controller) UpdateAct(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var input uc.UpdateActInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.Logger.Error("Error decoding request body for UpdateAct")
		c.handleError(w, err)
		return
	}

	output, err := c.UpdateActUC.Execute(ctx, input)
	if err != nil {
		c.Logger.Error("Error in UpdateAct use case execution")
		c.handleError(w, err)
	}

	c.respondJSON(w, output)
}

// GetActByID gets an act by its ID.
func (c *Controller) GetActByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		c.Logger.Error("Error decoding id", "error", err)
		c.handleError(w, err)
		return
	}

	output, err := c.GetActByIDUC.Execute(ctx, uc.GetActByIDInput{ID: id})
	if err != nil {
		c.Logger.Error("Error in GetActByID use case execution", "error", err)
		c.handleError(w, err)
	}

	c.respondJSON(w, output)
}
