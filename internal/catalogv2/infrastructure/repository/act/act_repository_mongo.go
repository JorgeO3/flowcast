// Package act implements the act repository contract
package act

import (
	"context"
	"fmt"

	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/entity"
	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/repository"
	"github.com/JorgeO3/flowcast/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoActRepository is a repository for the act entity.
type MongoActRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

// NewRepository creates a new instance of MongoActRepository.
func NewRepository(db *mongo.Database, collection string) repository.ActRepository {
	return &MongoActRepository{
		db:         db,
		collection: db.Collection(collection),
	}
}

// CreateOne inserts a new act in the database.
func (m *MongoActRepository) CreateOne(ctx context.Context, act *entity.Act) (string, error) {
	res, err := m.collection.InsertOne(ctx, act)
	if err != nil {
		return "", mongodb.MapError(err)
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}

// UpdateOne updates an act in the database.
func (m *MongoActRepository) UpdateOne(ctx context.Context, act *entity.Act) error {
	filter := bson.M{"_id": act.ID}
	update := bson.M{"$set": act}

	_, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return mongodb.MapError(err)
	}
	return nil
}

// ReadOne retrieves an act from the database.
func (m *MongoActRepository) ReadOne(ctx context.Context, id string) (*entity.Act, error) {
	var act entity.Act
	filter := bson.M{"_id": id}

	err := m.collection.FindOne(ctx, filter).Decode(&act)
	if err != nil {
		return nil, mongodb.MapError(err)
	}

	return &act, nil
}

// CreateMany inserts multiple acts in the database.
func (m *MongoActRepository) CreateMany(ctx context.Context, acts []entity.Act) ([]string, error) {
	var docs []interface{}
	for i := range acts {
		docs = append(docs, &acts[i])
	}

	res, err := m.collection.InsertMany(ctx, docs)
	if err != nil {
		return nil, mongodb.MapError(err)
	}

	ids := make([]string, 0, len(res.InsertedIDs))
	for _, id := range res.InsertedIDs {
		ids = append(ids, id.(primitive.ObjectID).String())
	}

	return ids, nil
}

// DeleteOne removes an act from the database.
func (m *MongoActRepository) DeleteOne(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}

	if _, err := m.collection.DeleteOne(ctx, filter); err != nil {
		fmt.Println(err)
		return mongodb.MapError(err)
	}

	return nil
}

// ReadMany retrieves multiple acts from the database.
func (m *MongoActRepository) ReadMany(ctx context.Context, genre string, limit, offset int64) ([]*entity.Act, error) {
	filter := bson.M{}
	if genre != "" {
		filter["genre"] = genre
	}
	opts := &options.FindOptions{
		Limit: &limit,
		Skip:  &offset,
	}

	cursor, err := m.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, mongodb.MapError(err)
	}
	defer cursor.Close(ctx)

	acts := make([]*entity.Act, 0)
	if err := cursor.All(ctx, &acts); err != nil {
		return nil, mongodb.MapError(err)
	}

	return acts, nil
}
