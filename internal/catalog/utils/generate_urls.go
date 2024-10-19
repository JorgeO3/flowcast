// Package utils provides utility functions for the catalog service.
package utils

import (
	"context"
	"time"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/assets"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/rawaudio"
)

// SongURL represents a song's downloadable link.
type SongURL struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

// GenerateImagePresignedURL generates a presigned URL for an image (e.g., album cover).
func GenerateImagePresignedURL(ctx context.Context, bucket, actID string, assetRepo assets.Repository) (string, error) {
	filePath := bucket + actID + "/" + ".jpg"
	return assetRepo.GeneratePresignedURL(ctx, filePath, time.Minute)
}

// GenerateImagePresignedURLsFromActs generates presigned URLs for all images in multiple acts.
func GenerateImagePresignedURLsFromActs(ctx context.Context, acts []*entity.Act, bucket string, assetRepo assets.Repository) ([]string, error) {
	var imageURLs []string

	for _, act := range acts {
		url, err := GenerateImagePresignedURL(ctx, bucket, act.ID.Hex(), assetRepo)
		if err != nil {
			return nil, err
		}

		imageURLs = append(imageURLs, url)
	}

	return imageURLs, nil
}

// GenerateSongPresignedURL generates a presigned URL for a specific song in the catalog.
func GenerateSongPresignedURL(ctx context.Context, bucket, actID, albumID, songID string, audioRepo rawaudio.Repository) (string, error) {
	filePath := bucket + actID + "/" + albumID + "/" + songID + ".wav"
	return audioRepo.GeneratePresignedURL(ctx, filePath, time.Minute)
}

// GenerateSongURLsFromAlbum generates presigned URLs for all songs in an album.
func GenerateSongURLsFromAlbum(ctx context.Context, album *entity.Album, bucket, actID string, audioRepo rawaudio.Repository) ([]SongURL, error) {
	songURLs := make([]SongURL, 0, len(album.Songs))

	for _, song := range album.Songs {
		if err := ValidateSongFile(&song.File); err != nil {
			return nil, err
		}

		url, err := GenerateSongPresignedURL(ctx, bucket, actID, album.ID.Hex(), song.ID.Hex(), audioRepo)
		if err != nil {
			return nil, err
		}

		songURLs = append(songURLs, SongURL{URL: url, Name: song.File.Name})
	}

	return songURLs, nil
}

// GenerateSongURLsFromAct generates presigned URLs for all songs in an act (across albums).
func GenerateSongURLsFromAct(ctx context.Context, bucket string, act *entity.Act, audioRepo rawaudio.Repository) ([]SongURL, error) {
	songURLs := make([]SongURL, 0, act.SongsLength())

	for _, album := range act.Albums {
		for _, song := range album.Songs {
			if err := ValidateSongFile(&song.File); err != nil {
				return nil, errors.NewValidation("invalid file", err)
			}

			url, err := GenerateSongPresignedURL(ctx, bucket, act.ID.Hex(), album.ID.Hex(), song.ID.Hex(), audioRepo)
			if err != nil {
				return nil, errors.NewInternal("error generating presigned URL", err)
			}

			songURLs = append(songURLs, SongURL{URL: url, Name: song.File.Name})
		}
	}

	return songURLs, nil
}

// GenerateSongURLsFromActs generates presigned URLs for all songs across multiple acts.
func GenerateSongURLsFromActs(ctx context.Context, acts []*entity.Act, bucket string, audioRepo rawaudio.Repository) ([]SongURL, error) {
	allSongURLs := make([]SongURL, 0)

	for _, act := range acts {
		songURLs, err := GenerateSongURLsFromAct(ctx, bucket, act, audioRepo)
		if err != nil {
			return nil, err
		}
		allSongURLs = append(allSongURLs, songURLs...)
	}

	return allSongURLs, nil
}
