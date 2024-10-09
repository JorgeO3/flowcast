package album

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repository provides an interface to interact with the album repository.
type Repository interface {
	CreateAlbum(ctx context.Context, actID primitive.ObjectID, album *entity.Album) (string, error)
	CreateAlbums(ctx context.Context, actID primitive.ObjectID, albums *[]entity.Album) ([]string, error)
	GetAlbumByID(ctx context.Context, actID, albumID primitive.ObjectID) (*entity.Album, error)
	GetAlbums(ctx context.Context, actID primitive.ObjectID) (*[]entity.Album, error)
	UpdateAlbum(ctx context.Context, actID primitive.ObjectID, album *entity.Album) error
	DeleteAlbum(ctx context.Context, actID, albumID primitive.ObjectID) error
}
