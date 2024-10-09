package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/JorgeO3/flowcast/configs"
	e "github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/usecase/act"
	"github.com/JorgeO3/flowcast/pkg/logger"
)

// Opts represents the functional options for the controller.
type Opts func(c *ActController)

// WithCreateActUC sets the CreateActUC in the controller.
func WithCreateActUC(uc *act.CreateActUC) Opts {
	return func(c *ActController) {
		c.CreateActUC = uc
	}
}

// WithUpdateActUC sets the UpdateActUC in the controller.
func WithUpdateActUC(uc *act.UpdateActUC) Opts {
	return func(c *ActController) {
		c.UpdateActUC = uc
	}
}

// WithGetActByIDUC sets the GetActByIDUC in the controller.
func WithGetActByIDUC(uc *act.GetActByIDUC) Opts {
	return func(c *ActController) {
		c.GetActByIDUC = uc
	}
}

// WithDeleteActUC sets the DeleteActUC in the controller.
func WithDeleteActUC(uc *act.DeleteActUC) Opts {
	return func(c *ActController) {
		c.DeleteActUC = uc
	}
}

// WithGetActsUC sets the GetActsUC in the controller.
func WithGetActsUC(uc *act.GetActsUC) Opts {
	return func(c *ActController) {
		c.GetActsUC = uc
	}
}

// WithCreateManyUC sets the CreateManyUC in the controller.
func WithCreateManyUC(uc *act.CreateActsUC) Opts {
	return func(c *ActController) {
		c.CreateManyUC = uc
	}
}

// WithLogger sets the logger in the controller.
func WithLogger(l logger.Interface) Opts {
	return func(c *ActController) {
		c.Logger = l
	}
}

// WithConfig sets the config in the controller.
func WithConfig(cfg *configs.CatalogConfig) Opts {
	return func(c *ActController) {
		c.Cfg = cfg
	}
}

// New creates a new instance of ActController.
func New(opts ...Opts) *ActController {
	c := &ActController{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// handleError processes errors and sends appropriate HTTP responses.
// It distinguishes between known catalog errors and unexpected internal errors,
// logging each accordingly.
func (c *ActController) handleError(w http.ResponseWriter, err error) {
	var catalogErr e.CatalogError
	if errors.As(err, &catalogErr) {
		http.Error(w, catalogErr.Msg(), catalogErr.Code())
		c.Logger.Error("Request failed - err", err)
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		c.Logger.Error("Unexpected error - err", err)
	}
}

// respondJSON serializes the given data to JSON and writes it to the response.
// It sets the appropriate Content-Type header and handles encoding errors.
func (c *ActController) respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		c.handleError(w, e.NewInternal("Failed to encode response", err))
	}
}

// parsePaginationParams extracts and validates 'limit' and 'offset' query parameters.
// It returns default values (limit=10, offset=0) if parameters are missing.
// Returns an error if parameters are present but invalid.
func parsePaginationParams(query url.Values) (int64, int64, error) {
	// Default values
	limit := int64(10)
	offset := int64(0)

	// Parse 'limit' if present
	if l := query.Get("limit"); l != "" {
		parsedLimit, err := strconv.ParseInt(l, 10, 64)
		if err != nil || parsedLimit < 1 {
			return 0, 0, fmt.Errorf("invalid 'limit' parameter")
		}
		limit = parsedLimit
	}

	// Parse 'offset' if present
	if o := query.Get("offset"); o != "" {
		parsedOffset, err := strconv.ParseInt(o, 10, 64)
		if err != nil || parsedOffset < 0 {
			return 0, 0, fmt.Errorf("invalid 'offset' parameter")
		}
		offset = parsedOffset
	}

	return limit, offset, nil
}

// decodeJSON decodifica el cuerpo de la solicitud JSON en la estructura proporcionada.
func (c *ActController) decodeJSON(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		c.Logger.Error("Error decoding request body", "error", err)
		c.handleError(w, e.NewBadRequest("Invalid request payload", err))
		return false
	}
	return true
}

// withTimeout crea un contexto con un timeout de 5 segundos.
func (c *ActController) withTimeout(r *http.Request) (context.Context, context.CancelFunc) {
	return context.WithTimeout(r.Context(), 5*time.Second)
}
