package entity

import "time"

// Member is a value object that represent a member of an act
type Member struct {
	Name              string    `json:"name,omitempty" bson:"name,omitempty" validate:"required,alpha"`
	Biography         string    `json:"biography,omitempty" bson:"biography,omitempty" validate:"required,alpha"`
	BirthDate         time.Time `json:"birthdate,omitempty" bson:"birth_date,omitempty" validate:"required,alpha"`
	ProfilePictureURL string    `json:"profilePictureUrl,omitempty" bson:"profile_picture_url,omitempty" validate:"required,url"`
	StartDate         time.Time `json:"startDate,omitempty" bson:"start_date,omitempty" validate:"required,alpha"`
	EndDate           time.Time `json:"endDate,omitempty" bson:"end_date,omitempty" validate:"required,alpha"`
}
