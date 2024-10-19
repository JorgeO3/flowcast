// Package usecase provides the use cases for the catalog service.
package usecase

import (
	"context"
	"time"

	e "github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/act"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/assets"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/rawaudio"
	"github.com/JorgeO3/flowcast/internal/catalog/utils"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// CreateActInput represents the input for the CreateAct use case.
type CreateActInput struct {
	Act e.Act `json:"act" validate:"required"`
}

// CreateActOutput represents the output for the CreateAct use case.
type CreateActOutput struct {
	ID       string          `json:"id"`
	SongURLs []utils.SongURL `json:"songURLs"`
	ImageURL string          `json:"imageUrl"`
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
	ActRepository act.Repository
	Logger        logger.Interface
	Validator     validator.Interface
	RaRepo        rawaudio.Repository
	AssRepo       assets.Repository
	Producer      redpanda.Producer
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
		uc.RaRepo = repo
	}
}

// WithCreateActAssRepository sets the AssetsRepository in the CreateActUC.
func WithCreateActAssRepository(repo assets.Repository) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.AssRepo = repo
	}
}

// WithCreateActProducer sets the Producer in the CreateActUC.
func WithCreateActProducer(producer redpanda.Producer) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.Producer = producer
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

	act := &input.Act
	act.GenerateIDs()

	id, err := uc.ActRepository.CreateAct(ctx, act)
	if err != nil {
		uc.Logger.Error("Error inserting act in db: %v", err)
		return nil, errors.HandleRepoError(err)
	}

	// Generar URL de las canciones
	songURLs, err := utils.GenerateSongURLsFromAct(ctx, "raw-audio/", act, uc.RaRepo)
	if err != nil {
		uc.Logger.Error("Error generating song links: %v", err)
		return nil, err
	}

	// Generar URL de la imagen
	profilePictureURL, err := utils.GenerateImagePresignedURL(ctx, "/images", act.ID.Hex(), uc.RaRepo)
	if err != nil {
		uc.Logger.Error("Error generating image url: %v", err)
		return nil, errors.NewInternal("error generating image url", err)
	}

	// Manejar la parte del evento

	uc.Logger.Info("Act created successfully")
	return &CreateActOutput{ID: id, SongURLs: songURLs, ImageURL: profilePictureURL}, nil
}
