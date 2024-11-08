// Package audprocess provides all the usecases for syncronizing the catalog service with the audio processing service.
package audprocess

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/audsync/events"
	"github.com/JorgeO3/flowcast/internal/audsync/repository"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// UpdateProcessInput holds the input data for the UpdateProcessUC usecase
type UpdateProcessInput struct {
	events.UpdateAudioProcessingsEvent
}

// UpdateProcessOutput holds the output data for the UpdateProcessUC usecase
type UpdateProcessOutput struct{}

// UpdateProcessUC is the usecase for creating a process
type UpdateProcessUC struct {
	Logger    logger.Interface
	Validator validator.Interface
	Repos     *repository.Repositories
}

// UpdateProcessUCOpts is a type for the options of the UpdateProcessUC usecase
type UpdateProcessUCOpts func(*UpdateProcessUC)

// WithUpdateProcessLogger sets the logger in the UpdateProcessUC usecase
func WithUpdateProcessLogger(logger logger.Interface) UpdateProcessUCOpts {
	return func(cp *UpdateProcessUC) {
		cp.Logger = logger
	}
}

// WithUpdateProcessValidator sets the validator in the UpdateProcessUC usecase
func WithUpdateProcessValidator(validator validator.Interface) UpdateProcessUCOpts {
	return func(cp *UpdateProcessUC) {
		cp.Validator = validator
	}
}

// WithUpdateProcessRepos sets the repositories in the UpdateProcessUC usecase
func WithUpdateProcessRepos(repos *repository.Repositories) UpdateProcessUCOpts {
	return func(cp *UpdateProcessUC) {
		cp.Repos = repos
	}
}

// UpdateateProcessUC creates a new UpdateProcessUC usecase
func NewUpdateProcessUC(opts ...UpdateProcessUCOpts) *UpdateProcessUC {
	uc := &UpdateProcessUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute executes the UpdateProcessUC usecase
func (cp *UpdateProcessUC) Execute(ctx context.Context, input *UpdateProcessInput) (*UpdateProcessOutput, error) {
	return &UpdateProcessOutput{}, nil
}
