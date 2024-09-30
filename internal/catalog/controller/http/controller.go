// Package http provides the HTTP Controller for the catalog service.
package http

import (
	"fmt"
	"net/http"

	"github.com/JorgeO3/flowcast/configs"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	uc "github.com/JorgeO3/flowcast/internal/catalog/usecase"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Controller handles HTTP requests for the catalog service and delegates operations to use cases.
type Controller struct {
	GetActsUC    *uc.GetActsUC
	DeleteActUC  *uc.DeleteActUC
	CreateActUC  *uc.CreateActUC
	UpdateActUC  *uc.UpdateActUC
	GetActByIDUC *uc.GetActByIDUC
	CreateManyUC *uc.CreateManyUC

	Logger logger.Interface
	Cfg    *configs.CatalogConfig
}

// CreateAct handles the creation of a new act.
// It decodes the JSON request body into CreateActInput, executes the create use case,
// and responds with the created act or an error.
func (c *Controller) CreateAct(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := c.withTimeout(r)
	defer cancel()

	var input uc.CreateActInput
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

	var input uc.UpdateActInput
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

// GetAct retrieves acts based on query parameters.
// - If 'id' is provided, it returns a single act.
// - If 'genre' is provided, it returns acts of that genre with pagination.
// - Otherwise, it returns all acts with pagination.
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

	output, err := c.GetActByIDUC.Execute(ctx, uc.GetActByIDInput{ID: id})
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

	fmt.Println(limit, offset, genre)

	input := uc.GetActsInput{
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

	idParam := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.Logger.Error("Invalid ID format - id: %s, error: %v", idParam, err)
		c.handleError(w, errors.NewBadRequest("Invalid ID format", err))
		return
	}

	output, err := c.DeleteActUC.Execute(ctx, uc.DeleteActInput{ID: id})
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

	var input uc.CreateManyInput
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
