package repository

import "go.mongodb.org/mongo-driver/mongo"

// MongoActRepository represents the repository for the act entity.
type MongoActRepository struct {
	client     *mongo.Client
	collection string
}

// NewMongoActRepository creates a new instance of MongoActRepository.
func NewMongoActRepository(client *mongo.Client, collection string) *MongoActRepository {
	return &MongoActRepository{
		client:     client,
		collection: collection,
	}
}

// CreateAct creates a new act in the database.
