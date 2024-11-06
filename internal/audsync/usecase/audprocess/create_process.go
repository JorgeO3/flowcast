// Package audprocess provides all the usecases for syncronizing the catalog service with the audio processing service.
package audprocess

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/audsync/repository"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// CreateProcessInput holds the input data for the CreateProcessUC usecase
type CreateProcessInput struct{}

// CreateProcessOutput holds the output data for the CreateProcessUC usecase
type CreateProcessOutput struct{}

// CreateProcessUC is the usecase for creating a process
type CreateProcessUC struct {
	Logger    logger.Interface
	Validator validator.Interface
	Consumer  redpanda.Consumer
	Repos     *repository.Repositories
}

// CreateProcessUCOpts is a type for the options of the CreateProcessUC usecase
type CreateProcessUCOpts func(*CreateProcessUC)

// WithCreateProcessLogger sets the logger in the CreateProcessUC usecase
func WithCreateProcessLogger(logger logger.Interface) CreateProcessUCOpts {
	return func(cp *CreateProcessUC) {
		cp.Logger = logger
	}
}

// WithCreateProcessValidator sets the validator in the CreateProcessUC usecase
func WithCreateProcessValidator(validator validator.Interface) CreateProcessUCOpts {
	return func(cp *CreateProcessUC) {
		cp.Validator = validator
	}
}

// WithCreateProcessConsumer sets the consumer in the CreateProcessUC usecase
func WithCreateProcessConsumer(consumer redpanda.Consumer) CreateProcessUCOpts {
	return func(cp *CreateProcessUC) {
		cp.Consumer = consumer
	}
}

// WithCreateProcessRepos sets the repositories in the CreateProcessUC usecase
func WithCreateProcessRepos(repos *repository.Repositories) CreateProcessUCOpts {
	return func(cp *CreateProcessUC) {
		cp.Repos = repos
	}
}

// NewCreateProcessUC creates a new CreateProcessUC usecase
func NewCreateProcessUC(opts ...CreateProcessUCOpts) *CreateProcessUC {
	uc := &CreateProcessUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute executes the CreateProcessUC usecase
func (cp *CreateProcessUC) Execute(ctx context.Context, input *CreateProcessInput) (*CreateProcessOutput, error) {
	return &CreateProcessOutput{}, nil
}
