package entity

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

// AudioFeaturesOption represent the functional options for the audio features entity
type AudioFeaturesOption func(*AudioFeatures)

// WithAudioFeaturesTempo set the tempo of the audio features
func WithAudioFeaturesTempo(tempo int) AudioFeaturesOption {
	return func(a *AudioFeatures) {
		a.Tempo = tempo
	}
}

// WithAudioFeaturesAudioKey set the audio key of the audio features
func WithAudioFeaturesAudioKey(audioKey string) AudioFeaturesOption {
	return func(a *AudioFeatures) {
		a.AudioKey = audioKey
	}
}

// WithAudioFeaturesMode set the mode of the audio features
func WithAudioFeaturesMode(mode string) AudioFeaturesOption {
	return func(a *AudioFeatures) {
		a.Mode = mode
	}
}

// WithAudioFeaturesLoudness set the loudness of the audio features
func WithAudioFeaturesLoudness(loudness float64) AudioFeaturesOption {
	return func(a *AudioFeatures) {
		a.Loudness = loudness
	}
}

// WithAudioFeaturesEnergy set the energy of the audio features
func WithAudioFeaturesEnergy(energy float64) AudioFeaturesOption {
	return func(a *AudioFeatures) {
		a.Energy = energy
	}
}

// WithAudioFeaturesDanceability set the danceability of the audio features
func WithAudioFeaturesDanceability(danceability float64) AudioFeaturesOption {
	return func(a *AudioFeatures) {
		a.Danceability = danceability
	}
}

// WithAudioFeaturesSpeechiness set the speechiness of the audio features
func WithAudioFeaturesSpeechiness(speechiness float64) AudioFeaturesOption {
	return func(a *AudioFeatures) {
		a.Speechiness = speechiness
	}
}

// WithAudioFeaturesAcousticness set the acousticness of the audio features
func WithAudioFeaturesAcousticness(acousticness float64) AudioFeaturesOption {
	return func(a *AudioFeatures) {
		a.Acousticness = acousticness
	}
}

// WithAudioFeaturesInstrumentalness set the instrumentalness of the audio features
func WithAudioFeaturesInstrumentalness(instrumentalness float64) AudioFeaturesOption {
	return func(a *AudioFeatures) {
		a.Instrumentalness = instrumentalness
	}
}

// WithAudioFeaturesLiveness set the liveness of the audio features
func WithAudioFeaturesLiveness(liveness float64) AudioFeaturesOption {
	return func(a *AudioFeatures) {
		a.Liveness = liveness
	}
}

// WithAudioFeaturesVelence set the valance of the audio features
func WithAudioFeaturesVelence(valance float64) AudioFeaturesOption {
	return func(a *AudioFeatures) {
		a.Velence = valance
	}
}

// NewAudioFeatures creates a new audio features entity
func NewAudioFeatures(options ...AudioFeaturesOption) *AudioFeatures {
	audioFeatures := &AudioFeatures{}
	for _, option := range options {
		option(audioFeatures)
	}
	return audioFeatures
}
