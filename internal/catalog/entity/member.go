package entity

// Member is a value object that represents a member of an act
type Member struct {
	Name              string `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Biography         string `json:"biography,omitempty" bson:"biography,omitempty" validate:"required"`
	BirthDate         string `json:"birthdate,omitempty" bson:"birth_date,omitempty" validate:"required,datetime=2006-01-02"`
	ProfilePictureURL string `json:"profilePictureUrl,omitempty" bson:"profile_picture_url,omitempty" validate:"required,url"`
	StartDate         string `json:"startDate,omitempty" bson:"start_date,omitempty" validate:"required,datetime=2006-01-02"`
	EndDate           string `json:"endDate,omitempty" bson:"end_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
}
