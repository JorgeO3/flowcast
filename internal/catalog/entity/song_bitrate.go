package entity

// SongBitrate represent the bitrate of a song
type SongBitrate struct {
	ID       int
	SongID   int
	Bitrate  int
	AudioURL string
}

// SongBitrateOption represent the functional options for the song bitrate entity
type SongBitrateOption func(*SongBitrate)

// WithSongBitrateID set the ID of the song bitrate
func WithSongBitrateID(id int) SongBitrateOption {
	return func(sb *SongBitrate) {
		sb.ID = id
	}
}

// WithSongBitrateSongID set the song ID of the song bitrate
func WithSongBitrateSongID(songID int) SongBitrateOption {
	return func(sb *SongBitrate) {
		sb.SongID = songID
	}
}

// WithSongBitrateBitrate set the bitrate of the song bitrate
func WithSongBitrateBitrate(bitrate int) SongBitrateOption {
	return func(sb *SongBitrate) {
		sb.Bitrate = bitrate
	}
}

// WithSongBitrateAudioURL set the audio URL of the song bitrate
func WithSongBitrateAudioURL(audioURL string) SongBitrateOption {
	return func(sb *SongBitrate) {
		sb.AudioURL = audioURL
	}
}

// NewSongBitrate create a new song bitrate entity
func NewSongBitrate(opts ...SongBitrateOption) *SongBitrate {
	songBitrate := &SongBitrate{}
	for _, opt := range opts {
		opt(songBitrate)
	}
	return songBitrate
}
