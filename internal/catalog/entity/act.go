package entity

import "time"

const (
	// Database is the name of the database
	Database = "catalog"

	// ActCollection is the name of the act collection in the database
	ActCollection = "acts"
)

// Act represent an musical act entity
type Act struct {
	ID                string    `json:"id,omitempty" bson:"_id" validate:"alpha"`
	Name              string    `json:"name,omitempty" bson:"name,omitempty" validate:"required,alpha"`
	Type              string    `json:"type,omitempty" bson:"type,omitempty" validate:"required,alpha"`
	Biography         string    `json:"biography,omitempty" bson:"biography,omitempty" validate:"required,alpha"`
	FormationDate     time.Time `json:"formationDate,omitempty" bson:"formation_date,omitempty" validate:"required,alphanum"`
	DisbandDate       time.Time `json:"disbandDate,omitempty" bson:"disband_date,omitempty" validate:"required,alphanum"`
	ProfilePictureURL string    `json:"profilePictureUrl,omitempty" bson:"profile_picture_url,omitempty" validate:"required,url"`
	Genres            []Genre   `json:"genres,omitempty" bson:"genres,omitempty" validate:"required"`
	Albums            []Album   `json:"albums,omitempty" bson:"albums,omitempty" validate:"required"`
	Members           []Member  `json:"members,omitempty" bson:"members,omitempty" validate:"required"`
}

// ActOption represent the functional options for the act entity
type ActOption func(*Act)

// WithActID set the ID of the act
func WithActID(id string) ActOption {
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