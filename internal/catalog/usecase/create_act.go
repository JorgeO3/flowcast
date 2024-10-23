// Package usecase provides the use cases for the catalog service.
package usecase

import (
	"context"
	"time"

	e "github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
	"github.com/JorgeO3/flowcast/internal/catalog/utils"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
	"github.com/JorgeO3/flowcast/pkg/validator"
	"github.com/google/uuid"
)

// CreateActInput represents the input for the CreateAct use case.
type CreateActInput struct {
	Act e.Act `json:"act" validate:"required"`
}

// CreateActOutput represents the output for the CreateAct use case.
type CreateActOutput struct {
	ID     string           `json:"id"`
	Assets []utils.AssetURL `json:"assets,omitempty"`
}

// CreateActEvent represents an audio event.
type CreateActEvent struct {
	UserID              string            `json:"userId"`
	EventID             string            `json:"eventId"`
	CreatedAt           time.Time         `json:"createdAt"`
	UpdatedAt           time.Time         `json:"updatedAt"`
	FileName            string            `json:"fileName"`
	FileSize            int64             `json:"fileSize"`
	FileType            string            `json:"fileType"`
	SourceLocation      string            `json:"sourceLocation"`
	DestinationLocation string            `json:"destinationLocation"`
	ProcessingState     string            `json:"processingState"`
	Metadata            map[string]string `json:"metadata"`
}

// CreateActUC is the use case for creating an musical actor.
type CreateActUC struct {
	Logger    logger.Interface
	Validator validator.Interface
	Producer  redpanda.Producer
	Repos     *repository.Repositories
}

// CreateActUCOpts represents the functional options for the CreateActUC.
type CreateActUCOpts func(uc *CreateActUC)

// WithCreateActLogger sets the logger in the CreateActUC.
func WithCreateActLogger(logger logger.Interface) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.Logger = logger
	}
}

// WithCreateActValidator sets the validator in the CreateActUC.
func WithCreateActValidator(validator validator.Interface) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.Validator = validator
	}
}

// WithCreateActProducer sets the Producer in the CreateActUC.
func WithCreateActProducer(producer redpanda.Producer) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.Producer = producer
	}
}

// WithCreateActRepos sets the Repos in the CreateActUC.
func WithCreateActRepos(repos *repository.Repositories) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.Repos = repos
	}
}

// NewCreateAct creates a new instance of CreateActUC.
func NewCreateAct(opts ...CreateActUCOpts) *CreateActUC {
	uc := &CreateActUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute executes the CreateAct use case.
func (uc *CreateActUC) Execute(ctx context.Context, input CreateActInput) (*CreateActOutput, error) {
	uc.Logger.Info("Creating a new musical act")

	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warn("Invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	processor := utils.NewAssetsProcessor(ctx, uc.Repos)
	output, err := processor.Create(&input.Act)
	if err != nil {
		uc.Logger.Error("Error processing assets: %v", err)
		return nil, err
	}

	id, err := uc.Repos.Act.CreateAct(ctx, &input.Act)
	if err != nil {
		uc.Logger.Error("Error inserting act in db: %v", err)
		return nil, errors.HandleRepoError(err)
	}

	// TODO: Finish the implementation of the event
	// Manejar la parte del evento
	event := CreateActEvent{
		UserID:    input.Act.UserID,
		EventID:   uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.Producer.Publish(ctx, event, e.CreateActTopic); err != nil {
		uc.Logger.Error("Error producing event: %v", err)
		return nil, errors.NewInternal("error producing event", err)
	}

	uc.Logger.Info("Act created successfully")
	return &CreateActOutput{Assets: output.AssetsURLs, ID: id}, nil
}
