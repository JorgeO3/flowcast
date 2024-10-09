package repository

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SongRepository represent the song repository contract
type SongRepository interface {
	CreateSong(ctx context.Context, song *entity.Song, actID, albumID primitive.ObjectID) (string, error)
	CreateSongs(ctx context.Context, songs []*entity.Song, actID, albumID primitive.ObjectID) ([]string, error)
	GetSongByID(ctx context.Context, actID, albumID, SongID primitive.ObjectID) (*entity.Song, error)
	UpdateSong(ctx context.Context, song *entity.Song, actID, albumID, SongID primitive.ObjectID) error
	DeleteSong(ctx context.Context, actID, albumID, SongID primitive.ObjectID) error
}
