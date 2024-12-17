package usecase

import (
	"context"

	e "github.com/JorgeO3/flowcast/internal/catalogv2/domain/entity"
	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/errors"
	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/repository"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/eventbus"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/logger"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/utils"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/validator"
)

// CreateManyInput represents the input for the CreateMany use case.
type CreateManyInput struct {
	Acts []e.Act `json:"acts" validate:"required,dive,required"`
}

// CreateManyOutput represents the output for the CreateMany use case.
type CreateManyOutput struct {
	IDs              []string            `json:"ids,omitempty"`
	AssetURLs        []utils.AssetURL    `json:"assets,omitempty"`
	ProcessingAssets []AudioServiceAsset `json:"processingAssets,omitempty"`
}

// CreateManyUC is the use case for creating multiple musical actors.
type CreateManyUC struct {
	Logger       logger.Interface
	Validator    validator.Interface
	Producer     eventbus.Producer
	ActRepo      repository.ActRepository
	AssetsRepo   repository.AssetsRepository
	RawaudioRepo repository.RawaudioRepository
}

// CreateManyUCOpts represents the functional options for the CreateManyUC.
type CreateManyUCOpts func(uc *CreateManyUC)

// WithCreateManyLogger sets the logger in the CreateManyUC.
func WithCreateManyLogger(logger logger.Interface) CreateManyUCOpts {
	return func(uc *CreateManyUC) {
		uc.Logger = logger
	}
}

// WithCreateManyValidator sets the validator in the CreateManyUC.
func WithCreateManyValidator(validator validator.Interface) CreateManyUCOpts {
	return func(uc *CreateManyUC) {
		uc.Validator = validator
	}
}

// TODO: Fix the redpanda producer generic type

// WithCreateManyProducer sets the producer in the CreateManyUC.
func WithCreateManyProducer(producer eventbus.Producer) CreateManyUCOpts {
	return func(uc *CreateManyUC) {
		uc.Producer = producer
	}
}

// WithCreateManyActRepo sets the repositories in the CreateManyUC.
func WithCreateManyActRepo(repo repository.ActRepository) CreateManyUCOpts {
	return func(uc *CreateManyUC) {
		uc.ActRepo = repo
	}
}

// WithCreateManyAssetsRepo sets the repositories in the CreateManyUC.
func WithCreateManyAssetsRepo(repo repository.AssetsRepository) CreateManyUCOpts {
	return func(uc *CreateManyUC) {
		uc.AssetsRepo = repo
	}
}

// WithCreateManyRawaudioRepo sets the repositories in the CreateManyUC.
func WithCreateManyRawaudioRepo(repo repository.RawaudioRepository) CreateManyUCOpts {
	return func(uc *CreateManyUC) {
		uc.RawaudioRepo = repo
	}
}

// NewCreateMany is the constructor for CreateManyUC use case
func NewCreateMany(opts ...CreateManyUCOpts) *CreateManyUC {
	uc := &CreateManyUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute executes the CreateMany use case
func (uc *CreateManyUC) Execute(ctx context.Context, input CreateManyInput) (*CreateManyOutput, error) {
	uc.Logger.Info("Creating multiple musical acts")

	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warnf("Invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	// Create acts return IDs
	_, err := uc.ActRepo.CreateMany(ctx, input.Acts)
	if err != nil {
		uc.Logger.Errorf("Error inserting acts in db: %v", err)
		return nil, err
	}

	processor := NewAssetsProcessor(ctx, uc.ActRepo, uc.RawaudioRepo, uc.AssetsRepo)
	output, err := processor.CreateMany(input.Acts)
	if err != nil {
		uc.Logger.Errorf("Error processing assets: %v", err)
		return nil, err
	}

	_ = handleCreatedAssets(output)
	// createdAssets := handleCreatedAssets(output)
	// FIXME: Implement the event correctly
	event := struct{}{}

	if err := uc.Producer.Publish(ctx, event, ""); err != nil {
		uc.Logger.Errorf("Error producing event: %v", err)
		return nil, err
	}

	// FIXME: Implement the output correctly
	return &CreateManyOutput{}, nil
}
