package entity

// AudioFeatures represent the audio features of a song
type AudioFeatures struct {
	Tempo            int
	Key              int
	Mode             int
	Danceability     int
	Energy           int
	Speechiness      int
	Acousticness     int
	Instrumentalness int
	Liveness         int
	Valance          int
}

// SongBiterate represent the biterate of a song
type SongBiterate struct {
	bitrate  string
	audioURL string
}

// Song represent a song entity
type Song struct {
	ID            string
	Title         string
	Duration      int
	Explicit      bool
	Popularity    int
	AudioFeatures AudioFeatures
	AlbumID       string
	AlbumName     string
	ArtistIDs     []string
	ArtistNames   []string
	ReleaseDate   string
	Genres        []string
	Biterates     []SongBiterate
}

// New creates a new Song entity
func New() (*Song, error) {
	song := &Song{}

	return song, nil
}
