package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/catalog/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoActRepository is a repository for the act entity.
type MongoActRepository struct {
	db         *mongo.Database
	collection string
}

// NewMongoActRepository creates a new instance of MongoActRepository.
func NewMongoActRepository(db *mongo.Database, collection string) ActRepository {
	return &MongoActRepository{
		db:         db,
		collection: collection,
	}
}

// CreateAct implements ActRepository.
func (m *MongoActRepository) Create(ctx context.Context, act *entity.Act) (string, error) {
	collection := m.db.Collection(m.collection)
	_, err := collection.InsertOne(ctx, act)
	return "", err
}

// GetActByID implements ActRepository.
func (m *MongoActRepository) GetByID(ctx context.Context, id string) (*entity.Act, error) {
	var act entity.Act

	collection := m.db.Collection(m.collection)
	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(&act)
	return &act, err
}

// UpdateAct implements ActRepository.
func (m *MongoActRepository) Update(ctx context.Context, id string, updates map[string]interface{}) error {
	panic("unimplemented")
}

// DeleteAct implements ActRepository.
func (m *MongoActRepository) Delete(ctx context.Context, id string) error {
	collection := m.db.Collection(m.collection)
	filter := bson.M{"id": id}

	_, err := collection.DeleteOne(ctx, filter)
	return err
}

// AddAlbum implements ActRepository.
func (m *MongoActRepository) AddAlbum(ctx context.Context, actID string, album entity.Album) error {
	collection := m.db.Collection(m.collection)
	filter := bson.M{"id": actID}
	update := bson.M{"$push": bson.M{"albums": album}}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

// AddMember implements ActRepository.
func (m *MongoActRepository) AddMember(ctx context.Context, actID string, member entity.Member) error {
	collection := m.db.Collection(m.collection)
	filter := bson.M{"id": actID}
	update := bson.M{"$push": bson.M{"members": member}}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

// RemoveAlbum implements ActRepository.
func (m *MongoActRepository) RemoveAlbum(ctx context.Context, actID string, albumID string) error {
	collection := m.db.Collection(m.collection)
	filter := bson.M{"id": actID}
	update := bson.M{"$pull": bson.M{"albums": bson.M{"id": albumID}}}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

// RemoveMember implements ActRepository.
func (m *MongoActRepository) RemoveMember(ctx context.Context, actID, memberName string) error {
	collection := m.db.Collection(m.collection)
	filter := bson.M{"id": actID}
	update := bson.M{"$pull": bson.M{"members": bson.M{"Name": memberName}}}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
