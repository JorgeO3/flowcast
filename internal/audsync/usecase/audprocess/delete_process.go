package audprocess

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/audsync/repository"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// DeleteProcessInput holds the input data for the DeleteProcessUC usecase
type DeleteProcessInput struct {
	ID string
}

// DeleteProcessOutput holds the output data for the DeleteProcessUC usecase
type DeleteProcessOutput struct{}

// DeleteProcessUC is the usecase for creating a process
type DeleteProcessUC struct {
	Logger    logger.Interface
	Validator validator.Interface
	Consumer  redpanda.Consumer
	Repos     *repository.Repositories
}

// DeleteProcessUCOpts is a type for the options of the DeleteProcessUC usecase
type DeleteProcessUCOpts func(*DeleteProcessUC)

// WithDeleteProcessLogger sets the logger in the DeleteProcessUC usecase
func WithDeleteProcessLogger(logger logger.Interface) DeleteProcessUCOpts {
	return func(cp *DeleteProcessUC) {
		cp.Logger = logger
	}
}

// WithDeleteProcessValidator sets the validator in the DeleteProcessUC usecase
func WithDeleteProcessValidator(validator validator.Interface) DeleteProcessUCOpts {
	return func(cp *DeleteProcessUC) {
		cp.Validator = validator
	}
}

// WithDeleteProcessConsumer sets the consumer in the DeleteProcessUC usecase
func WithDeleteProcessConsumer(consumer redpanda.Consumer) DeleteProcessUCOpts {
	return func(cp *DeleteProcessUC) {
		cp.Consumer = consumer
	}
}

// WithDeleteProcessRepos sets the repositories in the DeleteProcessUC usecase
func WithDeleteProcessRepos(repos *repository.Repositories) DeleteProcessUCOpts {
	return func(cp *DeleteProcessUC) {
		cp.Repos = repos
	}
}

// NewDeleteProcessUC Deletes a new DeleteProcessUC usecase
func NewDeleteProcessUC(opts ...DeleteProcessUCOpts) *DeleteProcessUC {
	uc := &DeleteProcessUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute executes the DeleteProcessUC usecase
func (cp *DeleteProcessUC) Execute(ctx context.Context, input *DeleteProcessInput) (*DeleteProcessOutput, error) {
	return &DeleteProcessOutput{}, nil
}
