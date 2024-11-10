// Package utils provides utility functions for the catalog service.
package usecase

import (
	"context"
	"fmt"
	"slices"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/utils"
)

// Action represents the action to be performed
type Action string

const (
	// Add reflects the addition of an asset
	Add Action = "add"
	// Update reflects the update of an asset
	Update Action = "update"
	// Delete reflects the deletion of an asset
	Delete Action = "delete"
)

// EntityType represents the type of entity
type EntityType string

const (
	// Act represents an act entity
	Act EntityType = "act"
	// Album represents an album entity
	Album EntityType = "album"
	// Song represents a song entity
	Song EntityType = "song"
)

// AssetChange represents a change in an asset
type AssetChange struct {
	Action     Action
	UserID     string
	ActID      string
	ActName    string
	AlbumID    string
	AlbumName  string
	SongID     string
	Type       entity.AssetType
	EntityType EntityType
	EntityID   string
	OldAsset   *entity.Asset
	NewAsset   *entity.Asset
}

// AssetsProcessorParams represents the parameters needed by the AssetsProcessor
type AssetsProcessorParams struct {
	Ctx   context.Context
	Repos *repository.Repositories
}

// AssetsProcessorOutput represents the output of the AssetsProcessor
type AssetsProcessorOutput struct {
	AddedAssets   []AssetChange
	DeletedAssets []AssetChange
	AssetsURLs    []utils.AssetURL
}

// AssetsProcessor is responsible for processing changes in assets
type AssetsProcessor struct {
	ctx   context.Context
	repos *repository.Repositories
}

// NewAssetsProcessor is a constructor for AssetsProcessor
func NewAssetsProcessor(ctx context.Context, repos *repository.Repositories) *AssetsProcessor {
	return &AssetsProcessor{
		ctx:   ctx,
		repos: repos,
	}
}

// Create processes assets for the creation of an Act
func (processor *AssetsProcessor) Create(act *entity.Act) (*AssetsProcessorOutput, error) {
	assets, err := processor.processAssets(nil, act)
	if err != nil {
		return nil, err
	}

	return processor.handleAssets(assets)
}

// CreateMany processes assets for the creation of multiple Acts
func (processor *AssetsProcessor) CreateMany(acts []entity.Act) (*AssetsProcessorOutput, error) {
	var changedAssets []AssetChange

	for _, act := range acts {
		assets, err := processor.processAssets(nil, &act)
		if err != nil {
			return nil, err
		}

		changedAssets = slices.Grow(changedAssets, len(assets))
		changedAssets = append(changedAssets, assets...)
	}

	return processor.handleAssets(changedAssets)
}

// Update processes assets for updating an Act
func (processor *AssetsProcessor) Update(oldAct, newAct *entity.Act) (*AssetsProcessorOutput, error) {
	assets, err := processor.processAssets(oldAct, newAct)
	if err != nil {
		return nil, err
	}

	return processor.handleAssets(assets)
}

// Delete processes assets for the deletion of an Act
func (processor *AssetsProcessor) Delete(act *entity.Act) (*AssetsProcessorOutput, error) {
	assets, err := processor.processAssets(act, nil)
	if err != nil {
		return nil, err
	}

	return processor.handleAssets(assets)
}

// processAssets handles the processing of assets for create, update, and delete operations
func (processor *AssetsProcessor) processAssets(oldAct, newAct *entity.Act) ([]AssetChange, error) {
	var assetsToProcess []AssetChange

	if oldAct == nil && newAct != nil {
		// Creating a new Act
		GenerateIDs(newAct)
		GenerateURLs(newAct)
		assetsToProcess = processor.collectAssetsForCreation(newAct)
	} else if newAct == nil && oldAct != nil {
		// Deleting a new act
		assetsToProcess = processor.collectAssetsForDeletion(oldAct)
	} else if newAct != nil && oldAct != nil {
		// Updating an existing act
		GenerateURLs(newAct)
		assetsToProcess = processor.collectAssetsForUpdate(oldAct, newAct)
	} else {
		return nil, fmt.Errorf("invalid input: both newAct and oldAct are nil")
	}

	return assetsToProcess, nil
}

// collectAssetsForCreation collects assets for creation
func (processor *AssetsProcessor) collectAssetsForCreation(act *entity.Act) []AssetChange {
	var assets []AssetChange

	// Profile picture
	if !isAssetEmpty(act.ProfilePicture) {
		assets = append(assets, AssetChange{
			Action:     Add,
			EntityType: Act,
			ActID:      act.ID,
			UserID:     act.UserID,
			ActName:    act.Name,
			EntityID:   act.ID,
			Type:       act.ProfilePicture.Type,
			NewAsset:   &act.ProfilePicture,
		})
	}

	// Albums and songs
	for _, album := range act.Albums {
		assets = append(assets, processor.collectAlbumAssetsForCreation(act, album)...)
	}

	return assets
}

