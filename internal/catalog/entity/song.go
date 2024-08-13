package entity

import "time"

// Song represent a song entity
type Song struct {
	ID            int            `bson:"_id"`
	Title         string         `bson:"title,omitempty"`
	ArtistID      int            `bson:"artist_id,omitempty"`
	AlbumID       int            `bson:"album_id,omitempty"`
	AudioFeatures AudioFeatures  `bson:"audio_features,omitempty"`
	Genre         Genre          `bson:"genre,omitempty"`
	ReleaseDate   time.Time      `bson:"release_date,omitempty"`
	Duration      int            `bson:"duration,omitempty"`
	TrackNumber   int            `bson:"track_number,omitempty"`
	Lyrics        string         `bson:"lyrics,omitempty"`
	Explicit      bool           `bson:"explicit,omitempty"`
	Bitrates      []AudioBitrate `bson:"bitrates,omitempty"`
}

// SongOption represent the functional options for the song entity
type SongOption func(*Song)

// WithSongID set the ID of the song
func WithSongID(id int) SongOption {
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

// WithSongArtistID set the artist ID of the song
func WithSongArtistID(artistID int) SongOption {
	return func(s *Song) {
		s.ArtistID = artistID
	}
}

// WithSongAlbumID set the album ID of the song
func WithSongAlbumID(albumID int) SongOption {
	return func(s *Song) {
		s.AlbumID = albumID
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

// WithSongTrackNumber set the track number of the song
func WithSongTrackNumber(trackNumber int) SongOption {
	return func(s *Song) {
		s.TrackNumber = trackNumber
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
