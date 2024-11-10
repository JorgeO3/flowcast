// Package entity provides the domain model for the catalog service.
package entity

// Album represent an album entity
type Album struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Title       string `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	ReleaseDate string `json:"releaseDate,omitempty" bson:"release_date,omitempty" validate:"required,datetime=2006-01-02"`
	Genre       Genre  `json:"genre,omitempty" bson:"genre,omitempty" validate:"required"`
	CoverArt    Asset  `json:"coverarturl,omitempty" bson:"cover_art_url,omitempty" validate:"required"`
	TotalTracks int    `json:"totaltracks,omitempty" bson:"total_tracks,omitempty" validate:"required,min=1"`
	Songs       []Song `json:"songs,omitempty" bson:"songs,omitempty" validate:"dive"`
}
