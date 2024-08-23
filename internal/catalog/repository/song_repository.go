package repository

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SongRepository represent the song repository contract
type SongRepository interface {
	CreateSong(context.Context, primitive.ObjectID, *entity.Song) error
	CreateManySongs(context.Context, primitive.ObjectID, []*entity.Song) error
	GetSongByID(context.Context, primitive.ObjectID) (*entity.Song, error)
	GetManySongs(context.Context, bson.M) ([]*entity.Song, error)
	UpdateSong(context.Context, *entity.Song) error
	UpdateManySongs(context.Context, bson.M, bson.M) error
	DeleteSong(context.Context, primitive.ObjectID) error
	DeleteManySongs(context.Context, bson.M) error
}
