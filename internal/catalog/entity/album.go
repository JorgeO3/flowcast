// Package entity provides the domain model for the catalog service.
package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Album represent an album entity
type Album struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id" validate:"required,alpha"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty" validate:"required,alpha"`
	ReleaseDate time.Time          `json:"releasedate,omitempty" bson:"release_date,omitempty" validate:"required,alpha"`
	Genre       Genre              `json:"genre,omitempty" bson:"genre,omitempty" validate:"required,alpha"`
	CoverArtURL string             `json:"coverarturl,omitempty" bson:"cover_art_url,omitempty" validate:"required,url"`
	TotalTracks int                `json:"totaltracks,omitempty" bson:"total_tracks,omitempty" validate:"required,alpha"`
	Songs       []Song             `json:"songs,omitempty" bson:"songs,omitempty" validate:"required"`
}

// AlbumOption represent the functional options for the album entity
type AlbumOption func(*Album)

// WithAlbumID set the ID of the album
func WithAlbumID(id string) AlbumOption {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	return func(a *Album) {
		a.ID = objectID
	}
}

// WithAlbumTitle set the title of the album
func WithAlbumTitle(title string) AlbumOption {
	return func(a *Album) {
		a.Title = title
	}
}

// WithAlbumReleaseDate set the release date of the album
func WithAlbumReleaseDate(releaseDate time.Time) AlbumOption {
	return func(a *Album) {
		a.ReleaseDate = releaseDate
	}
}

// WithAlbumGenre set the genre of the album
func WithAlbumGenre(genre Genre) AlbumOption {
	return func(a *Album) {
		a.Genre = genre
	}
}

// WithAlbumCoverArtURL set the cover art URL of the album
func WithAlbumCoverArtURL(coverArtURL string) AlbumOption {
	return func(a *Album) {
		a.CoverArtURL = coverArtURL
	}
}

// WithAlbumTotalTracks set the total tracks of the album
func WithAlbumTotalTracks(totalTracks int) AlbumOption {
	return func(a *Album) {
		a.TotalTracks = totalTracks
	}
}

// WithAlbumSongs set the songs of the album
func WithAlbumSongs(songs []Song) AlbumOption {
	return func(a *Album) {
		a.Songs = songs
	}
}

// NewAlbum create a new album entity
func NewAlbum(opts ...AlbumOption) Album {
	album := Album{}
	for _, opt := range opts {
		opt(&album)
	}
	return album
}
