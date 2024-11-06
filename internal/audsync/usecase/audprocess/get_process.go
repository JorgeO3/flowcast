package audprocess

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/audsync/repository"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// GetProcessInput holds the input data for the GetProcessUC usecase
type GetProcessInput struct {
	ID string
}

// GetProcessOutput holds the output data for the GetProcessUC usecase
type GetProcessOutput struct{}

// GetProcessUC is the usecase for creating a process
type GetProcessUC struct {
	Logger    logger.Interface
	Validator validator.Interface
	Consumer  redpanda.Consumer
	Repos     *repository.Repositories
}

// GetProcessUCOpts is a type for the options of the GetProcessUC usecase
type GetProcessUCOpts func(*GetProcessUC)

// WithGetProcessLogger sets the logger in the GetProcessUC usecase
func WithGetProcessLogger(logger logger.Interface) GetProcessUCOpts {
	return func(cp *GetProcessUC) {
		cp.Logger = logger
	}
}

// WithGetProcessValidator sets the validator in the GetProcessUC usecase
func WithGetProcessValidator(validator validator.Interface) GetProcessUCOpts {
	return func(cp *GetProcessUC) {
		cp.Validator = validator
	}
}

// WithGetProcessConsumer sets the consumer in the GetProcessUC usecase
func WithGetProcessConsumer(consumer redpanda.Consumer) GetProcessUCOpts {
	return func(cp *GetProcessUC) {
		cp.Consumer = consumer
	}
}

// WithGetProcessRepos sets the repositories in the GetProcessUC usecase
func WithGetProcessRepos(repos *repository.Repositories) GetProcessUCOpts {
	return func(cp *GetProcessUC) {
		cp.Repos = repos
	}
}

// NewGetProcessUC Gets a new GetProcessUC usecase
func NewGetProcessUC(opts ...GetProcessUCOpts) *GetProcessUC {
	uc := &GetProcessUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute executes the GetProcessUC usecase
func (cp *GetProcessUC) Execute(ctx context.Context, input *GetProcessInput) (*GetProcessOutput, error) {
	return &GetProcessOutput{}, nil
}
