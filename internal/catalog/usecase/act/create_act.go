// Package act provides the use cases for the catalog service.
package act

import (
	"context"

	e "github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/act"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/rawaudio"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// CreateActInput represents the input for the CreateAct use case.
type CreateActInput struct {
	e.Act
}

func (cai CreateActInput) toEntity() *e.Act {
	return e.NewAct(
		e.WithActID(cai.ID),
		e.WithActName(cai.Name),
		e.WithActType(cai.Type),
		e.WithActGenres(cai.Genres),
		e.WithActAlbums(cai.Albums),
		e.WithActMembers(cai.Members),
		e.WithActBiography(cai.Biography),
		e.WithActDisbandDate(cai.DisbandDate),
		e.WithActFormationDate(cai.FormationDate),
		e.WithActProfilePictureURL(cai.ProfilePictureURL),
	)
}

// CreateActOutput represents the output for the CreateAct use case.
type CreateActOutput struct {
	ID string
}

// CreateActUC is the use case for creating an musical actor.
type CreateActUC struct {
	ActRepository act.Repository
	Logger        logger.Interface
	Validator     validator.Interface
	RaRepository  rawaudio.Repository
}

// CreateActUCOpts represents the functional options for the CreateActUC.
type CreateActUCOpts func(uc *CreateActUC)

// WithCreateActRepository sets the ActRepository in the CreateActUC.
func WithCreateActRepository(repo act.Repository) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.ActRepository = repo
	}
}

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

// WithCreateActRARepository sets the RawAudioRepository in the CreateActUC.
func WithCreateActRARepository(repo rawaudio.Repository) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.RaRepository = repo
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

	id, err := uc.ActRepository.CreateAct(ctx, input.toEntity())
	if err != nil {
		uc.Logger.Error("Error inserting act in db: %v", err)
		return nil, errors.HandleRepoError(err)
	}

	uc.Logger.Info("Act created successfully")
	return &CreateActOutput{ID: id}, nil
}
