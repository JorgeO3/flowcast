// Package usecase provides the use cases for the catalog service.
package usecase

import (
	"context"
	"time"

	e "github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/infrastructure/kafka"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
	"github.com/JorgeO3/flowcast/internal/catalog/utils"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/validator"
	"github.com/google/uuid"
)

// CreateActInput represents the input for the CreateAct use case.
type CreateActInput struct {
	Act e.Act `json:"act" validate:"required"`
}

// CreateActOutput represents the output for the CreateAct use case.
type CreateActOutput struct {
	AssetsURLs       []utils.AssetURL    `json:"assets,omitempty"`
	ProcessingAssets []AudioServiceAsset `json:"processingAssets,omitempty"`
}

// AudioServiceAsset represents a asset in the audio service.
type AudioServiceAsset struct {
	EntityType          string    `json:"entityType"`
	AlbumID             string    `json:"albumId"`
	AlbumName           string    `json:"albumName"`
	SongID              string    `json:"songId"`
	AssetID             string    `json:"assetId"`
	FilePath            string    `json:"filePath"`
	AssetType           string    `json:"assetType"`
	AssetName           string    `json:"assetName"`
	Status              string    `json:"status"`
	ProcessingStartTime time.Time `json:"processingStartTime"`
	ProcessingEndTime   time.Time `json:"processingEndTime"`
	ErrorMsg            string    `json:"errorMsg"`
	UserID              string    `json:"userId"`
	ActID               string    `json:"actId"`
	ActName             string    `json:"actName"`
}

// CreateActEvent represents an audio event.
type CreateActEvent struct {
	EventID string              `json:"eventId"`
	Assets  []AudioServiceAsset `json:"assets"`
}

// CreateActUC is the use case for creating an musical actor.
type CreateActUC struct {
	Logger    logger.Interface
	Validator validator.Interface
	Producer  *kafka.Producer
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
func WithCreateActProducer(producer *kafka.Producer) CreateActUCOpts {
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

func handleCreatedAssets(output *utils.AssetsProcessorOutput) []AudioServiceAsset {
	createdAssets := make([]AudioServiceAsset, len(output.DeletedAssets))
	for _, asset := range output.DeletedAssets {
		createdAssets = append(createdAssets, AudioServiceAsset{
			UserID:              asset.UserID,
			ActID:               asset.ActID,
			ActName:             asset.ActName,
			AlbumID:             asset.AlbumID,
			SongID:              asset.SongID,
			AssetID:             uuid.New().String(),
			AssetType:           string(asset.Type),
			AssetName:           asset.NewAsset.Name,
			AlbumName:           asset.AlbumName,
			EntityType:          string(asset.EntityType),
			Status:              "processing",
			ProcessingStartTime: time.Now(),
			ProcessingEndTime:   time.Now(),
			FilePath:            asset.NewAsset.URL,
			ErrorMsg:            "",
		})
	}
	return createdAssets
}

// Execute executes the CreateAct use case.
func (uc *CreateActUC) Execute(ctx context.Context, input CreateActInput) (*CreateActOutput, error) {
	uc.Logger.Info("Creating a new musical act")

	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warn("Invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	if _, err := uc.Repos.Act.CreateAct(ctx, &input.Act); err != nil {
		uc.Logger.Error("Error inserting act in db: %v", err)
		return nil, errors.HandleRepoError(err)
	}

	processor := utils.NewAssetsProcessor(ctx, uc.Repos)
	output, err := processor.Create(&input.Act)
	if err != nil {
		uc.Logger.Error("Error processing assets: %v", err)
		return nil, err
	}

	createdAssets := handleCreatedAssets(output)

	// Create Act event
	event := CreateActEvent{
		EventID: uuid.New().String(),
		Assets:  createdAssets,
	}

	if err := uc.Producer.Publish(ctx, event, e.CreateActTopic); err != nil {
		uc.Logger.Error("Error producing event: %v", err)
		return nil, errors.NewInternal("error producing event", err)
	}

	uc.Logger.Info("Act created successfully")
	return &CreateActOutput{
		ProcessingAssets: createdAssets,
		AssetsURLs:       output.AssetsURLs,
	}, nil
}