// collectAssetsForDeletion collects assets for deletion
func (processor *AssetsProcessor) collectAssetsForDeletion(act *entity.Act) []AssetChange {
	var assets []AssetChange

	// Profile picture
	if !isAssetEmpty(act.ProfilePicture) {
		assets = append(assets, AssetChange{
			Action:     Delete,
			UserID:     act.UserID,
			EntityType: Act,
			ActID:      act.ID,
			EntityID:   act.ID,
			Type:       act.ProfilePicture.Type,
			OldAsset:   &act.ProfilePicture,
		})
	}

	// Albums and songs
	for _, album := range act.Albums {
		assets = append(assets, processor.collectAlbumAssetsForDeletion(act, album)...)
	}

	return assets
}

// collectAssetsForUpdate collects assets for update
func (processor *AssetsProcessor) collectAssetsForUpdate(oldAct, newAct *entity.Act) []AssetChange {
	var assets []AssetChange

	// Profile picture
	if !assetsEqual(newAct.ProfilePicture, oldAct.ProfilePicture) {
		assets = append(assets, AssetChange{
			Action:     Update,
			EntityType: Act,
			ActID:      newAct.ID,
			UserID:     newAct.UserID,
			EntityID:   newAct.ID,
			Type:       newAct.ProfilePicture.Type,
			NewAsset:   &newAct.ProfilePicture,
			OldAsset:   &oldAct.ProfilePicture,
		})
	}

	// Albums and songs
	assets = append(assets, processor.compareAlbums(oldAct, newAct)...)

	return assets
}

// collectAlbumAssetsForCreation collects album assets for creation
func (processor *AssetsProcessor) collectAlbumAssetsForCreation(act *entity.Act, album entity.Album) []AssetChange {
	var assets []AssetChange

	// Album cover art
	if !isAssetEmpty(album.CoverArt) {
		assets = append(assets, AssetChange{
			Action:     Add,
			EntityType: Album,
			ActID:      act.ID,
			UserID:     act.UserID,
			AlbumID:    album.ID,
			EntityID:   album.ID,
			Type:       album.CoverArt.Type,
			NewAsset:   &album.CoverArt,
		})
	}

	// Songs
	for _, song := range album.Songs {
		// If the song has a file
		if !isAssetEmpty(song.File) {
			assets = append(assets, AssetChange{
				Action:     Add,
				EntityType: Song,
				ActID:      act.ID,
				UserID:     act.UserID,
				AlbumID:    album.ID,
				SongID:     song.ID,
				EntityID:   song.ID,
				Type:       song.File.Type,
				NewAsset:   &song.File,
			})
		}

		// If the song has cover art
		if !isAssetEmpty(song.CoverArt) {
			assets = append(assets, AssetChange{
				Action:     Add,
				EntityType: Song,
				ActID:      act.ID,
				UserID:     act.UserID,
				AlbumID:    album.ID,
				SongID:     song.ID,
				EntityID:   song.ID,
				Type:       song.CoverArt.Type,
				NewAsset:   &song.CoverArt,
			})
		}
	}

	return assets
}

// collectAlbumAssetsForDeletion collects album assets for deletion
func (processor *AssetsProcessor) collectAlbumAssetsForDeletion(act *entity.Act, album entity.Album) []AssetChange {
	var assets []AssetChange

	// Album cover art
	if !isAssetEmpty(album.CoverArt) {
		assets = append(assets, AssetChange{
			Action:     Delete,
			EntityType: Album,
			UserID:     act.UserID,
			ActID:      act.ID,
			AlbumID:    album.ID,
			EntityID:   album.ID,
			Type:       album.CoverArt.Type,
			OldAsset:   &album.CoverArt,
		})
	}

	// Songs
	for _, song := range album.Songs {
		assets = append(assets, processor.collectSongAssetsForDeletion(act, album.ID, song)...)
	}

	return assets
}

