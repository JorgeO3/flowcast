package repository

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
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

// NewMongoActRepository creates a new instance of MongoActRepository.
func NewMongoActRepository(db *mongo.Database, collection string) ActRepository {
	return &MongoActRepository{
		db:         db,
		collection: db.Collection(collection),
	}
}

// CreateAct implements ActRepository.
func (m *MongoActRepository) CreateAct(ctx context.Context, act *entity.Act) (string, error) {
	res, err := m.collection.InsertOne(ctx, act)
	if err != nil {
		return "", mongodb.MapError(err)
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}

// UpdateAct implements ActRepository.
func (m *MongoActRepository) UpdateAct(ctx context.Context, act *entity.Act) error {
	filter := bson.M{"_id": act.ID}
	update := bson.M{"$set": act}

	_, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return mongodb.MapError(err)
	}
	return nil
}

// GetActByID implements ActRepository.
func (m *MongoActRepository) GetActByID(ctx context.Context, id primitive.ObjectID) (*entity.Act, error) {
	var act entity.Act
	filter := bson.M{"_id": id}
	err := m.collection.FindOne(ctx, filter).Decode(&act)
	if err != nil {
		return nil, mongodb.MapError(err)
	}
	return &act, nil
}

// CreateManyActs implements ActRepository.
func (m *MongoActRepository) CreateManyActs(context.Context, []*entity.Act) error {
	panic("unimplemented")
}

// DeleteAct implements ActRepository.
func (m *MongoActRepository) DeleteAct(context.Context, primitive.ObjectID) error {
	panic("unimplemented")
}

// DeleteManyActs implements ActRepository.
func (m *MongoActRepository) DeleteManyActs(context.Context, primitive.M) error {
	panic("unimplemented")
}

// GetManyActs implements ActRepository.
func (m *MongoActRepository) GetManyActs(context.Context, primitive.M, *options.FindOptions) ([]*entity.Act, error) {
	panic("unimplemented")
}
