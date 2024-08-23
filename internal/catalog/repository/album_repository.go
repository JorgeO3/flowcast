package repository

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AlbumRepository provides an interface to interact with the album repository.
type AlbumRepository interface {
	CreateAlbum(context.Context, primitive.ObjectID, *entity.Album) error
	CreateManyAlbums(context.Context, primitive.ObjectID, []*entity.Album) error
	GetAlbumByID(context.Context, primitive.ObjectID) (*entity.Album, error)
	GetManyAlbums(context.Context, bson.M) ([]*entity.Album, error)
	UpdateAlbum(context.Context, *entity.Album) error
	UpdateManyAlbums(context.Context, bson.M, bson.M) error
	DeleteAlbum(context.Context, primitive.ObjectID) error
	DeleteManyAlbums(context.Context, bson.M) error
}
