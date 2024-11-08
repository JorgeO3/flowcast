// Package http provides the http controller for the audsync service.
package http

import (
	"net/http"

	"github.com/JorgeO3/flowcast/configs"
	e "github.com/JorgeO3/flowcast/internal/audsync/errors"
	apuc "github.com/JorgeO3/flowcast/internal/audsync/usecase/audprocess"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/go-chi/chi"
)

// Controller is the http controller for the audsync service
type Controller struct {
	GetProcessUC     *apuc.GetProcessUC
	GetManyProcessUC *apuc.GetManyProcessUC

	Logger logger.Interface
	Cfg    *configs.AudsyncConfig
}

// GetProcess execute the GetProcess use case
func (c *Controller) GetProcess(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := c.withTimeout(r)
	defer cancel()

	id := chi.URLParam(r, "id")
	if id == "" {
		c.Logger.Error("Invalid ID format - id: %s", id)
		c.handleError(w, e.NewValidation("Invalid ID format", nil))
		return
	}

	input := &apuc.GetProcessInput{ID: id}
	output, err := c.GetProcessUC.Execute(ctx, input)
	if err != nil {
		c.Logger.Error("Error executing GetProcess use case - err", err)
		c.handleError(w, err)
		return
	}

	c.respondJSON(w, output)
}

// GetManyProcess execute the GetManyProcess use case
func (c *Controller) GetManyProcess(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := c.withTimeout(r)
	defer cancel()

	query := r.URL.Query()
	limit, offset, err := parsePaginationParams(query)
	if err != nil {
		c.Logger.Error("Invalid pagination parameters - err %v", err)
		c.handleError(w, e.NewBadRequest("Invalid pagination parameters", err))
		return
	}

	input := &apuc.GetManyProcessInput{
		Limit:  limit,
		Offset: offset,
	}

	output, err := c.GetManyProcessUC.Execute(ctx, input)
	if err != nil {
		c.Logger.Error("Error executing GetManyProcess use case - err %v", err)
		c.handleError(w, err)
		return
	}

	c.respondJSON(w, output)
}

// CreateProcess execute the CreateProcess use case
// func (c *Controller) CreateProcess(w http.ResponseWriter, r *http.Request) {
// 	ctx, cancel := c.withTimeout(r)
// 	defer cancel()

// 	var input apuc.CreateProcessInput
// 	if !c.decodeJSON(w, r, &input) {
// 		return
// 	}

// 	output, err := c.CreateProcessUC.Execute(ctx, &input)
// 	if err != nil {
// 		c.Logger.Error("Error executing CreateProcess use case - err", err)
// 		c.handleError(w, err)
// 		return
// 	}

// 	c.respondJSON(w, output)
// }

// UpdateProcess execute the UpdateProcess use case
// func (c *Controller) UpdateProcess(w http.ResponseWriter, r *http.Request) {
// 	ctx, cancel := c.withTimeout(r)
// 	defer cancel()

// 	var input apuc.UpdateProcessInput
// 	if !c.decodeJSON(w, r, &input) {
// 		return
// 	}

// 	output, err := c.UpdateProcessUC.Execute(ctx, &input)
// 	if err != nil {
// 		c.Logger.Error("Error executing UpdateProcess use case - err", err)
// 		c.handleError(w, err)
// 		return
// 	}

// 	c.respondJSON(w, output)
// }

// DeleteProcess execute the DeleteProcess use case
// func (c *Controller) DeleteProcess(w http.ResponseWriter, r *http.Request) {
// 	ctx, cancel := c.withTimeout(r)
// 	defer cancel()

// 	id := chi.URLParam(r, "id")
// 	if id == "" {
// 		c.Logger.Error("Invalid ID format - id: %s", id)
// 		c.handleError(w, e.NewValidation("Invalid ID format", nil))
// 		return
// 	}

// 	input := &apuc.DeleteProcessInput{ID: id}
// 	output, err := c.DeleteProcessUC.Execute(ctx, input)
// 	if err != nil {
// 		c.Logger.Error("Error executing DeleteProcess use case - err", err)
// 		c.handleError(w, err)
// 		return
// 	}

// 	c.respondJSON(w, output)
// }
