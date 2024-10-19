package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// Database is the name of the database
	Database = "catalog"

	// ActCollection is the name of the act collection in the database
	ActCollection = "acts"

	// CreateActTopic is the Kafka topic for act creation events.
	CreateActTopic = "catalog.acts.created"

	// UpdateActTopic is the Kafka topic for act update events.
	UpdateActTopic = "catalog.acts.updated"

	// DeleteActTopic is the Kafka topic for act deletion events.
	DeleteActTopic = "catalog.acts.deleted"

	// CreateActsTopic is the Kafka topic for bulk act creation events.
	CreateActsTopic = "catalog.acts.bulk_created"
)

// Image represent an image object value
type Image struct {
	Name string `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Ext  string `json:"ext,omitempty" bson:"ext,omitempty" validate:"required"`
	Size int64  `json:"size,omitempty" bson:"size,omitempty" validate:"required"`
}

// Act represent an musical act entity
type Act struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	UserID         string             `json:"userId,omitempty" bson:"user_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	ProfilePicture Image              `json:"profilePictureUrl,omitempty" bson:"profile_picture_url,omitempty" validate:"required"`
	Genres         []Genre            `json:"genres,omitempty" bson:"genres,omitempty" validate:"required,dive"`
	Albums         []Album            `json:"albums,omitempty" bson:"albums,omitempty" validate:"dive"`
}

// ActOption represent the functional options for the act entity
type ActOption func(*Act)

// WithActID set the ID of the act
func WithActID(id any) ActOption {
	return func(a *Act) {
		a.ID = GetObjectID(id)
	}
}

// GetObjectID get the object ID
func GetObjectID(id any) primitive.ObjectID {
	if !id.(primitive.ObjectID).IsZero() {
		return id.(primitive.ObjectID)
	}

	if id.(string) != "" {
		if oid, err := primitive.ObjectIDFromHex(id.(string)); err == nil {
			return oid
		}
	}
	return primitive.NewObjectID()
}

// WithActUserID set the user ID of the act
func WithActUserID(userID string) ActOption {
	return func(a *Act) {
		a.UserID = userID
	}
}

// WithActName set the name of the act
func WithActName(name string) ActOption {
	return func(a *Act) {
		a.Name = name
	}
}

// WithActProfilePictureURL set the profile picture URL of the act
func WithActProfilePictureURL(profilePicture Image) ActOption {
	return func(a *Act) {
		a.ProfilePicture = profilePicture
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

// NewAct create a new act entity
func NewAct(opts ...ActOption) *Act {
	act := &Act{}
	for _, opt := range opts {
		opt(act)
	}
	return act
}

// SongsLength return the total number of songs in the act
func (a *Act) SongsLength() int {
	length := 0
	for _, album := range a.Albums {
		length += len(album.Songs)
	}
	return length
}

// GetSongs return all the songs in the act
func (a *Act) GetSongs() []Song {
	songs := make([]Song, a.SongsLength())

	for _, album := range a.Albums {
		for _, song := range album.Songs {
			songs = append(songs, song)
		}
	}

	return songs
}

// GenerateIDs generates new ObjectIDs for the act and its albums and songs.
//
// NOTE: This method is only used for creating new acts
// and should not be used for any other use case
func (a *Act) GenerateIDs() {
	a.ID = primitive.NewObjectID()
	for i := range a.Albums {
		a.Albums[i].ID = primitive.NewObjectID()
		for j := range a.Albums[i].Songs {
			a.Albums[i].Songs[j].ID = primitive.NewObjectID()
		}
	}
}
