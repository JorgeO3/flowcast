package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// Database is the name of the database
	Database = "catalog"

	// ActCollection is the name of the act collection in the database
	ActCollection = "acts"
)

// Act represent an musical act entity
type Act struct {
	ID                primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name              string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Type              string             `json:"type,omitempty" bson:"type,omitempty" validate:"required,oneof='Band' 'Solo Artist' 'Duo'"`
	Biography         string             `json:"biography,omitempty" bson:"biography,omitempty" validate:"required"`
	FormationDate     string             `json:"formationDate,omitempty" bson:"formation_date,omitempty" validate:"required,datetime=2006-01-02"`
	DisbandDate       string             `json:"disbandDate,omitempty" bson:"disband_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	ProfilePictureURL string             `json:"profilePictureUrl,omitempty" bson:"profile_picture_url,omitempty" validate:"required,url"`
	Genres            []Genre            `json:"genres,omitempty" bson:"genres,omitempty" validate:"required,dive"`
	Albums            []Album            `json:"albums,omitempty" bson:"albums,omitempty" validate:"dive"`
	Members           []Member           `json:"members,omitempty" bson:"members,omitempty" validate:"dive"`
}

// ActOption represent the functional options for the act entity
type ActOption func(*Act) error

// WithActID set the ID of the act
func WithActID(id primitive.ObjectID) ActOption {
	return func(a *Act) error {
		a.ID = id
		return nil
	}
}

func WithNewActID() ActOption {
	return func(a *Act) error {
		a.ID = primitive.NewObjectID()
		return nil
	}
}

// WithActName set the name of the act
func WithActName(name string) ActOption {
	return func(a *Act) error {
		a.Name = name
		return nil
	}
}

// WithActType set the type of the act
func WithActType(actType string) ActOption {
	return func(a *Act) error {
		a.Type = actType
		return nil
	}
}

// WithActBiography set the biography of the act
func WithActBiography(biography string) ActOption {
	return func(a *Act) error {
		a.Biography = biography
		return nil
	}
}

// WithActFormationDate set the formation date of the act
func WithActFormationDate(formationDate string) ActOption {
	return func(a *Act) error {
		a.FormationDate = formationDate
		return nil
	}
}

// WithActDisbandDate set the disband date of the act
func WithActDisbandDate(disbandDate string) ActOption {
	return func(a *Act) error {
		a.DisbandDate = disbandDate
		return nil
	}
}

// WithActProfilePictureURL set the profile picture URL of the act
func WithActProfilePictureURL(profilePictureURL string) ActOption {
	return func(a *Act) error {
		a.ProfilePictureURL = profilePictureURL
		return nil
	}
}

// WithActGenres set the genres of the act
func WithActGenres(genres []Genre) ActOption {
	return func(a *Act) error {
		a.Genres = genres
		return nil
	}
}

// WithActAlbums set the albums of the act
func WithActAlbums(albums []Album) ActOption {
	return func(a *Act) error {
		a.Albums = albums
		return nil
	}
}

// WithActMembers set the members of the act
func WithActMembers(members []Member) ActOption {
	return func(a *Act) error {
		a.Members = members
		return nil
	}
}

// NewAct create a new act entity
func NewAct(opts ...ActOption) (*Act, error) {
	act := &Act{}
	for _, opt := range opts {
		if err := opt(act); err != nil {
			return nil, err
		}
	}
	return act, nil
}
