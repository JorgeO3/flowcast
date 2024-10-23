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

// CreateActsInput represents the input for the CreateActs use case.
type CreateActsInput struct {
	Acts []*e.Act `json:"acts" validate:"required,dive,required"`
}

// CreateActsOutput represents the output for the CreateActs use case.
type CreateActsOutput struct {
	IDs    []string         `json:"ids,omitempty"`
	Assets []utils.AssetURL `json:"assets,omitempty"`
}

// CreateActsEvent represents a song link.
type CreateActsEvent struct {
	UserID    string    `json:"userId"`
	EventID   string    `json:"eventId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CreateActsUC is the use case for creating multiple musical actors.
type CreateActsUC struct {
	ActRepository act.Repository
	Logger        logger.Interface
	Validator     validator.Interface
	RaRepo        rawaudio.Repository
	AssRepo       assets.Repository
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

// WithCreateActsAssRepository sets the AssetsRepository in the CreateActsUC.
func WithCreateActsAssRepository(repo assets.Repository) CreateActsUCOpts {
	return func(uc *CreateActsUC) {
		uc.AssRepo = repo
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

	ids, err := uc.ActRepository.CreateActs(ctx, input.Acts)
	if err != nil {
		uc.Logger.Error("Error inserting acts in db: %v", err)
		return nil, errors.HandleRepoError(err)
	}

	// Generar las urls firmadas y retornarlas al usario
	songLinks, err := utils.GenerateSongURLsFromActs(ctx, input.Acts, "raw-audio/", uc.RaRepo)
	if err != nil {
		uc.Logger.Error("Error generating song links: %v", err)
		return nil, err
	}

	imageURLs, err := utils.GenerateImagePresignedURLsFromActs(ctx, input.Acts, "assets/", uc.AssRepo)
	if err != nil {
		uc.Logger.Error("Error generating image links: %v", err)
		return nil, err
	}

	// TODO: Definir evento de acuerdo al frontend
	event := CreateActsEvent{
		UserID:    "user_id",
		EventID:   "event_id",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.Producer.Publish(event, e.CreateActsTopic); err != nil {
		uc.Logger.Error("Error producing event: %v", err)
		return nil, err
	}

	return &CreateActsOutput{IDs: ids, SongURLs: songLinks, ImageURLs: imageURLs}, nil
}
