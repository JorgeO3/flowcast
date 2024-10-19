package usecase

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	e "github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
	"github.com/JorgeO3/flowcast/internal/catalog/utils"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/mongotx"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// UpdateActInput represents the input required to update a musical act.
type UpdateActInput struct {
	Act    e.Act  `json:"act" validate:"required,dive"`
	UserID string `json:"userId" validate:"required"`
}

// UpdateActOutput represents the result of updating a musical act.
type UpdateActOutput struct {
	ID       string          `json:"id,omitempty"`
	SongURLs []utils.SongURL `json:"songURLs,omitempty"`
	ImageURL string          `json:"imageUrl,omitempty"`
}

// BaseEvent defines the common structure for all events.
type BaseEvent struct {
	EventID   string    `json:"eventId"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}

// UpdateActEvent represents an event triggered after updating an act.
type UpdateActEvent struct {
	BaseEvent
	DeletedSongs []string `json:"deletedSongs,omitempty"` // Names of deleted songs
	AddedSongs   []string `json:"addedSongs,omitempty"`   // Names of added songs
}

// UpdateActUC is the use case responsible for updating a musical act.
type UpdateActUC struct {
	repos     *repository.Repositories
	logger    logger.Interface
	validator validator.Interface

	producer       redpanda.Producer
	transactionMgr mongotx.TxManager
}

// UpdateActUCOpt represents a functional option for configuring UpdateActUC.
type UpdateActUCOpt func(*UpdateActUC)

// optionHelper is a helper function to create functional options.
func optionHelper(setter func(*UpdateActUC)) UpdateActUCOpt {
	return func(uc *UpdateActUC) {
		setter(uc)
	}
}

// WithUpdateRepositories sets the repositories in the UpdateActUC.
func WithUpdateRepositories(repo *repository.Repositories) UpdateActUCOpt {
	return func(uc *UpdateActUC) {
		uc.repos = repo
	}
}

// WithUpdateActLogger sets the logger in the UpdateActUC.
func WithUpdateActLogger(logger logger.Interface) UpdateActUCOpt {
	return func(uc *UpdateActUC) {
		uc.logger = logger
	}
}

// WithUpdateActValidator sets the Validator in UpdateActUC.
func WithUpdateActValidator(v validator.Interface) UpdateActUCOpt {
	return optionHelper(func(uc *UpdateActUC) {
		uc.validator = v
	})
}

// WithUpdateActProducer sets the Producer in UpdateActUC.
func WithUpdateActProducer(producer redpanda.Producer) UpdateActUCOpt {
	return func(uc *UpdateActUC) {
		uc.producer = producer
	}
}

// WithUpdateActTransactionManager sets the TransactionManager in UpdateActUC.
func WithUpdateActTransactionManager(txMgr mongotx.TxManager) UpdateActUCOpt {
	return func(uc *UpdateActUC) {
		uc.transactionMgr = txMgr
	}
}

// NewUpdateAct creates a new instance of UpdateActUC with the provided options.
func NewUpdateAct(opts ...UpdateActUCOpt) *UpdateActUC {
	uc := &UpdateActUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// getSongDifferences identifies deleted and added songs between existingSongs and updatedSongs.
func getSongDifferences(existingSongs, updatedSongs []e.Song) (deletedSongs, addedSongs []e.Song) {
	existingSongMap := make(map[primitive.ObjectID]e.Song, len(existingSongs))
	updatedSongMap := make(map[primitive.ObjectID]e.Song, len(updatedSongs))

	// Populate existing songs map
	for _, song := range existingSongs {
		existingSongMap[song.ID] = song
	}

	// Populate updated songs map
	for _, song := range updatedSongs {
		updatedSongMap[song.ID] = song
	}

	// Identify deleted songs
	for songID, song := range existingSongMap {
		if _, exists := updatedSongMap[songID]; !exists {
			deletedSongs = append(deletedSongs, song)
		}
	}

	// Identify added songs
	for songID, song := range updatedSongMap {
		if _, exists := existingSongMap[songID]; !exists {
			addedSongs = append(addedSongs, song)
		}
	}

	return
}

// handleAddedSongs processes the newly added songs by generating pre-signed URLs.
func (uc *UpdateActUC) handleAddedSongs(ctx context.Context, act *e.Act, addedSongs []e.Song) ([]utils.SongURL, error) {
	var songURLs []utils.SongURL

	// Create a map of song ID to album ID for faster lookup
	songToAlbumMap := make(map[primitive.ObjectID]string, len(act.Albums))
	for _, album := range act.Albums {
		for _, song := range album.Songs {
			songToAlbumMap[song.ID] = album.ID.Hex()
		}
	}

	for _, song := range addedSongs {
		albumID, exists := songToAlbumMap[song.ID]
		if !exists {
			uc.logger.Error("Album not found for song ID %s", song.ID.Hex())
			return nil, errors.NewInternal("album not found for song", nil)
		}

		url, err := utils.GenerateSongPresignedURL(ctx, "audio/", act.ID.Hex(), albumID, song.ID.Hex(), uc.repos.RawAudio)
		if err != nil {
			uc.logger.Error("Failed to generate pre-signed URL for song ID %s: %v", song.ID.Hex(), err)
			return nil, errors.NewInternal("failed to generate pre-signed URL", err)
		}
		songURLs = append(songURLs, utils.SongURL{URL: url, Name: song.File.Name})
	}

	return songURLs, nil
}

// handleDeletedSongs processes the songs that have been deleted.
// Currently, this function is a placeholder and should be implemented as needed.
func (uc *UpdateActUC) handleDeletedSongs(ctx context.Context, deletedSongs []e.Song) ([]utils.SongURL, error) {
	// Implement the necessary logic to handle deleted songs.
	// This could involve removing references, cleaning up resources, etc.
	return nil, nil
}

// handleImageUpdate processes the profile picture update by generating a pre-signed URL if necessary.
func (uc *UpdateActUC) handleImageUpdate(ctx context.Context, existingAct, updatedAct *e.Act) (string, error) {
	// Check if the profile picture has changed
	if existingAct.ProfilePicture != updatedAct.ProfilePicture && updatedAct.ProfilePicture != (e.Image{}) {
		imageURL, err := utils.GenerateImagePresignedURL(ctx, "images/", updatedAct.ID.Hex(), uc.repos.Assets)
		if err != nil {
			uc.logger.Error("Failed to generate image URL: %v", err)
			return "", errors.NewInternal("failed to generate image URL", err)
		}
		return imageURL, nil
	}
	return "", nil
}

func createSongToAlbumMap(act *e.Act) map[primitive.ObjectID]string {
	songToAlbumMap := make(map[primitive.ObjectID]string)
	for _, album := range act.Albums {
		for _, song := range album.Songs {
			songToAlbumMap[song.ID] = album.ID.Hex()
		}
	}
	return songToAlbumMap
}

func (uc *UpdateActUC) updateActTransaction(ctx context.Context, input UpdateActInput) (*UpdateActOutput, error) {
	existingAct, err := uc.repos.Act.GetActByID(ctx, input.Act.ID)
	if err != nil {
		return nil, errors.HandleRepoError(err)
	}

	deletedSongs, addedSongs := getSongDifferences(existingAct.GetSongs(), input.Act.GetSongs())

	if err := uc.repos.Act.UpdateAct(ctx, &input.Act); err != nil {
		return nil, errors.HandleRepoError(err)
	}

	imageURL, err := uc.handleImageUpdate(ctx, existingAct, &input.Act)
	if err != nil {
		return nil, err
	}

	// Generate pre-signed URLs for added songs
	songURLs, err := uc.handleAddedSongs(ctx, &input.Act, addedSongs)
	if err != nil {
		return nil, err
	}

	// Publish an event for the act update
	deletedSongURLs, err := uc.handleDeletedSongs(ctx, deletedSongs)

	return &UpdateActOutput{
		ID:       input.Act.ID.Hex(),
		SongURLs: songURLs,
		ImageURL: imageURL,
	}, nil
}

// Execute performs the UpdateAct use case, updating the act and handling associated changes.
func (uc *UpdateActUC) Execute(ctx context.Context, input UpdateActInput) (*UpdateActOutput, error) {
	uc.logger.Info("Starting act update in the catalog")

	// Validate the input
	if err := uc.validator.Validate(input); err != nil {
		uc.logger.Warn("Invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	var (
		deletedSongs []e.Song
		addedSongs   []e.Song
		songURLs     []utils.SongURL
		imageURL     string
	)

	output, err := uc

	// Execute the transaction to update the act
	err := uc.transactionMgr.Run(ctx, func(ctx context.Context) error {
		// Retrieve the existing act from the repository
		existingAct, err := uc.actRepo.GetActByID(ctx, input.Act.ID)
		if err != nil {
			uc.logger.Error("Failed to retrieve act: %v", err)
			return errors.HandleRepoError(err)
		}

		// Identify differences between existing songs and updated songs
		deletedSongs, addedSongs = getSongDifferences(existingAct.GetSongs(), input.Act.GetSongs())

		// Update the act in the repository
		if err := uc.actRepo.UpdateAct(ctx, &input.Act); err != nil {
			uc.logger.Error("Failed to update act: %v", err)
			return errors.HandleRepoError(err)
		}

		// Handle image update
		var imgURL string
		imgURL, err = uc.handleImageUpdate(ctx, existingAct, &input.Act)
		if err != nil {
			return err
		}
		imageURL = imgURL

		return nil
	})

	if err != nil {
		uc.logger.Error("Act update failed: %v", err)
		return nil, err
	}

	// Handle added songs
	if len(addedSongs) > 0 {
		addedSongURLs, err := uc.handleAddedSongs(ctx, &input.Act, addedSongs)
		if err != nil {
			return nil, err
		}
		songURLs = append(songURLs, addedSongURLs...)
	}

	// Handle deleted songs
	if len(deletedSongs) > 0 {
		deletedSongURLs, err := uc.handleDeletedSongs(ctx, deletedSongs)
		if err != nil {
			return nil, err
		}
		songURLs = append(songURLs, deletedSongURLs...)
	}

	uc.logger.Info("Act updated successfully")
	return &UpdateActOutput{
		ID:       input.Act.ID.Hex(),
		SongURLs: songURLs,
		ImageURL: imageURL,
	}, nil
}
