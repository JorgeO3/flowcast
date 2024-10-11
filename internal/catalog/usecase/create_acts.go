package usecase

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/act"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/rawaudio"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// CreateActsInput represents the input for the CreateActs use case.
type CreateActsInput struct {
	Acts []*entity.Act `json:"acts" validate:"required,dive,required"`
}

// CreateActsOutput represents the output for the CreateActs use case.
type CreateActsOutput struct {
	IDs []string `json:"ids"`
}

// CreateActsUC is the use case for creating multiple musical actors.
type CreateActsUC struct {
	ActRepository act.Repository
	Logger        logger.Interface
	Validator     validator.Interface
	RaRepo        rawaudio.Repository
	Producer      redpanda.Producer
}

// CreateActsUCOpts represents the functional options for the CreateActsUC.
type CreateActsUCOpts func(uc *CreateActsUC)

// WithCreateActsRepository sets the ActRepository in the CreateActsUC.
func WithCreateActsRepository(repo act.Repository) CreateActsUCOpts {
	return func(uc *CreateActsUC) {
		uc.ActRepository = repo
	}
}

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

// WithCreateActsRaRepository sets the RawAudioRepository in the CreateActsUC.
func WithCreateActsRaRepository(repo rawaudio.Repository) CreateActsUCOpts {
	return func(uc *CreateActsUC) {
		uc.RaRepo = repo
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

func (uc *CreateActsUC) generateManySongLinks(acts []*entity.Act) ([]SongLink, error) {
	var songLinks []SongLink

	for _, act := range acts {
		links, err := generateSongLinks(context.Background(), act, act.ID.Hex(), uc.RaRepo)
		if err != nil {
			return nil, err
		}
		songLinks = append(songLinks, links...)
	}

	return songLinks, nil
}

// Execute executes the CreateActs use case
func (uc *CreateActsUC) Execute(ctx context.Context, input CreateActsInput) (*CreateActsOutput, error) {
	uc.Logger.Info("Creating multiple musical acts")

	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warn("Invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	ids, err := uc.ActRepository.CreateActs(ctx, input.Acts)
	if err != nil {
		uc.Logger.Error("Error inserting acts in db: %v", err)
		return nil, errors.HandleRepoError(err)
	}

	// Generar las urls firmadas y retornarlas al usario

	return &CreateActsOutput{IDs: ids}, nil
}
