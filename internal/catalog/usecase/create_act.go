// Package usecase provides the use cases for the catalog service.
package usecase

import (
	"context"
	"time"

	e "github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/act"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/rawaudio"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// CreateActInput represents the input for the CreateAct use case.
type CreateActInput struct {
	Act e.Act `json:"act" validate:"required"`
}

// SongLink represents a song link.
type SongLink struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

// CreateActOutput represents the output for the CreateAct use case.
type CreateActOutput struct {
	ID        string     `json:"id"`
	SongLinks []SongLink `json:"song_links"`
}

// AudioEvent represents an audio event.
type AudioEvent struct {
	EventID             string            `json:"event_id"`
	FileName            string            `json:"file_name"`
	FileSize            int64             `json:"file_size"`
	FileType            string            `json:"file_type"`
	SourceLocation      string            `json:"source_location"`
	DestinationLocation string            `json:"destination_location"`
	ProcessingState     string            `json:"processing_state"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
	UserID              string            `json:"user_id"`
	Metadata            map[string]string `json:"metadata"`
}

// CreateActUC is the use case for creating an musical actor.
type CreateActUC struct {
	ActRepository act.Repository
	Logger        logger.Interface
	Validator     validator.Interface
	RaRepo        rawaudio.Repository
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

func generateSongPresignedURL(ctx context.Context, userID, albumID, songID string, RaRepo rawaudio.Repository) (string, error) {
	fileName := "raw-audio/" + userID + "/" + albumID + "/" + songID + ".wav"
	return RaRepo.GeneratePresignedURL(ctx, fileName, time.Minute)
}

func validateFile(file *e.AudioFile) error {
	if file.Ext != "wav" {
		return errors.NewValidation("invalid file extension", nil)
	}

	maxFileSize := 1024 * 1024 * 10 // 10MB
	if file.Size > maxFileSize {
		return errors.NewValidation("invalid file size", nil)
	}
	return nil
}

func generateSongLinks(ctx context.Context, act *e.Act, actID string, RaRepo rawaudio.Repository) ([]SongLink, error) {
	songsLength := act.SongsLength()
	songLinks := make([]SongLink, 0, songsLength)

	for _, album := range act.Albums {
		for _, song := range album.Songs {
			if err := validateFile(&song.File); err != nil {
				return nil, errors.NewValidation("invalid file", err)
			}

			url, err := generateSongPresignedURL(ctx, actID, album.ID.Hex(), song.ID.Hex(), RaRepo)
			if err != nil {
				return nil, errors.NewInternal("error generating presigned url", err)
			}

			songLinks = append(songLinks, SongLink{URL: url, Name: song.File.Name})
		}
	}

	return songLinks, nil
}

// Execute executes the CreateAct use case.
func (uc *CreateActUC) Execute(ctx context.Context, input CreateActInput) (*CreateActOutput, error) {
	uc.Logger.Info("Creating a new musical act")

	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warn("Invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	id, err := uc.ActRepository.CreateAct(ctx, &input.Act)
	if err != nil {
		uc.Logger.Error("Error inserting act in db: %v", err)
		return nil, errors.HandleRepoError(err)
	}

	songLinks, err := generateSongLinks(ctx, &input.Act, id, uc.RaRepo)
	if err != nil {
		uc.Logger.Error("Error generating song links: %v", err)
		return nil, err
	}

	// Postear evento en el topic de actos

	uc.Logger.Info("Act created successfully")
	return &CreateActOutput{ID: id, SongLinks: songLinks}, nil
}
