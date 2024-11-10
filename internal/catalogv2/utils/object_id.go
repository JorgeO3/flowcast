package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

// ObjectIDFrom get the object ID
func ObjectIDFrom(id any) primitive.ObjectID {
	if !id.(primitive.ObjectID).IsZero() {
		return id.(primitive.ObjectID)
	}

	if id.(string) != "" {
		if oid, err := primitive.ObjectIDFromHex(id.(string)); err == nil {
			return oid
		}
	}
	return primitive.NewObjectID()
}
