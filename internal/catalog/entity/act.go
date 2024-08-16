package entity

import "time"

const collectionAct = "acts"

// Act represent an musical act entity
type Act struct {
	ID                int
	Name              string
	Type              string
	Biography         string
	FormationDate     time.Time
	DisbandDate       time.Time
	ProfilePictureURL string
	Genres            []Genre
	Albums            []Album
	Members           []Member
}

// ActOption represent the functional options for the act entity
type ActOption func(*Act)

// WithActID set the ID of the act
func WithActID(id int) ActOption {
	return func(a *Act) {
		a.ID = id
	}
}

// WithActName set the name of the act
func WithActName(name string) ActOption {
	return func(a *Act) {
		a.Name = name
	}
}

// WithActType set the type of the act
func WithActType(actType string) ActOption {
	return func(a *Act) {
		a.Type = actType
	}
}

// WithActBiography set the biography of the act
func WithActBiography(biography string) ActOption {
	return func(a *Act) {
		a.Biography = biography
	}
}

// WithActFormationDate set the formation date of the act
func WithActFormationDate(formationDate time.Time) ActOption {
	return func(a *Act) {
		a.FormationDate = formationDate
	}
}

// WithActDisbandDate set the disband date of the act
func WithActDisbandDate(disbandDate time.Time) ActOption {
	return func(a *Act) {
		a.DisbandDate = disbandDate
	}
}

// WithActProfilePictureURL set the profile picture URL of the act
func WithActProfilePictureURL(profilePictureURL string) ActOption {
	return func(a *Act) {
		a.ProfilePictureURL = profilePictureURL
	}
}

// WithActGenres set the genres of the act
func WithActGenres(genres []Genre) ActOption {
	return func(a *Act) {
		a.Genres = genres
	}
}

// WithActAlbums set the albums of the act
func WithActAlbums(albums []Album) ActOption {
	return func(a *Act) {
		a.Albums = albums
	}
}

// WithActMembers set the members of the act
func WithActMembers(members []Member) ActOption {
	return func(a *Act) {
		a.Members = members
	}
}

// NewAct create a new act entity
func NewAct(opts ...ActOption) *Act {
	act := &Act{}
	for _, opt := range opts {
		opt(act)
	}
	return act
}
