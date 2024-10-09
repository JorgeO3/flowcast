package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Song represent a song entity
type Song struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Title         string             `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	AudioFeatures AudioFeatures      `json:"audioFeatures,omitempty" bson:"audio_features,omitempty" validate:"required"`
	File          AudioFile          `json:"file,omitempty" bson:"file,omitempty" validate:"required"`
	Genre         Genre              `json:"genre,omitempty" bson:"genre,omitempty" validate:"required"`
	ReleaseDate   string             `json:"releaseDate,omitempty" bson:"release_date,omitempty" validate:"required,datetime=2006-01-02"`
	Duration      int                `json:"duration,omitempty" bson:"duration,omitempty" validate:"required,min=1"`
	Lyrics        string             `json:"lyrics,omitempty" bson:"lyrics,omitempty" validate:"required"`
	Explicit      bool               `json:"explicit,omitempty" bson:"explicit,omitempty"`
	Bitrates      []AudioBitrate     `json:"bitrates,omitempty" bson:"bitrates,omitempty" validate:"required,dive"`
}

// SongOption represent the functional options for the song entity
type SongOption func(*Song)

// WithSongID set the ID of the song
func WithSongID(id primitive.ObjectID) SongOption {
	return func(s *Song) {
		s.ID = GetObjectID(id)
	}
}

// WithSongTitle set the title of the song
func WithSongTitle(title string) SongOption {
	return func(s *Song) {
		s.Title = title
	}
}

// WithSongAudioFeatures set the audio features of the song
func WithSongAudioFeatures(audioFeatures AudioFeatures) SongOption {
	return func(s *Song) {
		s.AudioFeatures = audioFeatures
	}
}

// WithSongFile set the file of the song
func WithSongFile(file AudioFile) SongOption {
	return func(s *Song) {
		s.File = file
	}
}

// WithSongGenre set the genre of the song
func WithSongGenre(genre Genre) SongOption {
	return func(s *Song) {
		s.Genre = genre
	}
}

// WithSongReleaseDate set the release date of the song
func WithSongReleaseDate(releaseDate string) SongOption {
	return func(s *Song) {
		s.ReleaseDate = releaseDate
	}
}

// WithSongDuration set the duration of the song
func WithSongDuration(duration int) SongOption {
	return func(s *Song) {
		s.Duration = duration
	}
}

// WithSongLyrics set the lyrics of the song
func WithSongLyrics(lyrics string) SongOption {
	return func(s *Song) {
		s.Lyrics = lyrics
	}
}

// WithSongExplicit set the explicit of the song
func WithSongExplicit(explicit bool) SongOption {
	return func(s *Song) {
		s.Explicit = explicit
	}
}

// WithSongBitrates set the bitrates of the song
func WithSongBitrates(bitrates []AudioBitrate) SongOption {
	return func(s *Song) {
		s.Bitrates = bitrates
	}
}

// NewSong create a new song entity
func NewSong(opts ...SongOption) *Song {
	song := &Song{}
	for _, opt := range opts {
		opt(song)
	}
	return song
}
