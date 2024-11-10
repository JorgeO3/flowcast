package entity

// Genre represent a value object
type Genre struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty" bson:"description,omitempty" validate:"required"`
}

// AudioFeatures represent the audio features of a song
type AudioFeatures struct {
	Tempo            int     `json:"tempo" bson:"tempo,omitempty" validate:"required,min=60,max=200"`
	AudioKey         string  `json:"audiokey" bson:"audiokey,omitempty" validate:"required"`
	Mode             string  `json:"mode" bson:"mode,omitempty" validate:"required,oneof=0 1"`
	Loudness         float64 `json:"loudness" bson:"loudness,omitempty" validate:"required,min=-60,max=0"`
	Energy           float64 `json:"energy" bson:"energy,omitempty" validate:"required,min=0,max=1"`
	Danceability     float64 `json:"danceability" bson:"danceability,omitempty" validate:"required,min=0,max=1"`
	Speechiness      float64 `json:"speechiness" bson:"speechiness,omitempty" validate:"required,min=0,max=1"`
	Acousticness     float64 `json:"acousticness" bson:"acousticness,omitempty" validate:"required,min=0,max=1"`
	Instrumentalness float64 `json:"instrumentalness" bson:"instrumentalness,omitempty" validate:"required,min=0,max=1"`
	Liveness         float64 `json:"liveness" bson:"liveness,omitempty" validate:"required,min=0,max=1"`
	Velence          float64 `json:"valence" bson:"valance,omitempty" validate:"required,min=0,max=1"`
}

// Song represent a song entity
type Song struct {
	ID            string        `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Title         string        `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	AudioFeatures AudioFeatures `json:"audioFeatures,omitempty" bson:"audio_features,omitempty" validate:"required"`
	File          Asset         `json:"file,omitempty" bson:"file,omitempty" validate:"required"`
	CoverArt      Asset         `json:"coverArt,omitempty" bson:"cover_art,omitempty"`
	Genre         Genre         `json:"genre,omitempty" bson:"genre,omitempty" validate:"required"`
	ReleaseDate   string        `json:"releaseDate,omitempty" bson:"release_date,omitempty" validate:"required,datetime=2006-01-02"`
	Duration      int           `json:"duration,omitempty" bson:"duration,omitempty" validate:"required,min=1"`
}
