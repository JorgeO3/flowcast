package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/catalog/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// type mongoAct struct {
// 	id            primitive.ObjectID `bson:"_id"`
// 	name          string             `bson:"name,omitempty"`
// 	actType       string             `bson:"type,omitempty"`
// 	biography     string             `bson:"biography,omitempty"`
// 	formationDate primitive.DateTime `bson:"formation_date,omitempty"`
// 	disbandDate   primitive.DateTime `bson:"disband_date,omitempty"`
// 	profilePicURL string             `bson:"profile_picture_url,omitempty"`
// 	genres        []mongoGenre       `bson:"genres,omitempty"`
// 	members       []mongoMember      `bson:"members,omitempty"`
// 	albums        []mongoAlbum       `bson:"albums,omitempty"`
// }

// func fromAct(act *entity.Act) *mongoAct {
// 	id, err := primitive.ObjectIDFromHex(act.ID)

// 	return &mongoAct{
// 		id: primitive.ObjectIDFromHex(act.ID),
// 	}
// }

// type mongoMember struct {
// 	name          string             `bson:"name,omitempty"`
// 	biography     string             `bson:"biography,omitempty"`
// 	birthDate     primitive.DateTime `bson:"birth_date,omitempty"`
// 	profilePicURL string             `bson:"profile_picture_url,omitempty"`
// 	startDate     primitive.DateTime `bson:"start_date,omitempty"`
// 	endDate       primitive.DateTime `bson:"end_date,omitempty"`
// }

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
func (m *MongoActRepository) CreateAct(ctx context.Context, act entity.Act) error {
	collection := m.db.Collection(m.collection)
	_, err := collection.InsertOne(ctx, act)
	return err
}

// GetActByID implements ActRepository.
func (m *MongoActRepository) GetActByID(ctx context.Context, id string) (*entity.Act, error) {
	var act entity.Act

	collection := m.db.Collection(m.collection)
	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(&act)
	return &act, err
}

// UpdateAct implements ActRepository.
func (m *MongoActRepository) UpdateAct(ctx context.Context, id string, updates map[string]interface{}) error {
	panic("unimplemented")
}

// DeleteAct implements ActRepository.
func (m *MongoActRepository) DeleteAct(ctx context.Context, id string) error {
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
