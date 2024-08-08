package entity

import "time"

// ArtistMember represent an artist member entity
type ArtistMember struct {
	Name              string
	Biography         string
	BirthDate         time.Time
	ProfilePictureURL string
	StartDate         time.Time
	EndDate           time.Time
}

// Act represent an musical act entity
type Act struct {
	ID                int
	Name              string
	Type              string
	Biography         string
	FormationDate     time.Time
	DisbandDate       time.Time
	ProfilePictureURL string
	Genre             Genre
	Albums            []Album
	members           []ArtistMember
}

// ActOption represent the functional options for the act entity
type ActOption func(*Act)

// WithActID set the ID of the act
