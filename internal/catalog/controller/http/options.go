// Package http provides the HTTP Controller for the catalog service.
package http

import (
	"github.com/JorgeO3/flowcast/configs"
	"github.com/JorgeO3/flowcast/internal/catalog/usecase"
	"github.com/JorgeO3/flowcast/pkg/logger"
)

// Opts represents the functional options for the controller.
type Opts func(c *Controller)

// WithCreateActUC sets the CreateActUC in the controller.
func WithCreateActUC(uc *usecase.CreateActUC) Opts {
	return func(c *Controller) {
		c.CreateActUC = uc
	}
}

// WithUpdateActUC sets the UpdateActUC in the controller.
func WithUpdateActUC(uc *usecase.UpdateActUC) Opts {
	return func(c *Controller) {
		c.UpdateActUC = uc
	}
}

// WithGetActByIDUC sets the GetActByIDUC in the controller.
func WithGetActByIDUC(uc *usecase.GetActByIDUC) Opts {
	return func(c *Controller) {
		c.GetActByIDUC = uc
	}
}

// WithDeleteActUC sets the DeleteActUC in the controller.
func WithDeleteActUC(uc *usecase.DeleteActUC) Opts {
	return func(c *Controller) {
		c.DeleteActUC = uc
	}
}

func WithGetActsUC(uc *usecase.GetActsUC) Opts {
	return func(c *Controller) {
		c.GetActsUC = uc
	}
}

// WithCreateManyUC sets the CreateManyUC in the controller.
func WithCreateManyUC(uc *usecase.CreateManyUC) Opts {
	return func(c *Controller) {
		c.CreateManyUC = uc
	}
}

// WithLogger sets the logger in the controller.
func WithLogger(l logger.Interface) Opts {
	return func(c *Controller) {
		c.Logger = l
	}
}

// WithConfig sets the config in the controller.
func WithConfig(cfg *configs.CatalogConfig) Opts {
	return func(c *Controller) {
		c.Cfg = cfg
	}
}

// New creates a new instance of Controller.
func New(opts ...Opts) *Controller {
	c := &Controller{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
