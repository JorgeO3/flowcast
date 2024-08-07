package controller

import (
	"encoding/json"
	"net/http"

	"gitlab.com/JorgeO3/flowcast/internal/auth/errors"
)

func (c *Controller) handleError(w http.ResponseWriter, err error) {
	if authErr, ok := err.(errors.AuthError); ok {
		http.Error(w, authErr.Message(), authErr.Code())
		c.Log.Error("Request failed", "error", authErr.Error())
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		c.Log.Error("Unexpected error", "error", err.Error())
	}
}

func (c *Controller) respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		c.handleError(w, errors.NewInternal("Failed to encode response", err))
	}
}
