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

// CreateActsInput represents the input for the CreateActs use case.
type CreateActsInput struct {
	Acts []e.Act `json:"acts" validate:"required,dive,required"`
}

// CreateActsOutput represents the output for the CreateActs use case.
type CreateActsOutput struct {
	IDs              []string            `json:"ids,omitempty"`
	AssetURLs        []utils.AssetURL    `json:"assets,omitempty"`
	ProcessingAssets []AudioServiceAsset `json:"processingAssets,omitempty"`
}

// CreateActsEvent represents a song link.
type CreateActsEvent struct {
	EventID   string    `json:"eventId"`
	UserIDs   []string  `json:"userIds"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CreateActsUC is the use case for creating multiple musical actors.
type CreateActsUC struct {
	Logger    logger.Interface
	Validator validator.Interface
	Producer  redpanda.Producer
	Repos     *repository.Repositories
}

// CreateActsUCOpts represents the functional options for the CreateActsUC.
type CreateActsUCOpts func(uc *CreateActsUC)

// WithCreateActsLogger sets the logger in the CreateActsUC.
func WithCreateActsLogger(logger logger.Interface) CreateActsUCOpts {
	return func(uc *CreateActsUC) {
		uc.Logger = logger
	}
}

// WithCreateActsValidator sets the validator in the CreateActsUC.
func WithCreateActsValidator(validator validator.Interface) CreateActsUCOpts {
	return func(uc *CreateActsUC) {
		uc.Validator = validator
	}
}

// WithCreateActsProducer sets the producer in the CreateActsUC.
func WithCreateActsProducer(producer redpanda.Producer) CreateActsUCOpts {
	return func(uc *CreateActsUC) {
		uc.Producer = producer
	}
}

// NewCreateActs is the constructor for CreateActsUC use case
func NewCreateActs(opts ...CreateActsUCOpts) *CreateActsUC {
	uc := &CreateActsUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute executes the CreateActs use case
func (uc *CreateActsUC) Execute(ctx context.Context, input CreateActsInput) (*CreateActsOutput, error) {
	uc.Logger.Info("Creating multiple musical acts")

	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warn("Invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	ids, err := uc.Repos.Act.CreateActs(ctx, input.Acts)
	if err != nil {
		uc.Logger.Error("Error inserting acts in db: %v", err)
		return nil, errors.HandleRepoError(err)
	}

	processor := utils.NewAssetsProcessor(ctx, uc.Repos)
	output, err := processor.CreateMany(input.Acts)
	if err != nil {
		uc.Logger.Error("Error processing assets: %v", err)
		return nil, err
	}

	createdAssets := handleCreatedAssets(output)
	event := CreateActsEvent{
		EventID:   uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.Producer.Publish(ctx, event, e.CreateActsTopic); err != nil {
		uc.Logger.Error("Error producing event: %v", err)
		return nil, err
	}

	return &CreateActsOutput{IDs: ids, AssetURLs: output.AssetsURLs, ProcessingAssets: createdAssets}, nil
}
