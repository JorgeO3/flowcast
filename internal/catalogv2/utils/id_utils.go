package utils

import (
	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func genID() string {
	return primitive.NewObjectID().Hex()
}

// GenerateIDs generates the IDs for the act and its assets
func GenerateIDs(act *entity.Act) {
	if !IsActEmpty(act) && act.ID == "" {
		act.ID = genID()
	}

	for i := range act.Albums {
		if !IsAlbumEmpty(&act.Albums[i]) && act.Albums[i].ID == "" {
			act.Albums[i].ID = genID()
		}

		for j := range act.Albums[i].Songs {
			if !IsSongEmpty(&act.Albums[i].Songs[j]) && act.Albums[i].Songs[j].ID == "" {
				act.Albums[i].Songs[j].ID = genID()
			}
		}
	}
}

// GenerateIDsFromActs generates the IDs for the acts and their assets
func GenerateIDsFromActs(acts []entity.Act) {
	for i := range acts {
		GenerateIDs(&acts[i])
	}
}

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

// GetUserIDs returns the user IDs from the acts
func GetUserIDs(acts []entity.Act) []string {
	var userIDs []string
	for i := range acts {
		userIDs = append(userIDs, acts[i].UserID)
	}
	return userIDs
}
