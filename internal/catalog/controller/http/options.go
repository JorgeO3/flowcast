package http

import (
	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/internal/catalog/usecase"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
)

// Opts represents the functional options for the controller.
type Opts func(c *Controller)

// WithCreateActUC sets the CreateActUC in the controller.
func WithCreateActUC(uc *usecase.CreateActUC) Opts {
	return func(c *Controller) {
		c.CreateActUC = uc
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