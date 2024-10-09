package album

import (
	"context"
	"errors"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoAlbumRepository is a repository for the album entity.
type MongoAlbumRepository struct {
	collection *mongo.Collection
}

// NewRepository creates a new instance of MongoAlbumRepository.
func NewRepository(db *mongo.Database, collection string) Repository {
	return &MongoAlbumRepository{
		collection: db.Collection(collection),
	}
}

// CreateAlbum implements AlbumRepository.
func (m *MongoAlbumRepository) CreateAlbum(ctx context.Context, actID primitive.ObjectID, album *entity.Album) (string, error) {
	if album.ID.IsZero() {
		album.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": actID}
	update := bson.M{
		"$push": bson.M{
			"albums": album,
		},
	}

	res, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", mongodb.MapError(err)
	}

	if res.ModifiedCount == 0 {
		return "", errors.New("act not found or album not added")
	}

	return album.ID.Hex(), nil
}

// CreateAlbums implements AlbumRepository.
func (m *MongoAlbumRepository) CreateAlbums(ctx context.Context, actID primitive.ObjectID, albums *[]entity.Album) ([]string, error) {
	for i := range *albums {
		if (*albums)[i].ID.IsZero() {
			(*albums)[i].ID = primitive.NewObjectID()
		}
	}

	filter := bson.M{"_id": actID}
	update := bson.M{
		"$push": bson.M{
			"albums": bson.M{
				"$each": albums,
			},
		},
	}

	res, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, mongodb.MapError(err)
	}

	// TODO: Check if the album is not found or no changes made
	if res.ModifiedCount == 0 {
		return nil, errors.New("act not found or albums not added")
	}

	ids := make([]string, len(*albums))
	for i, album := range *albums {
		ids[i] = album.ID.Hex()
	}

	return ids, nil
}

// DeleteAlbum implements AlbumRepository.
func (m *MongoAlbumRepository) DeleteAlbum(ctx context.Context, actID, albumID primitive.ObjectID) error {
	filter := bson.M{"_id": actID}
	update := bson.M{"$pull": bson.M{"albums": bson.M{"_id": albumID}}}

	res, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return mongodb.MapError(err)
	}

	// TODO: Check if the album is not found or no changes made
	if res.ModifiedCount == 0 {
		return errors.New("act or album not found")
	}

	return nil
}

// GetAlbumByID implements AlbumRepository.
func (m *MongoAlbumRepository) GetAlbumByID(ctx context.Context, actID, albumID primitive.ObjectID) (*entity.Album, error) {
	pipeline := []bson.M{
		{"$match": bson.M{"_id": actID}},
		{"$unwind": "$albums"},
		{"$match": bson.M{"albums._id": albumID}},
		{"$replaceRoot": bson.M{"newRoot": "$albums"}},
	}

	cursor, err := m.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, mongodb.MapError(err)
	}
	defer cursor.Close(ctx)

	var album entity.Album
	if cursor.Next(ctx) {
		err = cursor.Decode(&album)
		if err != nil {
			return nil, mongodb.MapError(err)
		}
		return &album, nil
	}

	return nil, errors.New("album not found")
}

// GetAlbums implements AlbumRepository.
func (m *MongoAlbumRepository) GetAlbums(ctx context.Context, actID primitive.ObjectID) (*[]entity.Album, error) {
	pipeline := []bson.M{
		{"$match": bson.M{"_id": actID}},
		{"$project": bson.M{"albums": 1, "_id": 0}},
	}

	cursor, err := m.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, mongodb.MapError(err)
	}
	defer cursor.Close(ctx)

	var result struct {
		Albums []entity.Album `bson:"albums"`
	}

	// TODO: Check if the album is not found or no changes made
	if cursor.Next(ctx) {
		err = cursor.Decode(&result)
		if err != nil {
			return nil, mongodb.MapError(err)
		}
		return &result.Albums, nil
	}

	return nil, errors.New("act not found or no albums")
}

// UpdateAlbum implements AlbumRepository.
func (m *MongoAlbumRepository) UpdateAlbum(ctx context.Context, actID primitive.ObjectID, album *entity.Album) error {
	filter := bson.M{"_id": actID, "albums._id": album.ID}
	update := bson.M{"$set": bson.M{"albums.$": album}}

	res, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return mongodb.MapError(err)
	}

	// TODO: Check if the album is not found or no changes made
	if res.ModifiedCount == 0 {
		return errors.New("act or album not found, or no changes made")
	}

	return nil
}