// compareAlbums compares albums between the old and new Act
func (processor *AssetsProcessor) compareAlbums(oldAct, newAct *entity.Act) []AssetChange {
	var assets []AssetChange

	oldAlbumMap := generateAlbumMap(oldAct.Albums)
	newAlbumMap := generateAlbumMap(newAct.Albums)

	// Albums added or updated
	for _, newAlbum := range newAct.Albums {
		oldAlbum, exists := oldAlbumMap[newAlbum.ID]
		if !exists {
			// New album
			assets = append(assets, processor.collectAlbumAssetsForCreation(newAct, newAlbum)...)
			continue
		}

		// Album cover art
		if !assetsEqual(newAlbum.CoverArt, oldAlbum.CoverArt) {
			assets = append(assets, AssetChange{
				Action:     Update,
				EntityType: Album,
				ActID:      newAct.ID,
				UserID:     newAct.UserID,
				AlbumID:    newAlbum.ID,
				EntityID:   newAlbum.ID,
				Type:       newAlbum.CoverArt.Type,
				NewAsset:   &newAlbum.CoverArt,
				OldAsset:   &oldAlbum.CoverArt,
			})
		}

		// Songs
		assets = append(assets, processor.compareSongs(newAct, oldAlbum, newAlbum)...)
	}

	// Albums deleted
	for _, oldAlbum := range oldAct.Albums {
		if _, exists := newAlbumMap[oldAlbum.ID]; !exists {
			// Deleted album
			assets = append(assets, processor.collectAlbumAssetsForDeletion(oldAct, oldAlbum)...)
		}
	}

	return assets
}

// compareSongs compares songs between the old and new album
func (processor *AssetsProcessor) compareSongs(act *entity.Act, oldAlbum, newAlbum entity.Album) []AssetChange {
	var assets []AssetChange

	oldSongMap := generateSongMap(oldAlbum.Songs)
	newSongMap := generateSongMap(newAlbum.Songs)

	// Songs added or updated
	for _, newSong := range newAlbum.Songs {
		oldSong, exists := oldSongMap[newSong.ID]
		if !exists {
			// New song
			assets = append(assets, processor.collectSongAssetsForCreation(act, newAlbum.ID, newSong)...)
			continue
		}

		// Song file
		if !assetsEqual(newSong.File, oldSong.File) {
			assets = append(assets, AssetChange{
				Action:     Update,
				EntityType: Song,
				ActID:      act.ID,
				UserID:     act.UserID,
				AlbumID:    newAlbum.ID,
				SongID:     newSong.ID,
				EntityID:   newSong.ID,
				Type:       newSong.File.Type,
				NewAsset:   &newSong.File,
				OldAsset:   &oldSong.File,
			})
		}

		// Song cover art
		if !assetsEqual(newSong.CoverArt, oldSong.CoverArt) {
			assets = append(assets, AssetChange{
				Action:     Update,
				EntityType: Song,
				ActID:      act.ID,
				UserID:     act.UserID,
				AlbumID:    newAlbum.ID,
				SongID:     newSong.ID,
				EntityID:   newSong.ID,
				Type:       newSong.CoverArt.Type,
				NewAsset:   &newSong.CoverArt,
				OldAsset:   &oldSong.CoverArt,
			})
		}
	}

	// Songs deleted
	for _, oldSong := range oldAlbum.Songs {
		if _, exists := newSongMap[oldSong.ID]; !exists {
			// Deleted song
			assets = append(assets, processor.collectSongAssetsForDeletion(act, oldAlbum.ID, oldSong)...)
		}
	}

	return assets
}

// collectSongAssetsForCreation collects song assets for creation
func (processor *AssetsProcessor) collectSongAssetsForCreation(act *entity.Act, albumID string, song entity.Song) []AssetChange {
	var assets []AssetChange

	assets = append(assets, AssetChange{
		Action:     Add,
		EntityType: Song,
		ActID:      act.ID,
		UserID:     act.UserID,
		AlbumID:    albumID,
		SongID:     song.ID,
		EntityID:   song.ID,
		Type:       song.File.Type,
		NewAsset:   &song.File,
	})
	// If the song has cover art
	if !isAssetEmpty(song.CoverArt) {
		assets = append(assets, AssetChange{
			Action:     Add,
			EntityType: Song,
			ActID:      act.ID,
			UserID:     act.UserID,
			AlbumID:    albumID,
			SongID:     song.ID,
			EntityID:   song.ID,
			Type:       song.CoverArt.Type,
			NewAsset:   &song.CoverArt,
		})
	}

	return assets
}

