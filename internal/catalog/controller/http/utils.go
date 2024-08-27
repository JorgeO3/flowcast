package http

import (
	"encoding/json"
	"net/http"

	"github.com/JorgeO3/flowcast/internal/catalog/errors"
)

func (c *Controller) handleError(w http.ResponseWriter, err error) {
	if catalogErr, ok := err.(errors.CatalogError); ok {
		http.Error(w, catalogErr.Msg(), catalogErr.Code())
		c.Logger.Error("Request failed")
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		c.Logger.Error("Unexpected error - %s", err.Error())
	}
}

func (c *Controller) respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		c.handleError(w, errors.NewInternal("Failed to encode response", err))
	}
}
