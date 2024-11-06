package audprocess

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/audsync/repository"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// GetManyProcessInput holds the input data for the GetManyProcessUC usecase
type GetManyProcessInput struct {
	Limit  int64
	Offset int64
}

// GetManyProcessOutput holds the output data for the GetManyProcessUC usecase
type GetManyProcessOutput struct{}

// GetManyProcessUC is the usecase for creating a process
type GetManyProcessUC struct {
	Logger    logger.Interface
	Validator validator.Interface
	Consumer  redpanda.Consumer
	Repos     *repository.Repositories
}

// GetManyProcessUCOpts is a type for the options of the GetManyProcessUC usecase
type GetManyProcessUCOpts func(*GetManyProcessUC)

// WithGetManyProcessLogger sets the logger in the GetManyProcessUC usecase
func WithGetManyProcessLogger(logger logger.Interface) GetManyProcessUCOpts {
	return func(cp *GetManyProcessUC) {
		cp.Logger = logger
	}
}

// WithGetManyProcessValidator sets the validator in the GetManyProcessUC usecase
func WithGetManyProcessValidator(validator validator.Interface) GetManyProcessUCOpts {
	return func(cp *GetManyProcessUC) {
		cp.Validator = validator
	}
}

// WithGetManyProcessConsumer sets the consumer in the GetManyProcessUC usecase
func WithGetManyProcessConsumer(consumer redpanda.Consumer) GetManyProcessUCOpts {
	return func(cp *GetManyProcessUC) {
		cp.Consumer = consumer
	}
}

// WithGetManyProcessRepos sets the repositories in the GetManyProcessUC usecase
func WithGetManyProcessRepos(repos *repository.Repositories) GetManyProcessUCOpts {
	return func(cp *GetManyProcessUC) {
		cp.Repos = repos
	}
}

// NewGetManyProcessUC GetManys a new GetManyProcessUC usecase
func NewGetManyProcessUC(opts ...GetManyProcessUCOpts) *GetManyProcessUC {
	uc := &GetManyProcessUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute executes the GetManyProcessUC usecase
func (cp *GetManyProcessUC) Execute(ctx context.Context, input *GetManyProcessInput) (*GetManyProcessOutput, error) {
	return &GetManyProcessOutput{}, nil
}
