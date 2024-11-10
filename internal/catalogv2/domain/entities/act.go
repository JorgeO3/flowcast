package entity

// TODO: Move the following code to the correct file
const (
	// Database is the name of the database
	Database = "catalog"

	// ActCollection is the name of the act collection in the database
	ActCollection = "acts"

	// CreateActTopic is the Kafka topic for act creation events.
	CreateActTopic = "catalog.acts.created"

	// UpdateActTopic is the Kafka topic for act update events.
	UpdateActTopic = "catalog.acts.updated"

	// DeleteActTopic is the Kafka topic for act deletion events.
	DeleteActTopic = "catalog.acts.deleted"

	// CreateActsTopic is the Kafka topic for bulk act creation events.
	CreateActsTopic = "catalog.acts.bulk_created"
)

// Act represent an musical act entity
type Act struct {
	ID             string  `json:"id,omitempty" bson:"_id,omitempty"`
	UserID         string  `json:"userId,omitempty" bson:"user_id,omitempty"`
	Name           string  `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	ProfilePicture Asset   `json:"profilePictureUrl,omitempty" bson:"profile_picture_url,omitempty" validate:"required"`
	Genres         []Genre `json:"genres,omitempty" bson:"genres,omitempty" validate:"required,dive"`
	Albums         []Album `json:"albums,omitempty" bson:"albums,omitempty" validate:"dive"`
}
