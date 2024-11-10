// Package utils provides the utility functions for the catalog service
package utils

import "github.com/JorgeO3/flowcast/internal/catalog/entity"

// IsActEmpty checks if the act is empty
func IsActEmpty(act *entity.Act) bool {
	return act == nil ||
		act.ID == "" ||
		act.Name == "" ||
		act.Albums == nil ||
		act.Genres == nil ||
		act.ProfilePicture == entity.Asset{} ||
		act.UserID == ""
}

// IsAlbumEmpty checks if the album is empty
func IsAlbumEmpty(album *entity.Album) bool {
	return album == nil ||
		album.ID == "" ||
		album.Title == "" ||
		album.ReleaseDate == "" ||
		album.Genre == entity.Genre{} ||
		album.CoverArt == entity.Asset{} ||
		album.TotalTracks == 0 ||
		album.Songs == nil
}

// IsSongEmpty checks if the song is empty
func IsSongEmpty(song *entity.Song) bool {
	return song == nil ||
		song.ID == "" ||
		song.Title == "" ||
		song.AudioFeatures == entity.AudioFeatures{} ||
		song.File == entity.Asset{} ||
		song.Genre == entity.Genre{} ||
		song.ReleaseDate == "" ||
		song.Duration == 0
}