// collectSongAssetsForDeletion collects song assets for deletion
func (processor *AssetsProcessor) collectSongAssetsForDeletion(act *entity.Act, albumID string, song entity.Song) []AssetChange {
	var assets []AssetChange

	assets = append(assets, AssetChange{
		Action:     Delete,
		EntityType: Song,
		ActID:      act.ID,
		UserID:     act.UserID,
		AlbumID:    albumID,
		SongID:     song.ID,
		EntityID:   song.ID,
		Type:       song.File.Type,
		OldAsset:   &song.File,
	})
	// If the song has cover art
	if !isAssetEmpty(song.CoverArt) {
		assets = append(assets, AssetChange{
			Action:     Delete,
			EntityType: Song,
			ActID:      act.ID,
			UserID:     act.UserID,
			AlbumID:    albumID,
			SongID:     song.ID,
			EntityID:   song.ID,
			Type:       song.CoverArt.Type,
			OldAsset:   &song.CoverArt,
		})
	}

	return assets
}

// FIXME: This solution is not scalable. We need to find a better way to compare assets.
// handleAssets processes the collected assets
func (processor *AssetsProcessor) handleAssets(assets []AssetChange) (*AssetsProcessorOutput, error) {
	var assetsURLs []AssetURL
	var addedAssets []AssetChange
	var deletedAssets []AssetChange

	// TODO: Refactor this switch statement with a map of functions
	// TODO: Standardize the input and output of the functions
	// var GetHandler = map[Action]func(AssetChange) ([]AssetURL, []AssetChange, error){}

	for _, asset := range assets {
		var err error

		switch asset.Action {
		case Delete:
			deletedAssets, err = processor.handleDeletedAsset(asset, deletedAssets)
		case Add:
			assetsURLs, addedAssets, err = processor.handleAddedAsset(asset, assetsURLs, addedAssets)
		case Update:
			assetsURLs, addedAssets, err = processor.handleUpdatedAsset(asset, assetsURLs, addedAssets)
		default:
			return nil, fmt.Errorf("unknown action: %s", asset.Action)
		}

		if err != nil {
			return nil, err
		}
	}

	return &AssetsProcessorOutput{
		AddedAssets:   addedAssets,
		DeletedAssets: deletedAssets,
		AssetsURLs:    assetsURLs,
	}, nil
}

// FIXME: This code can be
// handleAddedAsset processes added assets
func (processor *AssetsProcessor) handleAddedAsset(asset AssetChange, assetsURLs []AssetURL, addedAssets []AssetChange) ([]AssetURL, []AssetChange, error) {
	// Generate a presigned URL for uploading the file
	var presignedURL string
	var err error

	// Use the appropriate repository depending on the asset type
	switch asset.Type {
	case entity.Audio:
		presignedURL, err = processor.singAudioURL(asset.NewAsset.URL)
	case entity.ImageAct, entity.ImageAlbum:
		presignedURL, err = processor.singImageURL(asset.NewAsset.URL)
	default:
		fmt.Printf("unknown asset: %+v", asset)
		return assetsURLs, addedAssets, fmt.Errorf("unknown asset type: %v", asset.Type)
	}

	if err != nil {
		return assetsURLs, addedAssets, fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	assetURL := AssetURL{
		URL:  presignedURL,
		Name: asset.NewAsset.Name,
		Type: asset.NewAsset.Type,
	}

	assetsURLs = append(assetsURLs, assetURL)
	addedAssets = append(addedAssets, asset)
	return assetsURLs, addedAssets, nil
}

func (processor *AssetsProcessor) singAudioURL(path string) (string, error) {
	return processor.repos.RawAudio.GeneratePresignedURL(processor.ctx, path, URLExpirationTime)
}

func (processor *AssetsProcessor) singImageURL(path string) (string, error) {
	return processor.repos.Assets.GeneratePresignedURL(processor.ctx, path, URLExpirationTime)
}

// handleUpdatedAsset processes updated assets
func (processor *AssetsProcessor) handleUpdatedAsset(asset AssetChange, assetsURLs []AssetURL, addedAssets []AssetChange) ([]AssetURL, []AssetChange, error) {
	// Treat updated assets as new assets
	return processor.handleAddedAsset(asset, assetsURLs, addedAssets)
}

// handleDeletedAsset processes deleted assets
func (processor *AssetsProcessor) handleDeletedAsset(asset AssetChange, deletedAssets []AssetChange) ([]AssetChange, error) {
	// Add the asset to the deletedAssets array
	deletedAssets = append(deletedAssets, asset)
	return deletedAssets, nil
}

// generateAlbumMap generates a map of albums by ID
func generateAlbumMap(albums []entity.Album) map[string]entity.Album {
	albumMap := make(map[string]entity.Album)
	for _, album := range albums {
		albumMap[album.ID] = album
	}
	return albumMap
}

// generateSongMap generates a map of songs by ID
func generateSongMap(songs []entity.Song) map[string]entity.Song {
	songMap := make(map[string]entity.Song)
	for _, song := range songs {
		songMap[song.ID] = song
	}
	return songMap
}
