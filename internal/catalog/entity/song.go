package entity

import "time"

// Song represent a song entity
type Song struct {
	ID            string         `json:"id,omitempty" bson:"_id" validate:""`
	Title         string         `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	AudioFeatures AudioFeatures  `json:"audioFeatures,omitempty" bson:"audio_features,omitempty" validate:"required"`
	Genre         Genre          `json:"genre,omitempty" bson:"genre,omitempty" validate:"required"`
	ReleaseDate   time.Time      `json:"releaseDate,omitempty" bson:"release_date,omitempty" validate:"required,alphanum"`
	Duration      int            `json:"duration,omitempty" bson:"duration,omitempty" validate:"required,int"`
	Lyrics        string         `json:"lyrics,omitempty" bson:"lyrics,omitempty" validate:"required,alpha"`
	Explicit      bool           `json:"explicit,omitempty" bson:"explicit,omitempty" validate:"required"`
	Bitrates      []AudioBitrate `json:"bitrates,omitempty" bson:"bitrates,omitempty" validate:"required"`
}

// SongOption represent the functional options for the song entity
type SongOption func(*Song)

// WithSongID set the ID of the song
func WithSongID(id string) SongOption {
	return func(s *Song) {
		s.ID = id
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

// WithSongGenre set the genre of the song
func WithSongGenre(genre Genre) SongOption {
	return func(s *Song) {
		s.Genre = genre
	}
}

// WithSongReleaseDate set the release date of the song
func WithSongReleaseDate(releaseDate time.Time) SongOption {
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
