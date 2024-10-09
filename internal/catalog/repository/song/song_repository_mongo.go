package repository

import (
	"context"
	"errors"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SongRepositoryMongo is a repository for the song entity.
type SongRepositoryMongo struct {
	collection *mongo.Collection
}

// CreateSongs implements SongRepository.
func (s *SongRepositoryMongo) CreateSongs(ctx context.Context, songs []*entity.Song, actID, albumID primitive.ObjectID) ([]string, error) {
	filter := bson.M{"_id": actID, "albums._id": albumID}
	update := bson.M{"$push": bson.M{"albums.$.songs": bson.M{"$each": songs}}}

	res, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, mongodb.MapError(err)
	}

	if res.ModifiedCount == 0 {
		return nil, errors.New("no documents were modified")
	}

	ids := make([]string, len(songs))
	for i, song := range songs {
		ids[i] = song.ID.Hex()
	}

	return ids, nil
}

// CreateSong implements SongRepository.
func (s *SongRepositoryMongo) CreateSong(ctx context.Context, song *entity.Song, actID, albumID primitive.ObjectID) (string, error) {
	if song.ID.IsZero() {
		song.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": actID, "albums._id": albumID}
	update := bson.M{"$push": bson.M{"albums.$.songs": song}}

	res, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", mongodb.MapError(err)
	}

	if res.ModifiedCount == 0 {
		return "", errors.New("no documents were modified")
	}

	return song.ID.Hex(), nil
}

// DeleteSong implements SongRepository.
func (s *SongRepositoryMongo) DeleteSong(ctx context.Context, actID, albumID, songID primitive.ObjectID) error {
	filter := bson.M{"_id": actID, "albums._id": albumID}
	update := bson.M{"$pull": bson.M{"albums.$.songs": bson.M{"_id": songID}}}

	res, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return mongodb.MapError(err)
	}

	if res.ModifiedCount == 0 {
		return errors.New("song not found or already deleted")
	}

	return nil
}

// GetSongByID implements SongRepository.
func (s *SongRepositoryMongo) GetSongByID(ctx context.Context, actID, albumID, songID primitive.ObjectID) (*entity.Song, error) {
	pipeline := []bson.M{
		{"$match": bson.M{"_id": actID, "albums._id": albumID, "albums.songs._id": songID}},
		{"$unwind": "$albums"},
		{"$unwind": "$albums.songs"},
		{"$match": bson.M{"albums.songs._id": songID}},
		{"$replaceRoot": bson.M{"newRoot": "$albums.songs"}},
	}

	cursor, err := s.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, mongodb.MapError(err)
	}
	defer cursor.Close(ctx)

	var song entity.Song
	if cursor.Next(ctx) {
		err = cursor.Decode(&song)
		if err != nil {
			return nil, mongodb.MapError(err)
		}
		return &song, nil
	}

	return nil, mongodb.MapError(mongo.ErrNoDocuments)
}

// UpdateSong implements SongRepository.
func (s *SongRepositoryMongo) UpdateSong(ctx context.Context, song *entity.Song, actID, albumID, songID primitive.ObjectID) error {
	filter := bson.M{
		"_id":              actID,
		"albums._id":       albumID,
		"albums.songs._id": songID,
	}
	update := bson.M{"$set": bson.M{"albums.$[a].songs.$[s]": song}}
	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"a._id": albumID},
			bson.M{"s._id": songID},
		},
	})

	res, err := s.collection.UpdateOne(ctx, filter, update, arrayFilters)
	if err != nil {
		return mongodb.MapError(err)
	}

	if res.ModifiedCount == 0 {
		return errors.New("song not found or no changes made")
	}

	return nil
}

// NewSongRepositoryMongo creates a new instance of SongRepositoryMongo.
func NewSongRepositoryMongo(db *mongo.Database, collection string) SongRepository {
	return &SongRepositoryMongo{
		collection: db.Collection(collection),
	}
}
