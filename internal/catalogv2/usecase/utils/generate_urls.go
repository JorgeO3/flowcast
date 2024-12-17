package utils

import (
	"fmt"
	"time"

	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/entity"
)

const (
	// RawAudioBucketName represents the name of the bucket for raw audio files
	RawAudioBucketName = "raw-audio"
	// AssetsBucketName represents the name of the bucket for assets
	AssetsBucketName = "assets"

	// URLExpirationTime represents the expiration time for the presigned URL
	URLExpirationTime = time.Minute
)

// AssetURL represents the URL of an asset
type AssetURL struct {
	URL  string           `json:"url"`
	Name string           `json:"name"`
	Type entity.AssetType `json:"type"`
}

func generateSongFileURL(actID, albumID, songID string) string {
	return fmt.Sprintf("/%s/%s/%s/%s.wav", RawAudioBucketName, actID, albumID, songID)
}

func generateActPictureURL(actID string) string {
	return fmt.Sprintf("/%s/%s.jpg", AssetsBucketName, actID)
}

func generateAlbumCoverArtURL(actID, albumID string) string {
	return fmt.Sprintf("/%s/%s/%s.jpg", AssetsBucketName, actID, albumID)
}

func generateSongCoverArtURL(actID, albumID, songID string) string {
	return fmt.Sprintf("/%s/%s/%s/%s.jpg", AssetsBucketName, actID, albumID, songID)
}

// GenerateURLs generates the URLs for the act and its assets
func GenerateURLs(act *entity.Act) {
	setAssetURL(&act.ProfilePicture, generateActPictureURL(act.ID))

	for i := range act.Albums {
		setAssetURL(&act.Albums[i].CoverArt, generateAlbumCoverArtURL(act.ID, act.Albums[i].ID))

		for j := range act.Albums[i].Songs {
			setAssetURL(&act.Albums[i].Songs[j].File, generateSongFileURL(act.ID, act.Albums[i].ID, act.Albums[i].Songs[j].ID))
			setAssetURL(&act.Albums[i].Songs[j].CoverArt, generateSongCoverArtURL(act.ID, act.Albums[i].ID, act.Albums[i].Songs[j].ID))
		}
	}
}

// GenerateURLsFromActs generates the URLs for the acts and their assets
func GenerateURLsFromActs(acts []entity.Act) {
	for i := range acts {
		GenerateURLs(&acts[i])
	}
}

// assetsEqual compara dos assets
func assetsEqual(a1, a2 entity.Asset) bool {
	return a1 == a2
}

// isAssetEmpty verifica si un asset es vacío (todos sus campos en cero)
func isAssetEmpty(a entity.Asset) bool {
	return a == entity.Asset{}
}

// setAssetURL sets the URL of an asset if it is not empty
func setAssetURL(asset *entity.Asset, url string) {
	if !isAssetEmpty(*asset) {
		asset.URL = url
	}
}
