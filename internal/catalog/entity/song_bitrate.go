package entity

// AudioBitrate is a value object that represent the bitrate of a song
type AudioBitrate struct {
	Bitrate  int    `json:"bitrate,omitempty" bson:"bitrate,omitempty" validate:"required,min=64,max=320"`
	AudioURL string `json:"audioUrl,omitempty" bson:"audio_url,omitempty" validate:"required,url"`
}

// AudioBitrateOption represent the functional options for the song bitrate entity
type AudioBitrateOption func(*AudioBitrate)

// WithAudioBitrateBitrate set the bitrate of the song bitrate
func WithAudioBitrateBitrate(bitrate int) AudioBitrateOption {
	return func(sb *AudioBitrate) {
		sb.Bitrate = bitrate
	}
}

// WithAudioBitrateAudioURL set the audio URL of the song bitrate
func WithAudioBitrateAudioURL(audioURL string) AudioBitrateOption {
	return func(sb *AudioBitrate) {
		sb.AudioURL = audioURL
	}
}

// NewAudioBitrate create a new song bitrate entity
func NewAudioBitrate(opts ...AudioBitrateOption) *AudioBitrate {
	AudioBitrate := &AudioBitrate{}
	for _, opt := range opts {
		opt(AudioBitrate)
	}
	return AudioBitrate
}
