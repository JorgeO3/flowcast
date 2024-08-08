package entity

import "time"

// Artist represent an artist entity
type Artist struct {
	ID                int
	Name              string
	Biography         string
	BirthDate         time.Time
	Genre             Genre
	ProfilePictureURL string
	Albums            []Album
}

// ArtistOption represent the functional options for the artist entity
type ArtistOption func(*Artist)

// WithArtistID set the ID of the artist
func WithArtistID(id int) ArtistOption {
	return func(a *Artist) {
		a.ID = id
	}
}

// WithArtistName set the name of the artist
func WithArtistName(name string) ArtistOption {
	return func(a *Artist) {
		a.Name = name
	}
}

// WithArtistBiography set the biography of the artist
func WithArtistBiography(biography string) ArtistOption {
	return func(a *Artist) {
		a.Biography = biography
	}
}

// WithArtistBirthDate set the birth date of the artist
func WithArtistBirthDate(birthDate time.Time) ArtistOption {
	return func(a *Artist) {
		a.BirthDate = birthDate
	}
}

// WithArtistGenre set the genre of the artist
func WithArtistGenre(genre Genre) ArtistOption {
	return func(a *Artist) {
		a.Genre = genre
	}
}

// WithArtistProfilePictureURL set the profile picture URL of the artist
func WithArtistProfilePictureURL(profilePictureURL string) ArtistOption {
	return func(a *Artist) {
		a.ProfilePictureURL = profilePictureURL
	}
}

// WithArtistAlbums set the albums of the artist
func WithArtistAlbums(albums []Album) ArtistOption {
	return func(a *Artist) {
		a.Albums = albums
	}
}

// NewArtist create a new artist entity
func NewArtist(opts ...ArtistOption) *Artist {
	artist := &Artist{}
	for _, opt := range opts {
		opt(artist)
	}
	return artist
}
