package entity

import "time"

// Member is a value object that represent a member of an act
type Member struct {
	Name              string    `bson:"name,omitempty"`
	Biography         string    `bson:"biography,omitempty"`
	BirthDate         time.Time `bson:"birth_date,omitempty"`
	ProfilePictureURL string    `bson:"profile_picture_url,omitempty"`
	StartDate         time.Time `bson:"start_date,omitempty"`
	EndDate           time.Time `bson:"end_date,omitempty"`
}
