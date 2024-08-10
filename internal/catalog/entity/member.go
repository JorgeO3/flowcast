package entity

import "time"

// ArtistMember is a value object that represent a member of an act
type Member struct {
	Name              string
	Biography         string
	BirthDate         time.Time
	ProfilePictureURL string
	StartDate         time.Time
	EndDate           time.Time
}
