// Package http provides the HTTP Controllers for the catalog service.
package http

import (
	"net/http"

	"github.com/JorgeO3/flowcast/configs"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/usecase"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Controller handles HTTP requests for the catalog service and delegates operations to use cases.
type Controller struct {
	GetActsUC    *usecase.GetActsUC
	DeleteActUC  *usecase.DeleteActUC
	CreateActUC  *usecase.CreateActUC
	UpdateActUC  *usecase.UpdateActUC
	GetActByIDUC *usecase.GetActByIDUC
	CreateManyUC *usecase.CreateActsUC

	Logger logger.Interface
	Cfg    *configs.CatalogConfig
}

// CreateAct handles the creation of a new act.
// It decodes the JSON request body into CreateActInput, executes the create use case,
// and responds with the created act or an error.
func (c *Controller) CreateAct(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := c.withTimeout(r)
	defer cancel()

	var input usecase.CreateActInput
	if !c.decodeJSON(w, r, &input) {
		return
	}

	output, err := c.CreateActUC.Execute(ctx, input)
	if err != nil {
		c.Logger.Error("Error executing CreateAct use case - err", err)
		c.handleError(w, err)
		return
	}

	c.respondJSON(w, output)
}

// UpdateAct handles updating an existing act.
// It decodes the JSON request body into UpdateActInput, executes the update use case,
// and responds with the updated act or an error.
func (c *Controller) UpdateAct(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := c.withTimeout(r)
	defer cancel()

	var input usecase.UpdateActInput
	if !c.decodeJSON(w, r, &input) {
		return
	}

	output, err := c.UpdateActUC.Execute(ctx, input)
	if err != nil {
		c.Logger.Error("Error executing UpdateAct use case - err", err)
		c.handleError(w, err)
		return
	}

	c.respondJSON(w, output)
}

// GetAct retrieves an act by its MongoDB ObjectID.
func (c *Controller) GetAct(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := c.withTimeout(r)
	defer cancel()

	idParam := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.Logger.Error("Invalid ID format - id: %s, error: %v", idParam, err)
		c.handleError(w, errors.NewBadRequest("Invalid ID format", err))
		return
	}

	output, err := c.GetActByIDUC.Execute(ctx, usecase.GetActByIDInput{ID: id})
	if err != nil {
		c.Logger.Error("Error executing GetActByID use case - err %v", err)
		c.handleError(w, err)
		return
	}

	c.respondJSON(w, output)
}

// GetActs retrieves all acts with pagination and optional filtering by genre.
// It executes the GetActs use case and responds with the results.
func (c *Controller) GetActs(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := c.withTimeout(r)
	defer cancel()

	query := r.URL.Query()
	genre := query.Get("genre")
	limit, offset, err := parsePaginationParams(query)
	if err != nil {
		c.Logger.Error("Invalid pagination parameters - err %v", err)
		c.handleError(w, errors.NewBadRequest("Invalid pagination parameters", err))
		return
	}

	input := usecase.GetActsInput{
		Limit:  limit,
		Offset: offset,
		Genre:  genre,
	}

	output, err := c.GetActsUC.Execute(ctx, input)
	if err != nil {
		c.Logger.Error("Error executing GetActs use case - err %v", err)
		c.handleError(w, err)
		return
	}

	c.respondJSON(w, output)
}

// DeleteAct handles the deletion of an act by its MongoDB ObjectID.
// It parses the ID from the URL, executes the delete use case, and responds with an error or success message.
func (c *Controller) DeleteAct(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := c.withTimeout(r)
	defer cancel()

	id := chi.URLParam(r, "id")
	output, err := c.DeleteActUC.Execute(ctx, usecase.DeleteActInput{ID: id})
	if err != nil {
		c.Logger.Error("Error executing DeleteAct use case - err %v", err)
		c.handleError(w, err)
		return
	}

	c.respondJSON(w, output)
}

// CreateMany handles the creation of multiple acts.
// It decodes the JSON request body into CreateManyInput, executes the create many use case,
// and responds with the created acts or an error.
func (c *Controller) CreateMany(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := c.withTimeout(r)
	defer cancel()

	var input usecase.CreateActsInput
	if !c.decodeJSON(w, r, &input) {
		return
	}

	output, err := c.CreateManyUC.Execute(ctx, input)
	if err != nil {
		c.Logger.Error("Error executing CreateMany use case - err %v", err)
		c.handleError(w, err)
		return
	}

	c.respondJSON(w, output)
}
